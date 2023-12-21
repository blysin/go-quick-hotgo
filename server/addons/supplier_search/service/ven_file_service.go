package service

import (
	"context"
	"github.com/antlabs/strsim"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
	"hotgo/addons/supplier_search/api/admin/ven"
	"hotgo/addons/supplier_search/dao"
	"hotgo/addons/supplier_search/logic/util"
	"hotgo/addons/supplier_search/model/entity"
	"hotgo/addons/supplier_search/model/input/venin"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	intEntity "hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	intService "hotgo/internal/service"
	"strings"
)

type SVenFile struct{}

func NewVenFileService() *SVenFile {
	return &SVenFile{}
}

var VenFileService = NewVenFileService()

// Model 充值订单ORM模型
func (s *SVenFile) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.VendorUploadFile.Ctx(ctx), option...)
}

// List 获取上传的文件列表
func (s *SVenFile) List(ctx context.Context, inp venin.VenFileListInp) (list []*entity.VendorUploadFile, err error) {
	mod := s.Model(ctx)

	mod = mod.Where("vendor_id = ?", inp.VendorId)

	if inp.FileName != nil && *inp.FileName != "" {
		mod = mod.Where("file_name like ?", "%"+*inp.FileName+"%")
	}
	if inp.CreatedAtStart != nil {
		mod = mod.Where("create_time >= ?", inp.CreatedAtStart)
	}
	if inp.CreatedAtEnd != nil {
		mod = mod.Where("create_time <= ?", inp.CreatedAtEnd)
	}
	err = mod.Order("id desc").Scan(&list)
	return
}

func (s *SVenFile) UploadFile(ctx context.Context, fileName string, attr *sysin.AttachmentListModel) (res *ven.UploadRes, err error) {
	// 解析excel文件
	// 获取当前用户
	user := contexts.GetUser(ctx)

	path, err := getFilePath(ctx, attr.Path)
	if err != nil {
		return nil, err
	}

	allColumns, presetColumn, err := analysisExcelHead(path)
	if err != nil {
		return nil, err
	}

	venFile := &entity.VendorUploadFile{
		VendorId:            0,
		FileName:            fileName,
		FileId:              attr.Id,
		ExceptionDataFileId: "",
		ValidNum:            0,
		ExceptionNum:        0,
		AllColumn:           strings.Join(allColumns, ","),
		CreateBy:            user.Id,
		UpdateBy:            user.Id,
	}
	mod := s.Model(ctx)
	id, err := mod.Data(venFile).InsertAndGetId()
	if err == nil {
		venFile.Id = id
	}

	//去掉fileName最后一个.之后的内容
	fileName = fileName[:strings.LastIndex(fileName, ".")]
	return &ven.UploadRes{
		PresetColumn: *presetColumn,
		Id:           id,
		FileName:     fileName,
		AllColumns:   allColumns,
	}, err
}

func (s *SVenFile) GetById(ctx context.Context, id int64) (entity.VendorUploadFile, error) {
	mod := s.Model(ctx)
	var venFile entity.VendorUploadFile
	err := mod.Where("id = ?", id).Scan(&venFile)
	return venFile, err
}

func (s *SVenFile) Update(ctx context.Context, file *entity.VendorUploadFile) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		mod := s.Model(ctx)
		_, err := mod.Data(file).Where("id = ?", file.Id).Update()
		return err
	})
}

func (s *SVenFile) ReUploadFile(ctx context.Context, vendor *entity.Vendor, attr *sysin.AttachmentListModel) (res *ven.ReUploadRes, err error) {
	path, err := getFilePath(ctx, attr.Path)
	if err != nil {
		return
	}

	presetColumn := &venin.PresetColumn{}

	err = gjson.DecodeTo(vendor.RequiredColumn, presetColumn)
	if err != nil {
		glog.Errorf(ctx, "解析预设字段失败，预设字段：%s", vendor.RequiredColumn)
		return nil, gerror.New("解析预设字段失败")
	}

	details, head, brandNames, err := analysisExcelDetail(path, presetColumn)
	if err != nil {
		return
	}

	userId := contexts.GetUserId(ctx)

	venFile := entity.VendorUploadFile{
		VendorId:  vendor.Id,
		FileName:  attr.Name,
		FileId:    attr.Id,
		ValidNum:  len(details),
		AllColumn: strings.Join(*head, ","),
		CreateBy:  userId,
		UpdateBy:  userId,
	}

	// 对比表头，如果不一致，返回错误
	srcHeads := make([]string, 0, 10)
	_ = gjson.DecodeTo(vendor.AllColumn, &srcHeads)
	for _, srcHead := range srcHeads {
		if !gstr.InArray(*head, srcHead) {
			return nil, gerror.New("上传的文件表头与原文件不一致")
		}
	}

	// 保存数据，1、新增file，2、新增或更新detail，
	existDetails, err := VenDetailService.ListByBrandNames(ctx, vendor.Id, brandNames)
	if err != nil {
		return
	}

	existDetailMap := make(map[string]*entity.VendorDetail, len(existDetails))
	for _, detail := range existDetails {
		existDetailMap[detail.Brand] = detail
	}

	updateList := make([]*entity.VendorDetail, 0, len(details))
	insertList := make([]*entity.VendorDetail, 0, len(details))

	rate, err := util.FromExchangeRate(ctx, vendor.Currency)
	if err != nil {
		return nil, gerror.Wrap(err, "获取汇率失败")
	}
	glog.Info(ctx, "当前汇率：", gjson.MustEncodeString(rate))

	presetDetails(details, vendor.Id, userId, rate)

	for _, detail := range details {
		existDetail := existDetailMap[detail.Brand]
		if existDetail != nil {
			detail.Id = existDetail.Id
			detail.Status = existDetail.Status
			detail.CreatedAt = existDetail.CreatedAt
			detail.CreateBy = existDetail.CreateBy
			updateList = append(updateList, detail)
		} else {
			insertList = append(insertList, detail)
		}
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		mod := s.Model(ctx)
		_, err := mod.Data(&venFile).Insert()
		if err != nil {
			return err
		}

		err = VenDetailService.SaveBatch(ctx, insertList)
		if err != nil {
			return err
		}

		err = VenDetailService.UpdateBatch(ctx, updateList)
		if err != nil {
			return err
		}

		return nil
	})

	return &ven.ReUploadRes{}, err

}

func analysisExcelDetail(path string, presetColumn *venin.PresetColumn) ([]*entity.VendorDetail, *[]string, *[]string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, nil, nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	//"Row"（行）和 "Cell"（单元格）

	// 读取第一行作为标题
	i := f.GetActiveSheetIndex()
	sheetName := f.GetSheetName(i)

	fullRows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, nil, nil, gerror.New("读取excel文件失败")
	}
	if (len(fullRows)) < 2 {
		return nil, nil, nil, gerror.New("excel文件内容为空")
	}

	vendorDetailArray := make([]*entity.VendorDetail, 0, len(fullRows)-1)
	head := fullRows[0]

	for index, row := range fullRows {
		if index == 0 {
			continue
		}
		vendorDetail := entity.VendorDetail{
			VendorData: "",
		}

		// 行转成map
		rowMap := make(map[string]string, len(row))
		for i, v := range row {
			rowMap[head[i]] = v
		}

		vendorDetail.VendorData = gjson.MustEncodeString(rowMap)

		vendorDetailArray = append(vendorDetailArray, &vendorDetail)
	}

	brandNames, err := setDetailValue(vendorDetailArray, presetColumn)
	if err != nil {
		return nil, nil, nil, err
	}

	return vendorDetailArray, &head, brandNames, nil
}

func setDetailValue(list []*entity.VendorDetail, presetColumn *venin.PresetColumn) (*[]string, error) {
	brandNames := make([]string, 10)
	for index, vendorDetail := range list {
		json, err := gjson.DecodeToJson(vendorDetail.VendorData)
		if err != nil {
			return nil, gerror.Wrap(err, "第"+gconv.String(index)+"行，解析数据失败")
		}

		if presetColumn.BrandName != "" {
			err := json.Get(presetColumn.BrandName).Scan(&vendorDetail.Brand)
			if err != nil {
				return nil, gerror.Wrap(err, "第"+gconv.String(index)+"行，解析品牌名失败")
			}
			brandNames = append(brandNames, vendorDetail.Brand)
			if vendorDetail.Brand == "" {
				return nil, gerror.New("第" + gconv.String(index) + "行，品牌名不能为空，字段名：" + presetColumn.BrandName)
			}
		}

		if presetColumn.BarCode != "" {
			err := json.Get(presetColumn.BarCode).Scan(&vendorDetail.Barcode)
			if err != nil {
				return nil, gerror.Wrap(err, "第"+gconv.String(index)+"行，解析条码失败")
			}
			if vendorDetail.Barcode == "" {
				return nil, gerror.New("第" + gconv.String(index) + "行，条码不能为空，字段名：" + presetColumn.BarCode)
			}
		}

		if presetColumn.EnName != "" {
			err := json.Get(presetColumn.EnName).Scan(&vendorDetail.EnglishName)
			if err != nil {
				return nil, gerror.Wrap(err, "第"+gconv.String(index)+"行，解析英文名失败")
			}
			if vendorDetail.EnglishName == "" {
				return nil, gerror.New("第" + gconv.String(index) + "行，英文名不能为空，字段名：" + presetColumn.EnName)
			}
		}

		if presetColumn.SupplyPrice != "" {
			vendorDetail.Cost = getRowValueInt64(json, presetColumn.SupplyPrice)

			err := json.Get(presetColumn.EnName).Scan(&vendorDetail.EnglishName)
			if err != nil {
				return nil, gerror.Wrap(err, "第"+gconv.String(index)+"行，解析英文名失败")
			}

			if vendorDetail.Cost <= 0 {
				return nil, gerror.New("第" + gconv.String(index) + "行，供货价不能为空或者负数，字段名：" + presetColumn.SupplyPrice)
			}
		}

		if presetColumn.SalePrice != "" {
			vendorDetail.SellingPrice = getRowValueInt64(json, presetColumn.SalePrice)
			if vendorDetail.SellingPrice <= 0 {
				return nil, gerror.New("第" + gconv.String(index) + "行，销售价不能为空或者负数，字段名：" + presetColumn.SalePrice)
			}
		}

		if presetColumn.VendorName != "" {
			err := json.Get(presetColumn.VendorName).Scan(&vendorDetail.Vendor)
			if err != nil {
				return nil, gerror.Wrap(err, "第"+gconv.String(index)+"行，解析供应商失败")
			}
			if vendorDetail.Vendor == "" {
				return nil, gerror.New("第" + gconv.String(index) + "行，供应商不能为空，字段名：" + presetColumn.VendorName)
			}
		}
	}
	// brandNames不允许重复
	hasDup, name := hasDuplicates(brandNames)
	if hasDup {
		return nil, gerror.New("存在重复品牌名:" + name)
	}
	return &brandNames, nil
}

func hasDuplicates(arr []string) (bool, string) {
	elements := make(map[string]bool)
	for _, v := range arr {
		if v == "" {
			continue
		}
		if elements[v] {
			return true, v
		}
		elements[v] = true
	}
	return false, ""
}

func getRowValueInt64(rowMap *gjson.Json, key string) int64 {
	v := (*rowMap).Get(key).String()
	if v == "" {
		return 0
	}
	return gconv.Int64(gconv.Float64(v) * 100)
}

func analysisExcelHead(path string) ([]string, *venin.PresetColumn, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	//"Row"（行）和 "Cell"（单元格）

	// 读取第一行作为标题
	i := f.GetActiveSheetIndex()
	sheetName := f.GetSheetName(i)

	rows, err := f.Rows(sheetName)
	if err != nil {
		return nil, nil, err
	}

	// 获取第一行
	firstRows, err := getFirstRows(rows)
	if err != nil {
		return nil, nil, err
	}

	//品牌名	条码	英文名	供货价	销售价	供应商
	presetColumn := venin.PresetColumn{}

	presetColumn.BrandName = matchKeyWord("品牌", firstRows)
	presetColumn.BarCode = matchKeyWord("条码", firstRows)
	presetColumn.EnName = matchKeyWord("英文名", firstRows)
	presetColumn.SupplyPrice = matchKeyWord("供货价", firstRows)
	presetColumn.SalePrice = matchKeyWord("销售价", firstRows)
	presetColumn.VendorName = matchKeyWord("供应商", firstRows)

	return firstRows, &presetColumn, nil
}

func getFilePath(ctx context.Context, attrPath string) (string, error) {
	//获取文件完整路径
	sp := g.Cfg().MustGet(ctx, "server.serverRoot")
	if sp.IsEmpty() {
		err := gerror.New("本地上传驱动必须配置静态路径!")
		return "", err
	}

	dir := gfile.Pwd()
	path := dir + "/" + strings.Trim(sp.String(), "/") + "/" + attrPath
	return path, nil
}

func matchKeyWord(key string, keywords []string) string {
	//匹配最佳的关键字
	for _, keyword := range keywords {
		if key == keyword {
			return keyword
		}
	}
	//计算相似度
	result := strsim.FindBestMatch(key, keywords)
	return result.Match.S
}

func getFirstRows(rows *excelize.Rows) ([]string, error) {
	results, cur := make([][]string, 0, 64), 0
	for rows.Next() {
		cur++
		row, err := rows.Columns()
		if err != nil {
			break
		}
		results = append(results, row)
		break
	}
	return results[0], rows.Close()
}

func (s *SVenFile) GetAttrByFileId(ctx context.Context, fileId int64) (res *intEntity.SysAttachment, venFileEntity *entity.VendorUploadFile, err error) {
	venFile, err := VenFileService.GetById(ctx, fileId)
	if err != nil {
		return
	}

	venFileEntity = &venFile

	fid := venFile.FileId
	res, err = intService.CommonUpload().GetFile(ctx, fid)
	return
}

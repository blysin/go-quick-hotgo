package service

import (
	"context"
	"github.com/antlabs/strsim"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/xuri/excelize/v2"
	"hotgo/addons/supplier_search/api/admin/ven"
	"hotgo/addons/supplier_search/dao"
	"hotgo/addons/supplier_search/model/entity"
	"hotgo/addons/supplier_search/model/input/venin"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/sysin"
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
	if inp.FileName != "" {
		mod = mod.Where("file_name like ?", "%"+inp.FileName+"%")
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
	// todo 解析excel文件
	// 获取当前用户
	user := contexts.GetUser(ctx)

	path, err := getFilePath(ctx, attr)
	if err != nil {
		return nil, err
	}

	allColumns, presetColumn, err := analysisExcel(path)
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

func analysisExcel(path string) ([]string, *venin.PresetColumn, error) {
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

func getFilePath(ctx context.Context, attr *sysin.AttachmentListModel) (string, error) {
	//获取文件完整路径
	sp := g.Cfg().MustGet(ctx, "server.serverRoot")
	if sp.IsEmpty() {
		err := gerror.New("本地上传驱动必须配置静态路径!")
		return "", err
	}

	dir := gfile.Pwd()
	path := dir + "/" + strings.Trim(sp.String(), "/") + "/" + attr.Path
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

package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"hotgo/addons/supplier_search/dao"
	"hotgo/addons/supplier_search/model/entity"
	"hotgo/addons/supplier_search/model/input/venin"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/sysin"
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

func (s *SVenFile) UploadFile(ctx context.Context, fileName string, attr *sysin.AttachmentListModel) (venFile *entity.VendorUploadFile, err error) {
	// todo 解析excel文件
	// 获取当前用户
	user := contexts.GetUser(ctx)

	venFile = &entity.VendorUploadFile{
		VendorId:            0,
		FileName:            fileName,
		FileId:              attr.Id,
		ExceptionDataFileId: "",
		ValidNum:            0,
		ExceptionNum:        0,
		AllColumn:           "",
		CreateBy:            user.Id,
		UpdateBy:            user.Id,
	}
	mod := s.Model(ctx)
	id, err := mod.Data(venFile).InsertAndGetId()
	if err == nil {
		venFile.Id = id
	}
	return
}

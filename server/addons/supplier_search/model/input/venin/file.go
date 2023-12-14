package venin

import (
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/addons/supplier_search/model/entity"
)

type VenFileListInp struct {
	FileName        string      `json:"fileName"            description:"文件名称"`
	CreateTimeStart *gtime.Time `json:"createTimeStart"          description:"创建时间-开始"`
	CreateTimeEnd   *gtime.Time `json:"createTimeEnd"          description:"创建时间-结束"`
}

type VenFileListModel struct {
	List []*entity.VendorUploadFile `json:"list"`
}

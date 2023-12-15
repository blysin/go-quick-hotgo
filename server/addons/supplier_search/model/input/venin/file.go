package venin

import (
	"github.com/gogf/gf/v2/os/gtime"
	"hotgo/addons/supplier_search/model/entity"
)

type VenFileListInp struct {
	FileName       string      `json:"fileName"            description:"文件名称"`
	CreatedAtStart *gtime.Time `json:"createdAtStart"          description:"创建时间-开始"`
	CreatedAtEnd   *gtime.Time `json:"createdAtEnd"          description:"创建时间-结束"`
}

type VenFileListModel struct {
	List []*entity.VendorUploadFile `json:"list"`
}

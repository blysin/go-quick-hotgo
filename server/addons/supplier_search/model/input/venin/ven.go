package venin

type VenSaveInp struct {
	Id           int64         `json:"id"            description:"主表id"`
	VendorName   string        `json:"vendorName"            description:"供应商名称"`
	FileId       int64         `json:"fileId"            description:"文件id"`
	Exchange     string        `json:"exchange"            description:"货币类型"`
	PresetColumn *PresetColumn `json:"presetColumn"            description:"预设列"`
}

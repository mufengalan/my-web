package convert

type ExportRes struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data []ExportData `json:"data"`
}

type ExportData struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	MD5          string `json:"md5"`
	FileSize     int64  `json:"fileSize"`
	LastModified int64  `json:"lastModified"`
	Type         string `json:"type"`
	ShowName     string `json:"showName"`
}

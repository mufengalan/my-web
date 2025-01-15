package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mufengalan/my-web/internal/service"
	"github.com/mufengalan/my-web/pkg/response"
	"github.com/mufengalan/my-web/pkg/utils"
)

type ExportHandler struct {
}

// ExportData ExportData
func (h *ExportHandler) ExportData(c *gin.Context) {
	//从headers获取token
	token := c.GetHeader("token")
	url := c.GetHeader("url")
	http := utils.NewHttp(token)
	resp, err := http.GetRequest(url)
	if err != nil {
		response.Err(c, -100, "请求失败")
		return
	}
	err = service.NewExportService().Export(resp)
	if err != nil {
		response.Err(c, -100, "导出失败")
		return
	}
	response.Success(c, nil)
}

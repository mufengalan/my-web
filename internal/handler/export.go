package handler

import (
	"net/http"

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
	httpClient := utils.NewHttp(token)
	resp, err := httpClient.GetRequest(url)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, 1000, err.Error())
		return
	}
	err = service.NewExportService().Export(resp)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, 1000, err.Error())
		return
	}
	response.Success(c, nil)
}

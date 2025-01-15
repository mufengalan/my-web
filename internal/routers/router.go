package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mufengalan/my-web/internal/handler"
)
import "github.com/mufengalan/my-web/internal/health"

// 路由注册
func RegisterRouter(router *gin.Engine) {
	exportHandler := handler.ExportHandler{}
	router.GET("health", health.Health)
	router.GET("example", handler.GetList)
	router.GET("export", exportHandler.ExportData)
}

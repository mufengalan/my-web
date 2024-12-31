package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mufengalan/my-web/internal/controller"
)
import "github.com/mufengalan/my-web/internal/health"

// 路由注册
func RegisterRouter(router *gin.Engine) {
	router.GET("/health", health.Health)
	router.GET("example", controller.GetList)
}

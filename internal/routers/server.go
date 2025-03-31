package routers

import (
	"errors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var srv *http.Server
var port = ":8080"
var timeoutInSeconds = 10

func InitRouter(logger *zap.Logger) {
	router := gin.New()
	// 替换默认中间件
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true)) // 日志格式
	router.Use(ginzap.RecoveryWithZap(logger, true))      // 异常恢复日志
	RegisterRouter(router)
	RunRouter(router)
}

func RunRouter(router *gin.Engine) {
	srv = &http.Server{
		Addr:    port,
		Handler: router,
	}
	// 统一处理 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "404 Not Found"})
	})
	if err := srv.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
		panic(err)
	}
}

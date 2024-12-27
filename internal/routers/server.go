package routers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var srv *http.Server
var port = ":8080"
var timeoutInSeconds = 10

func InitRouter() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
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

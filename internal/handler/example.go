package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mufengalan/my-web/internal/service"
	"github.com/mufengalan/my-web/pkg/response"
)

func GetList(c *gin.Context) {
	src := service.NewExampleService()
	list, err := src.GetExampleList()
	if err != nil {
		response.Err(c, -100, err.Error())
		return
	}
	response.Success(c, list)
}

package health

import "github.com/gin-gonic/gin"

func Health(ginCtx *gin.Context) {
	ginCtx.JSON(200, gin.H{
		"message": "hello world",
	})
}

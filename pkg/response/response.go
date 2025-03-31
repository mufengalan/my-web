package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

const (
	SuccessCode       = 0
	InvalidParamsCode = 10001
	UnauthorizedCode  = 10002
	InternalErrorCode = 10003
)

func Result(c *gin.Context, httpStatus int, code int, msg string, data interface{}) {
	c.JSON(httpStatus, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success(c *gin.Context, data interface{}) {
	Result(c, http.StatusOK, SuccessCode, "success", data)
}

func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	Result(c, http.StatusOK, SuccessCode, message, data)
}

func Created(c *gin.Context, data interface{}) {
	Result(c, http.StatusCreated, SuccessCode, "created", data)
}

func Updated(c *gin.Context, data interface{}) {
	Result(c, http.StatusOK, SuccessCode, "updated", data)
}

func Deleted(c *gin.Context) {
	Result(c, http.StatusNoContent, SuccessCode, "deleted", nil)
}

func Err(c *gin.Context, httpStatus int, code int, msg string) {
	Result(c, httpStatus, code, msg, nil)
}

func ErrInvalidParams(c *gin.Context) {
	Err(c, http.StatusBadRequest, InvalidParamsCode, "invalid parameters")
}

func ErrInternal(c *gin.Context) {
	Err(c, http.StatusInternalServerError, InternalErrorCode, "internal server error")
}

func ErrUnauthorized(c *gin.Context) {
	Err(c, http.StatusUnauthorized, UnauthorizedCode, "unauthorized")
}

func ErrWithData(c *gin.Context, httpStatus int, code int, msg string, data interface{}) {
	Result(c, httpStatus, code, msg, data)
}

package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	ERROR   = 400
	SUCCESS = 200
	UNAUTH  = 401
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func FailWithMessage(c *gin.Context, msg string) {
	Result(ERROR, map[string]interface{}{}, msg, c)
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(SUCCESS, data, "ok", c)
}

func OkWithMessage(c *gin.Context, msg string) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

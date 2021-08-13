package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rophie123/patient_registration/utils"
	"net/http"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())
	router.POST("/user/reg", Reg)
	router.POST("/user/login", Login)
	router.GET("/user/info", NeedLogin(), Info)
	router.POST("/photo/upload", Upload)
	router.GET("/user/list", NeedLogin(), List)
	router.StaticFS("/upload", http.Dir("./upload"))
	return router
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, X-Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func NeedLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("X-Token")
		if tokenString == "" {
			c.String(401, "Illegal operation")
			return
		}
		userName := utils.ParseToken(tokenString)
		if userName != "admin" {
			c.String(401, "Illegal operation")
			return
		}
		c.Next()
	}
}

func ReturnError(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code":    400,
		"message": msg,
	})
}

func ReturnData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "ok",
		"data":    data,
	})
}

func ReturnOK(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": msg,
	})
}

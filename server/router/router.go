package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rophie123/patient_registration/server/api/user"
	"github.com/rophie123/patient_registration/server/middleware"
	"net/http"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	userApi := &user.UserApiHandler{}
	router.POST("/user/reg", userApi.Reg)
	router.POST("/user/login", userApi.Login)
	router.POST("/photo/upload", userApi.Upload)

	router.GET("/user/info", middleware.NeedLogin(), userApi.Info)
	router.GET("/user/list", middleware.NeedLogin(), userApi.List)

	router.StaticFS("/upload", http.Dir("./upload"))

	return router
}

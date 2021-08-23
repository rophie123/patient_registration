package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rophie123/patient_registration/server/model/common/response"
	"github.com/rophie123/patient_registration/server/utils"
)

func NeedLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("X-Token")
		if tokenString == "" {
			c.String(response.UNAUTH, "Illegal operation")
			return
		}
		userName := utils.ParseToken(tokenString)
		if userName != "admin" {
			c.String(response.UNAUTH, "Illegal operation")
			return
		}
		c.Next()
	}
}

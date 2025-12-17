package router

import (
	"usercenter/app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	const pre = "/api"

	api := r.Group(pre)
	{
		api.POST("/auth", userController.AuthPassword)
		api.POST("/oauth", userController.OauthPassword)
		api.POST("/email", userController.EmailReset)
		api.POST("/activation/notVerify", userController.Activite)

		//不需要邮箱验证
		api.POST("/repass", userController.RePass)

		api.POST("/del", userController.DelAccount)
	}
}

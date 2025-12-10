package userController

import (
	"errors"
	"usercenter/app/apiExpection"
	"usercenter/app/utility"

	"github.com/gin-gonic/gin"
	"github.com/zjutjh/WeJH-SDK/oauth"
	"github.com/zjutjh/WeJH-SDK/oauth/oauthException"
)

type OauthData struct {
	StudentId string `json:"stu_id"`
	Password  string `json:"password"`
}

func OauthPassword(c *gin.Context) {
	var data OauthData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.JsonResponse(400, "请求参数错误", nil, c)
		return
	}
	_, e := oauth.Login(data.StudentId, data.Password)
	if e != nil {
		switch {
		case errors.Is(e, oauthException.ClosedError):
			utility.JsonResponse(apiExpection.ClosedError.Code, apiExpection.ClosedError.Msg, nil, c)
		case errors.Is(e, oauthException.WrongPassword):
			utility.JsonResponse(apiExpection.WrongPassword.Code, apiExpection.WrongPassword.Msg, nil, c)
		case errors.Is(e, oauthException.NotActivatedError):
			utility.JsonResponse(apiExpection.NotActivatedError.Code, apiExpection.NotActivatedError.Msg, nil, c)
		case errors.Is(e, oauthException.WrongAccount):
			utility.JsonResponse(apiExpection.WrongAccount.Code, apiExpection.WrongAccount.Msg, nil, c)
		case errors.Is(e, oauthException.OtherError):
			utility.JsonResponse(apiExpection.OtherError("其他错误").Code, apiExpection.OtherError("其他错误").Msg, nil, c)
		default:
			utility.JsonResponseInternalServerError(c)
		}
		return
	}
	utility.JsonSuccessResponse(c, nil)
}

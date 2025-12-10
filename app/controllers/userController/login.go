package userController

import (
	"errors"
	"usercenter/app/services/userService"
	"usercenter/app/utility"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct {
	StudentId string `json:"stu_id"`
	Password  string `json:"password"`
}

func AuthPassword(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.JsonResponse(400, "请求参数错误", nil, c)
		return
	}
	_, err = userService.GetUserByStudentId(data.StudentId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utility.JsonResponse(404, "该用户不存在", nil, c)
		return
	} else if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	flag := userService.CheckUserBYStudentIdAndPassword(data.StudentId, data.Password)
	if !flag {
		utility.JsonResponse(409, "密码错误", nil, c)
		return
	}
	utility.JsonResponse(200, "OK", nil, c)
}

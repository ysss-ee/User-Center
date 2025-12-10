package userService

import (
	"errors"
	"time"
	"usercenter/app/model"
	"usercenter/app/utility"
	"usercenter/config/database"

	"gorm.io/gorm"
)

func CreateUser(password, email, sid string, userType uint8) error {
	user, err := GetUserByStudentId(sid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err == nil {
		if !CheckUserBYStudentIdAndPassword(sid, password) {
			return errors.New("密码错误")
		}
		// 已存在且密码正确，无需重复写库
		return nil
	}

	user = &model.User{
		Password:   utility.Encryrpt(password),
		StudentId:  sid,
		Email:      email,
		Type:       userType,
		CreateTime: time.Now(),
	}
	return database.DB.Create(user).Error
}

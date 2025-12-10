package userService

import (
	"usercenter/app/model"
	"usercenter/app/utility"
	"usercenter/config/database"

	"gorm.io/gorm"
)

func UpdateUserEmailByStudentId(studentId, email string) error {
	result := database.DB.Model(&model.User{}).
		Where("student_id = ?", studentId).
		Update("email", email)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func UpdateUserPasswordByStudentId(studentId, password string) error {
	pass := utility.Encryrpt(password)
	result := database.DB.Model(&model.User{}).
		Where("student_id = ?", studentId).
		Update("password", pass)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

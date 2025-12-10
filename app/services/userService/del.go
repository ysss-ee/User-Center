package userService

import (
	"usercenter/app/model"
	"usercenter/config/database"

	"gorm.io/gorm"
)

func DelAccount(stuID string) error {
	result := database.DB.Where("student_id = ?", stuID).Delete(&model.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

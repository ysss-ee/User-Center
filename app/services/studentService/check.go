package studentService

import (
	"usercenter/app/model"
	"usercenter/config/database"
)

func CheckStudentBYSIDAndIID(sid string, iid string) bool {
	student := model.Student{}
	result := database.DB.Where(
		&model.Student{
			StudentId: sid,
		},
	).First(&student)
	student_iid_6 := student.Iid[len(student.Iid)-6:]
	iid_6 := iid[len(iid)-6:]
	if student_iid_6 != iid_6 || result.Error != nil {
		return false
	} else {
		return true
	}
}

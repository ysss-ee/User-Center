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
	if result.Error != nil || len(iid) < 6 || len(student.Iid) < 6 {
		return false
	}
	student_iid_6 := student.Iid[len(student.Iid)-6:]
	iid_6 := iid[len(iid)-6:]
	return student_iid_6 == iid_6
}

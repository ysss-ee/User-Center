package model

type Student struct {
	StudentId string `gorm:"type:varchar(20)"`
	Iid       string `gorm:"type:varchar(100)"`
}

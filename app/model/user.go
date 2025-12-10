package model

import (
	"time"
)

type User struct {
	StudentId  string
	UserId     int `gorm:"primary_key;AUTO_INCREMENT"`
	Password   string
	Email      string
	Type       uint8 // 0: 本科生 1: 研究生
	CreateTime time.Time
}

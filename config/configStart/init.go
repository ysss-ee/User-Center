package configStart

import (
	"usercenter/config/config"
	"usercenter/config/database"
)

func Init() {
	config.InitConfig()
	database.Init()
}

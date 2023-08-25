package global

import (
	"github.com/ikun666/go_webserver/utils"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	RedisClient *utils.RedisClient
)

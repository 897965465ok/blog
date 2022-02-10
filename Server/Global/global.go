package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Viper *viper.Viper = viper.GetViper()
)

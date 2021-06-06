package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func CreateData() {
	driverName := viper.GetString("datasource.driverName") // 数据库名字
	host := viper.GetString("datasource.host")             // 公网ip
	port := viper.GetString("datasource.port")             // 数据库端口
	database := viper.GetString("datasource.database")     // 库
	username := viper.GetString("datasource.username")     // 用户名
	password := viper.GetString("datasource.password")     // password
	charset := viper.GetString("datasource.rootcharset")   // 字符集
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(driverName, args)
	db.SingularTable(true)
	DB = db
	if err != nil {
		panic("errr" + err.Error())
	}

	// 自动建表
	// db.AutoMigrate(&model.User{})
	// db.AutoMigrate(&model.Banner{})
	// user := &model.User{
	// 	Name:      "小明",
	// 	Telephone: "15278164395",
	// 	Password:  "897965465ok",
	// }
	// first := &model.User{}

	// db.Create(user)
	// db.First(first)
	// fmt.Println("first", first)
}

func GetDB() *gorm.DB {

	return DB
}

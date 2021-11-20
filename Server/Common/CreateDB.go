package common

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetDB() *gorm.DB {
	host := viper.GetString("datasource.host")           // 公网ip
	port := viper.GetString("datasource.port")           // 数据库端口
	database := viper.GetString("datasource.database")   // 库
	username := viper.GetString("datasource.username")   // 用户名
	password := viper.GetString("datasource.password")   // password
	charset := viper.GetString("datasource.rootcharset") // 字符集
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        args,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
			// NameReplacer:  strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
		},
	})

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
	return DB
}

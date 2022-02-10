package common

import (
	"fmt"
	model "main/Model"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/Config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")           // 公网ip
	port := viper.GetString("datasource.port")           // 数据库端口
	database := viper.GetString("datasource.database")   // 库
	username := viper.GetString("datasource.username")   // 用户名
	password := viper.GetString("datasource.password")   // password
	charset := viper.GetString("datasource.rootcharset") // 字符集
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	DB, err := gorm.Open(mysql.Open(args),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
				// NameReplacer:  strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称
			}})

	if err != nil {
		panic("errr" + err.Error())
	}
	sqlDB, _ := DB.DB()
	func() {
		if !DB.Migrator().HasTable(&model.Tags{}) {
			DB.AutoMigrate(&model.Tags{})
		}
		if !DB.Migrator().HasTable(&model.Favorites{}) {
			DB.AutoMigrate(&model.Favorites{})
		}
		if !DB.Migrator().HasTable(&model.User{}) {
			DB.AutoMigrate(&model.User{})
		}
		if !DB.Migrator().HasTable(&model.Comment{}) {
			DB.AutoMigrate(&model.Comment{})
		}
		if !DB.Migrator().HasTable(&model.Article{}) {
			DB.AutoMigrate(&model.Article{})
		}
		if !DB.Migrator().HasTable(&model.ImgUrl{}) {
			DB.AutoMigrate(&model.ImgUrl{})
		}
		if !DB.Migrator().HasTable(&model.Banner{}) {
			DB.AutoMigrate(&model.Banner{})
		}
	}()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(256)
	// 设置时间吧。
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	return DB
}

package common

import (
	"fmt"

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
	sqlDB, err := DB.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	return DB
}

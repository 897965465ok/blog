package main

import (
	common "main/Common"
	global "main/Global"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	common.InitConfig()
	global.DB = common.InitDB()
	sql, _ := global.DB.DB()
	defer sql.Close()
	r := gin.Default()
	r = CollectRouter(r)
	r.Run(viper.GetString("server.host"))
}

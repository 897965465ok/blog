package main

import (
	common "main/Common"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	common.CreateData()
	r := gin.Default()
	r = CollectRouter(r)
	defer common.DB.Close()
	r.Run(viper.GetString("server.host"))
}

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

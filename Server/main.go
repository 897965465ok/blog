package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	r := gin.Default()
	r = CollectRouter(r)
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

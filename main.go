package main

import (
	"io"
	"multicliws/lib/redislib"
	"multicliws/routers"
	"multicliws/websocket"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	initLogger()
	redislib.InitRedis()

	route := gin.Default()
	routers.InitWebRouter(route)
	websocket.InitWebSocket(route)

	route.Run(":8080")
}

func initConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initLogger() {
	filePath := viper.GetString("log.path")
	f, _ := os.Create(filePath)
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)
	gin.SetMode(viper.GetString("log.mode"))
	gin.DisableConsoleColor()
}

package main

import (
	"github.com/PrashantDesale2004/go-sms-verify-yt/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//initialise the config
	app := api.Config{Router: router}
	//routes
	app.Routes()
	router.Run(":8000")
}

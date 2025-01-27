package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mod/config"
)

func main() {
	app := gin.Default()
	app.Use(gin.Recovery())


	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":    "server is running ftm!",
			"statusCode": 200,
		})
	})

	fmt.Println("Air is runnig...")

	port, err := config.GetEnvProperty("Port")
	if err != nil {
		fmt.Println("Error getting Port:", err)
		return
	}

	if port == "" {
		port = "8080"
	}

	app.Run(":" + port)
}
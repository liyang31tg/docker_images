package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	tmpG := gin.Default()
	tmpG.Static("/static", "./static")
	fmt.Println("server is start")
	tmpG.GET("/1", func(c *gin.Context) {
		home := os.Getenv("HOME")
		c.JSON(200, map[string]string{"status": home})
	})
	tmpG.Run()
}

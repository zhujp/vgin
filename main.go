package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/zhujp/vgin/app/util/setting"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(setting.HTTPPort) // listen and serve on 0.0.0.0:8080
}

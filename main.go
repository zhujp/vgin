package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/zhujp/vgin/app/util/setting"
	"github.com/zhujp/vgin/app/controller/post"
)

func main() {
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	post_router := r.Group("/post")
	{
		post_router.GET("/",post.Lists)
	}
	r.Run(setting.HTTPPort) // listen and serve on 0.0.0.0:8080
}

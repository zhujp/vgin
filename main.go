package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhujp/vgin/app/controller/post"
	"github.com/zhujp/vgin/app/util/setting"
	// "net/http"
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
		// post_router.GET("/", post.Lists)
		post_router.POST("/create", post.Create)
		post_router.GET("/:id", post.Detail)
		post_router.PUT("/:id", post.Update)
		post_router.DELETE("/:id", post.Delete)
	}
	r.Run(setting.HTTPPort) // listen and serve on 0.0.0.0:8080
}

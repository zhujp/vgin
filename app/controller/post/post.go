package post

import (
	"github.com/gin-gonic/gin"
)

// func Lists() {
// 	posts.GetPosts()
// }

func Create(c *gin.Context) {
	name := c.PostForm("name") //接收post数据
	// id := c.Query("id")  //接收url的参数
	// page := c.DefaultQuery("page", "0")
	c.JSON(200, gin.H{
		"message": name,
	})
}

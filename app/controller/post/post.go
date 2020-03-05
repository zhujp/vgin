package post

import (
	"github.com/gin-gonic/gin"
	"github.com/zhujp/vgin/app/model/posts"
)

// func Lists() {
// 	posts.GetPosts()
// }

func Create(c *gin.Context) {
	title := c.PostForm("title") //接收post数据
	body := c.PostForm("body")   //接收post数据
	post := &posts.Post{Title: title, Body: body}
	result, err := post.CreatePost()
	// id := c.Query("id")  //接收url的参数
	// page := c.DefaultQuery("page", "0")
	msg := "成功"
	if result == 0 {
		msg = err.Error()
	}
	c.JSON(200, gin.H{
		"message": msg,
	})
}

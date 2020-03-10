package post

import (
	"github.com/gin-gonic/gin"
	"github.com/zhujp/vgin/app/model/posts"
	"github.com/zhujp/vgin/app/util/response"
)

// func Lists() {
// 	posts.GetPosts()
// }

func Create(c *gin.Context) {
	title := c.PostForm("title") //接收post数据
	body := c.PostForm("body")   //接收post数据
	post := &posts.Post{Title: title, Body: body}
	err := post.CreatePost()
	// id := c.Query("id")  //接收url的参数
	// page := c.DefaultQuery("page", "0")
	msg := "成功"
	if err != nil {
		msg = err.Error()
	}

	response.Response(c, 1, msg, nil, "")
}

func Detail(c *gin.Context) {
	id := c.Param("id")
	post, err := posts.GetPost(id)

	if err != nil {
		response.Response(c, 0, "数据获取失败", nil, err.Error())
		return
	}

	response.Response(c, 1, "", post, "")
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {
	id := c.Param("id")
	err := posts.DeletePost(id)
	if err != nil {
		response.Response(c, 0, "删除失败", nil, err.Error())
		return
	}

	response.Response(c, 1, "删除成功", nil, "")

}

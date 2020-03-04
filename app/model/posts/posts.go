package posts

import "github.com/zhujp/vgin/app/model"

type Post struct {
	model.Model

	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetPosts(pageNum int, pageSize int, maps interface{}) (posts []Post) {

	model.Db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&posts)
	return
}

func CreatePost(title string, body string) {

}

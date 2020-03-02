package posts

import "github.com/zhujp/vgin/app/model"

type Tag struct {
    model.Model

    Title string `json:"name"`
}

func GetPosts(pageNum int, pageSize int, maps interface {}) (tags []Tag) {

	model.Db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}
package posts

import (
	"github.com/jinzhu/gorm"
	"github.com/zhujp/vgin/app/model"
)

type Post struct {
	gorm.Model

	Title string `json:"title"`
	Body  string `json:"body"`
	// CreatedAt int    `json:"created_at"`
}

func GetPosts(pageNum int, pageSize int, maps interface{}) (posts []Post) {

	model.Db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&posts)
	return
}

func (post *Post) CreatePost() (id int, err error) {
	// post.CreatedAt = int(time.Now().Unix())
	err = model.Db.Create(&post).Error
	if err != nil {
		return 0, err
	}

	return 1, nil
}

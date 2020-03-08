package posts

import (
	// "github.com/jinzhu/gorm"
	"github.com/zhujp/vgin/app/model"
)

type Post struct {
	model.Model

	Title string `json:"title"`
	Body  string `json:"body"`
	Views int `json:"views;default:'0'"`
	AuthorId int `json:"author_id;default:'0'"`
	Status int `json:"status;default:'0'"`
}

func GetPosts(pageNum int, pageSize int, maps interface{}) (posts []Post) {

	model.Db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&posts)
	return
}

func (post *Post) CreatePost() error {
	err := model.Db.Create(&post).Error
	if err != nil {
		return err
	}

	return nil
}

func GetPost(id string) (*Post,error) {

	post := new (Post)
	err := model.Db.First(&post, id).Error
	if err != nil {
		return nil,err
	}

	return post,nil
}


func DeletePost(id string) error {
	post := new (Post)
	err := model.Db.First(&post,id).Error

	if err != nil {
		return err
	}

	err = model.Db.Delete(&post).Error
	if err != nil {
		return err
	}

	return nil

}

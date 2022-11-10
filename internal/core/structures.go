package models

import "github.com/jinzhu/gorm"

type News struct {
	gorm.Model
	NewsId int64 `json:"newsId"`
	//CreatedAt  string `json:"createdAt"`
	Header     string `json:"header"`
	Body       string `json:"body"`
	PictureURL string `json:"pictureURL"`
}

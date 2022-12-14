package entities

import (
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Header     string `gorm:"column:header"`
	Body       string `gorm:"column:body"`
	PictureURL string `gorm:"column:pictureURL"`
}

func (t *News) TableName() string {
	return "news"
}

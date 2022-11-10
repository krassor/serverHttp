package database

import (
	"github.com/jinzhu/gorm"
	models "github.com/krassor/serverHttp/internal/core"
	u "github.com/krassor/serverHttp/pkg/utils"
)

func Validate(news *models.News) (map[string]interface{}, bool) {

	temp := &models.News{}

	//проверка на наличие ошибок и дубликатов электронных писем
	err := GetDB().Table("news").Where("header = ?", news.Header).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Header != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func Create(news *models.News) map[string]interface{} {

	if resp, ok := Validate(news); !ok {
		return resp
	}

	GetDB().Create(news)

	if news.ID <= 0 {
		return u.Message(false, "Failed to create news, connection error.")
	}

	response := u.Message(true, "News has been created")
	response["news"] = news
	return response
}

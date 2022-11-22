package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/krassor/serverHttp/internal/models/entities"
)

type NewsRepositoryInterface interface {
	FindAll() ([]entities.News, error)
	FindByID(id uint) (entities.News, error)
	Save(news entities.News) (entities.News, error)
	Delete(news entities.News) error
}

type newsRepository struct {
	DB *gorm.DB
}

func NewNewsRepostiory(DB *gorm.DB) NewsRepositoryInterface {
	return &newsRepository{
		DB: DB,
	}
}

func (n *newsRepository) FindAll() ([]entities.News, error) {
	var news []entities.News
	result := n.DB.Find(&news)
	if result.Error != nil {
		return []entities.News{}, result.Error
	}
	return news, nil
}

func (n *newsRepository) FindByID(id uint) (entities.News, error) {
	var news entities.News
	result := n.DB.First(&news, id)
	if result.Error != nil {
		return entities.News{}, result.Error
	}
	return news, nil
}

func (n *newsRepository) Save(news entities.News) (entities.News, error) {
	result := n.DB.Save(&news)
	if result.Error != nil {
		return entities.News{}, result.Error
	}
	return news, nil
}

func (n *newsRepository) Delete(news entities.News) error {
	result := n.DB.Delete(&news)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//-----------------------------------------------------------------
/*
func Validate(news *models.News) (map[string]interface{}, bool) {

	temp := &models.News{}

	//проверка на наличие ошибок и дубликатов электронных писем
	err := GetDB().Table("news").Where("header = ?", news.Header).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Header != "" {
		return u.Message(false, "The news already exist."), false
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

func GetAllNews() map[string]interface{} {

	newsSlice := make([]*models.News, 0)
	err := GetDB().Table("news").Where("id = ?", models.News.ID).Find(&newsSlice).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return newsSlice

}
*/

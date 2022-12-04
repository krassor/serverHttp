package repositories

import (
	"context"

	"github.com/krassor/serverHttp/internal/models/entities"
	"gorm.io/gorm"
)

type NewsRepositoryInterface interface {
	FindAll(ctx context.Context) ([]entities.News, error)
	FindByID(ctx context.Context, id uint) (entities.News, error)
	Save(ctx context.Context, news entities.News) (entities.News, error)
	Delete(ctx context.Context, news entities.News) error
}

type newsRepository struct {
	DB *gorm.DB
}

func NewNewsRepostiory(DB *gorm.DB) NewsRepositoryInterface {
	return &newsRepository{
		DB: DB,
	}
}

func (n *newsRepository) FindAll(ctx context.Context) ([]entities.News, error) {
	var news []entities.News

	result := n.DB.WithContext(ctx).Find(&news)
	if result.Error != nil {
		return []entities.News{}, result.Error
	}
	return news, nil
}

func (n *newsRepository) FindByID(ctx context.Context, id uint) (entities.News, error) {
	var news entities.News
	result := n.DB.WithContext(ctx).First(&news, id)
	if result.Error != nil {
		return entities.News{}, result.Error
	}
	return news, nil
}

func (n *newsRepository) Save(ctx context.Context, news entities.News) (entities.News, error) {
	result := n.DB.WithContext(ctx).Save(&news)
	if result.Error != nil {
		return entities.News{}, result.Error
	}
	return news, nil
}

func (n *newsRepository) Delete(ctx context.Context, news entities.News) error {
	result := n.DB.WithContext(ctx).Delete(&news)
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

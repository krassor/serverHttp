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

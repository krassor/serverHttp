package services

import (
	//"database/sql"
	//"golang-starter/infrastructures/logger"
	"context"
	"log"

	"github.com/krassor/serverHttp/internal/models/dto"
	"github.com/krassor/serverHttp/internal/models/entities"
	"github.com/krassor/serverHttp/internal/repositories"
	"gorm.io/gorm"
)

type NewsService interface {
	GetNews(ctx context.Context) ([]entities.News, error)
	GetNewsByID(ctx context.Context, newsID int) (entities.News, error)
	CreateNewNews(ctx context.Context, data dto.NewsRequestBody) (entities.News, error)
	DeleteNews(ctx context.Context, newsID int) error
}

type newsService struct {
	NewsRepository repositories.NewsRepositoryInterface
}

func NewNewsService(newsRepository repositories.NewsRepositoryInterface) NewsService {
	return &newsService{
		NewsRepository: newsRepository,
	}
}

func (repo newsService) GetNews(ctx context.Context) ([]entities.News, error) {

	news, err := repo.NewsRepository.FindAll(ctx)
	if err != nil {
		//logger.Log.Errorln(err)
		log.Println(err)
		return []entities.News{}, err
	}

	return news, nil
}

func (repo newsService) GetNewsByID(ctx context.Context, newsID int) (entities.News, error) {
	news, err := repo.NewsRepository.FindByID(ctx, uint(newsID))
	if err != nil {
		//logger.Log.Errorln(err)
		log.Println(err)
		return entities.News{}, err
	}

	return news, nil
}

func (repo newsService) CreateNewNews(ctx context.Context, data dto.NewsRequestBody) (entities.News, error) {
	news := entities.News{
		Header:     data.Header,
		Body:       data.Body,
		PictureURL: data.PictureURL,
	}

	result, err := repo.NewsRepository.Save(ctx, news)
	if err != nil {
		//logger.Log.Errorln(err)
		log.Println(err)
		return entities.News{}, err
	}

	return result, nil
}

func (repo newsService) DeleteNews(ctx context.Context, newsID int) error {

	news, err := repo.GetNewsByID(ctx, newsID)
	if err != nil {
		log.Println(err)
		return err
	}
	if news.ID == 0 {
		return gorm.ErrRecordNotFound
	}

	return repo.NewsRepository.Delete(ctx, news)
}

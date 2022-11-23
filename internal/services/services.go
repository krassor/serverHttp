package services

import (
	//"database/sql"
	//"golang-starter/infrastructures/logger"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/krassor/serverHttp/internal/models/dto"
	"github.com/krassor/serverHttp/internal/models/entities"
	"github.com/krassor/serverHttp/internal/repositories"
)

type NewsService interface {
	GetNews() ([]entities.News, error)
	GetNewsByID(newsID int) (entities.News, error)
	CreateNewNews(data dto.NewsRequestBody) (entities.News, error)
	DeleteNews(newsID int) error
}

type newsService struct {
	NewsRepository repositories.NewsRepositoryInterface
}

func NewNewsService(newsRepository repositories.NewsRepositoryInterface) NewsService {
	return &newsService{
		NewsRepository: newsRepository,
	}
}

func (repo newsService) GetNews() ([]entities.News, error) {
	// var products []entities.Products
	news, err := repo.NewsRepository.FindAll()
	if err != nil {
		//logger.Log.Errorln(err)
		log.Println(err)
		return []entities.News{}, err
	}

	return news, nil
}

func (repo newsService) GetNewsByID(newsID int) (entities.News, error) {
	news, err := repo.NewsRepository.FindByID(uint(newsID))
	if err != nil {
		//logger.Log.Errorln(err)
		log.Println(err)
		return entities.News{}, err
	}

	return news, nil
}

func (repo newsService) CreateNewNews(data dto.NewsRequestBody) (entities.News, error) {
	news := entities.News{
		Header:     data.Header,
		Body:       data.Body,
		PictureURL: data.PictureURL,
	}

	result, err := repo.NewsRepository.Save(news)
	if err != nil {
		//logger.Log.Errorln(err)
		log.Println(err)
		return entities.News{}, err
	}

	return result, nil
}

func (repo newsService) DeleteNews(newsID int) error {

	news, err := repo.GetNewsByID(newsID)
	if err != nil {
		log.Println(err)
		return err
	}
	if news.ID == 0 {
		return gorm.ErrRecordNotFound
	}

	return repo.NewsRepository.Delete(news)
}

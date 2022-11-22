package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/krassor/serverHttp/internal/models/dto"
	"github.com/krassor/serverHttp/internal/models/entities"
	"github.com/krassor/serverHttp/internal/services"
	"github.com/krassor/serverHttp/pkg/utils"

	"github.com/go-chi/chi/v5"
)

type NewsHandlers interface {
	GetNews(w http.ResponseWriter, r *http.Request)
	GetNewsByID(w http.ResponseWriter, r *http.Request)
	CreateNewNews(w http.ResponseWriter, r *http.Request)
	DeleteNewsByID(w http.ResponseWriter, r *http.Request)
}

type newsHandlers struct {
	NewsService services.NewsService
}

func NewHttpHandler(newsService services.NewsService) NewsHandlers {
	return &newsHandlers{
		NewsService: newsService,
	}
}

func (newsHandler newsHandlers) GetNews(w http.ResponseWriter, r *http.Request) {
	// get all news
	news, err := newsHandler.NewsService.GetNews()
	var newsResponse interface{}
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}
	newsResponse = CreateNewsListResponse(news)
	utils.Json(w, http.StatusOK, newsResponse)

}

func (newsHandler newsHandlers) GetNewsByID(w http.ResponseWriter, r *http.Request) {
	//не забыть вытащить ID
	rawNewsId := chi.URLParam(r, "newsId")
	newsId, err := strconv.Atoi(rawNewsId)
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}
	news, err := newsHandler.NewsService.GetNewsByID(newsId)
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}
	newsResponse := CreateNewsResponse(news)
	utils.Json(w, http.StatusOK, newsResponse)
}

func (newsHandler newsHandlers) DeleteNewsByID(w http.ResponseWriter, r *http.Request) {
	//не забыть вытащить ID
	rawNewsId := chi.URLParam(r, "newsId")
	newsId, err := strconv.Atoi(rawNewsId)
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}
	err = newsHandler.NewsService.DeleteNews(newsId)
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}
	newsResponse := utils.Message(true, "news deleted")
	utils.Json(w, http.StatusOK, newsResponse)
	
}

func (newsHandler newsHandlers) CreateNewNews(w http.ResponseWriter, r *http.Request) {
	news := dto.NewsRequestBody{}
	err := json.NewDecoder(r.Body).Decode(&news)
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}

	newsEntity, err := newsHandler.NewsService.CreateNewNews(news)
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}
	//newsResponseDto := CreateNewsResponse(newsEntity)
	var newsResponseParams dto.NewsResponseParams
	newsResponseParams.NewsID = newsEntity.ID
	newsResponse := utils.Message(false, newsResponseParams)
	utils.Json(w, http.StatusOK, newsResponse)
}

func CreateNewsResponse(news entities.News) dto.NewsResponseBody {
	return dto.NewsResponseBody{
		Header:     news.Header,
		Body:       news.Body,
		PictureURL: news.PictureURL,
	}
}

func CreateNewsListResponse(newslist []entities.News) []dto.NewsResponseBody {
	var newsListResponse []dto.NewsResponseBody
	for _, n := range newslist {
		news := CreateNewsResponse(n)
		newsListResponse = append(newsListResponse, news)
	}
	return newsListResponse
}

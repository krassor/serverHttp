package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
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
	GetFiles(w http.ResponseWriter, r *http.Request)
}

type newsHandlers struct {
	NewsService services.NewsService
}

func NewHttpHandler(newsService services.NewsService) NewsHandlers {
	return &newsHandlers{
		NewsService: newsService,
	}
}

func (newsHandler *newsHandlers) GetFiles(w http.ResponseWriter, r *http.Request) {
	workDir := "/home"
	filesDir := "/home/files" //http.Dir(filepath.Join(workDir, "files"))

	fp := filepath.Join(workDir, filepath.Clean(r.URL.Path))
	// Return a 404 if the file doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}
	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	fs := http.StripPrefix("/files/", http.FileServer(http.Dir(filesDir)))

	fs.ServeHTTP(w, r)

	//Example: http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
}

func (newsHandler *newsHandlers) GetNews(w http.ResponseWriter, r *http.Request) {
	// get all news
	news, err := newsHandler.NewsService.GetNews(r.Context())
	var newsResponse interface{}
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}
	newsResponse = createNewsListResponse(news)
	utils.Json(w, http.StatusOK, newsResponse)

}

func (newsHandler *newsHandlers) GetNewsByID(w http.ResponseWriter, r *http.Request) {
	rawNewsId := chi.URLParam(r, "newsId")
	newsId, err := strconv.Atoi(rawNewsId)
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}
	news, err := newsHandler.NewsService.GetNewsByID(r.Context(), newsId)
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}
	newsResponse := createNewsResponse(news)
	utils.Json(w, http.StatusOK, newsResponse)
}

func (newsHandler *newsHandlers) DeleteNewsByID(w http.ResponseWriter, r *http.Request) {

	rawNewsId := chi.URLParam(r, "newsId")
	newsId, err := strconv.Atoi(rawNewsId)
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}
	err = newsHandler.NewsService.DeleteNews(r.Context(), newsId)
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}
	newsResponse := utils.Message(true, "news deleted")
	utils.Json(w, http.StatusOK, newsResponse)

}

func (newsHandler *newsHandlers) CreateNewNews(w http.ResponseWriter, r *http.Request) {
	news := dto.NewsRequestBody{}
	err := json.NewDecoder(r.Body).Decode(&news)
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}

	newsEntity, err := newsHandler.NewsService.CreateNewNews(r.Context(), news)
	if err != nil {
		utils.Err(w, http.StatusInternalServerError, err)
		return
	}
	//newsResponseDto := CreateNewsResponse(newsEntity)
	var newsResponseParams dto.NewsResponseParams
	newsResponseParams.NewsID = newsEntity.ID
	newsResponse := utils.Message(true, newsResponseParams)
	utils.Json(w, http.StatusOK, newsResponse)
}

func createNewsResponse(news entities.News) dto.NewsResponseBody {
	return dto.NewsResponseBody{
		Header:     news.Header,
		Body:       news.Body,
		PictureURL: news.PictureURL,
	}
}

func createNewsListResponse(newslist []entities.News) []dto.NewsResponseBody {
	var newsListResponse []dto.NewsResponseBody
	for _, n := range newslist {
		news := createNewsResponse(n)
		newsListResponse = append(newsListResponse, news)
	}
	return newsListResponse
}

package services

import (
	"encoding/json"
	"net/http"

	models "github.com/krassor/serverHttp/internal/core"
	"github.com/krassor/serverHttp/internal/database"
	u "github.com/krassor/serverHttp/pkg/utils"
)

var CreateNews = func(w http.ResponseWriter, r *http.Request) {

	news := &models.News{}
	err := json.NewDecoder(r.Body).Decode(news) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := database.Create(news) //Создать аккаунт
	u.Respond(w, resp)
}

package httpServer

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	models "github.com/krassor/serverHttp/internal/core"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving", r.Host, "for", r.URL.Path)
	myT := template.Must(template.ParseGlob("github.com/krassor/serverHttp/html/home.gohtml"))
	myT.ExecuteTemplate(w, "github.com/krassor/serverHttp/html/home.gohtml", nil)

}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	var testNewsSlice []models.News
	testNews := models.News{
		NewsId: 1,
		//CreatedAt:  time.Now().Format(time.RFC3339),
		Header:     "Заголовок про монету",
		Body:       "Новость про монету",
		PictureURL: "https://cdnstatic.rg.ru/crop1850x1234/uploads/images/206/44/87/iStock-1142796736.jpg",
	}

	testNewsSlice = append(testNewsSlice, testNews)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(testNewsSlice)
	fmt.Println(testNewsSlice)
}

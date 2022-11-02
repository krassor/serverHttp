package httpServer

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	models "github.com/krassor/serverHttp/internal/core"
)

var DATASLICE []models.Coin

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

func coinsHandler(w http.ResponseWriter, r *http.Request) {

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(DATASLICE)
	fmt.Println(DATASLICE)

}

// func CHANGE(k string, n Coin) bool {
// 	DATA[k] = n
// 	return true
// }

func APPEND(n models.Coin) {
	DATASLICE = append(DATASLICE, n)
}

func listAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing the contents of the KV store!")
	fmt.Fprintf(w, "<a href=\"/\" style=\"margin-right: 20px;\">Home sweet home!</a>")
	fmt.Fprintf(w, "<a href=\"/list\" style=\"margin-right: 20px;\">List all elements!</a>")
	fmt.Fprintf(w, "<a href=\"/change\" style=\"margin-right: 20px;\">Change an element!</a>")
	fmt.Fprintf(w, "<a href=\"/insert\" style=\"margin-right: 20px;\">Insert new element!</a>")
	fmt.Fprintf(w, "<a href=\"/api/coins\" style=\"margin-right: 20px;\">JSON coins</a>")
	fmt.Fprintf(w, "<h1>The contents of the KV store are:</h1>")
	fmt.Fprintf(w, "<ul>")
	for k, v := range DATASLICE {
		fmt.Fprintf(w, "<li>")
		fmt.Fprintf(w, "<strong>%v</strong> with value: %v\n", k, v)
		fmt.Fprintf(w, "</li>")
	}
	fmt.Fprintf(w, "</ul>")
	fmt.Println(DATASLICE)
}

func changeElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Changing an element of the KV store!")
	tmpl := template.Must(template.ParseFiles("github.com/krassor/serverHttp/html/update.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	//key := r.FormValue("Key")
	y, _ := strconv.Atoi(r.FormValue("Year"))
	n := models.Coin{
		Country: r.FormValue("Country"),
		Region:  r.FormValue("Region"),
		Year:    y,
	}
	fmt.Println("n: ", n)

	// if !CHANGE(key, n) {
	// 	fmt.Println("Update operation failed!")
	// } else {
	APPEND(n)
	tmpl.Execute(w, struct{ Success bool }{true})
	//}
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	testNews := models.News{
		Id:         1,
		CreatedAt:  time.Now().Format(time.RFC3339),
		Header:     "Заголовок про монету",
		Body:       "Новость про монету",
		PictureURL: "https://cdnstatic.rg.ru/crop1850x1234/uploads/images/206/44/87/iStock-1142796736.jpg",
	}

	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(testNews)
	fmt.Println(testNews)
}

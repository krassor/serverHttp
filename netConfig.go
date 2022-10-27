package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strconv"
	"time"

	p "github.com/krassor/serverHttp/proto/pb"
	"google.golang.org/grpc"
)

type Coin struct {
	Country string
	Region  string
	Year    int
}

type ErrorJson struct {
	ErrorText string
}

type MessageServer struct {
}

var portGrpc = ":8080"

var DATA = make(map[string]Coin)
var DATASLICE []Coin
var DATAFILE = "/tmp/dataFile.gob"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving", r.Host, "for", r.URL.Path)
	myT := template.Must(template.ParseGlob("home.gohtml"))
	myT.ExecuteTemplate(w, "home.gohtml", nil)

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

func CHANGE(k string, n Coin) bool {
	DATA[k] = n
	return true
}

func APPEND(n Coin) {
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
	tmpl := template.Must(template.ParseFiles("update.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	//key := r.FormValue("Key")
	y, _ := strconv.Atoi(r.FormValue("Year"))
	n := Coin{
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

func (MessageServer) SayIt(ctx context.Context, r *p.Request) (*p.Response, error) {
	fmt.Println("Request Text:", r.Text)
	fmt.Println("Request SubText:", r.Subtext)
	response := &p.Response{
		Text:    r.Text,
		Subtext: "Got it!",
	}
	return response, nil
}

func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("using default port: ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}

	r := http.NewServeMux()

	srv := &http.Server{
		Addr:         PORT,
		Handler:      r,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	r.HandleFunc("/time", timeHandler)
	r.HandleFunc("/", defaultHandler)
	r.HandleFunc("/api/coins", coinsHandler)
	r.HandleFunc("/change", changeElement)
	r.HandleFunc("/list", listAll)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
	server := grpc.NewServer()
	var messageServer MessageServer
	p.RegisterMessageServiceServer(server, messageServer)
	listen, err := net.Listen("tcp", portGrpc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Serving requests...")
	server.Serve(listen)

}

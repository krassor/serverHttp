package main

import (
	"bufio"
	//"context"
	"encoding/json"
	"fmt"
	"html/template"

	//"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strconv"
	"strings"
	"time"

	grpcm "github.com/krassor/serverHttp/internal/transport/grpc"
	sm "github.com/krassor/serverHttp/pkg/supportModule"
)

type Coin struct {
	Country string
	Region  string
	Year    int
}

type ErrorJson struct {
	ErrorText string
}

var DATA = make(map[string]Coin)
var DATASLICE []Coin
var DATAFILE = "/tmp/dataFile.gob"

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
	tmpl := template.Must(template.ParseFiles("github.com/krassor/serverHttp/html/update.gohtml"))
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

func main() {

	//arguments := os.Args
	// if len(arguments) == 1 {
	// 	fmt.Println("using default http port: ", PORT)
	// 	fmt.Println("using default grpc port: ", portGrpc)
	// } else {
	// 	PORT = ":" + arguments[1]
	// }

	go func() {
		//fmt.Println(time.Now().Format(time.RFC3339), " :", "Server HTTP starting...")
		sm.PrintlnWithTimeShtamp("Server HTTP starting...")
		PORT := ":8001"
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

		sm.PrintlnWithTimeShtamp("Server HTTP listening...")
		err := srv.ListenAndServe()
		defer srv.Close()
		if err != nil {
			fmt.Println(err)
			return
		}

	}()

	go func() {
		if err := grpcm.ServerGrpcStart("8080"); err != nil {
			fmt.Println(err)
		}
	}()

	time.Sleep(1 * time.Second)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(string(text))) == "stop" {
			fmt.Println("Program exiting...")
			return
		}
	}
	//fmt.Println("End program")
}

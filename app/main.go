package main

import (
	"context"
	"time"

	"github.com/krassor/serverHttp/internal/graceful"
	"github.com/krassor/serverHttp/internal/repositories"
	"github.com/krassor/serverHttp/internal/services"
	httpServer "github.com/krassor/serverHttp/internal/transport/rest"
	"github.com/krassor/serverHttp/internal/transport/rest/handlers"
	"github.com/krassor/serverHttp/internal/transport/rest/routers"
)

//var DATA = make(map[string]Coin)

//var DATAFILE = "/tmp/dataFile.gob"

func main() {

	db := repositories.InitDB()
	newsRepository := repositories.NewNewsRepostiory(db)
	newsService := services.NewNewsService(newsRepository)
	newsHandler := handlers.NewHttpHandler(newsService)
	newsHandlerImpl := handlers.NewHttpHandlerImpl(newsHandler)
	newsRouter := routers.NewHttpRouter(newsHandlerImpl)
	newsHttpServer := httpServer.NewHttpServer(newsRouter)

	maxSecond := 5 * time.Second
	graceful.GracefulShutdown(
		context.TODO(),
		maxSecond,
		map[string]graceful.Operation{
			"http": func(ctx context.Context) error {
				return newsHttpServer.Shutdown(ctx)
			},
		},
	)

	newsHttpServer.Listen()

	// time.Sleep(1 * time.Second)
	// for {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print(">> ")
	// 	text, _ := reader.ReadString('\n')
	// 	if strings.ToLower(strings.TrimSpace(string(text))) == "stop" {
	// 		fmt.Println("Program exiting...")
	// 		return
	// 	}
	// }
	//fmt.Println("End program")
}

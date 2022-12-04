package main

import (
	"context"
	"time"

	"github.com/krassor/serverHttp/internal/graceful"
	"github.com/krassor/serverHttp/internal/logger"
	"github.com/krassor/serverHttp/internal/repositories"
	"github.com/krassor/serverHttp/internal/services"
	httpServer "github.com/krassor/serverHttp/internal/transport/rest"
	"github.com/krassor/serverHttp/internal/transport/rest/handlers"
	"github.com/krassor/serverHttp/internal/transport/rest/routers"
)

//var DATA = make(map[string]Coin)

//var DATAFILE = "/tmp/dataFile.gob"

func main() {

	logger.InitLogger()

	db := repositories.InitDB()
	newsRepository := repositories.NewNewsRepostiory(db)
	newsService := services.NewNewsService(newsRepository)
	newsHandler := handlers.NewHttpHandler(newsService)
	newsHandlerImpl := handlers.NewHttpHandlerImpl(newsHandler)
	newsRouter := routers.NewHttpRouter(newsHandlerImpl)
	newsHttpServer := httpServer.NewHttpServer(newsRouter)

	maxSecond := 10 * time.Second
	graceful.GracefulShutdown(
		context.Background(),
		maxSecond,
		map[string]graceful.Operation{
			"http": func(ctx context.Context) error {
				return newsHttpServer.Shutdown(ctx)
			},
		},
	)

	newsHttpServer.Listen()

}

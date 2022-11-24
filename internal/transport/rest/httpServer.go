package httpServer

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/krassor/serverHttp/internal/transport/rest/routers"
)

type HttpImpl struct {
	HttpRouter *routers.HttpRouterImpl
	httpServer *http.Server
}

func NewHttpServer(
	HttpRouter *routers.HttpRouterImpl,
) *HttpImpl {
	return &HttpImpl{
		HttpRouter: HttpRouter,
	}
}

func (p *HttpImpl) setupRouter(app *chi.Mux) {
	p.HttpRouter.Router(app)
}

func (p *HttpImpl) Listen() {
	app := chi.NewRouter()

	p.setupRouter(app)

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Println(e)
	}

	serverPort := os.Getenv("http_port")
	p.httpServer = &http.Server{
		Addr:    serverPort,
		Handler: app,
	}
	log.Printf("Server started on Port %s ", serverPort)
	p.httpServer.ListenAndServe()
}

func (p *HttpImpl) Shutdown(ctx context.Context) error {
	if err := p.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

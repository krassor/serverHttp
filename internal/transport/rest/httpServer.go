package httpServer

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	um "github.com/krassor/serverHttp/pkg/utils"
)

func ServerHttpStart() error {

	router := mux.NewRouter()

	port := os.Getenv("PORT") //Получить порт из файла .env; мы не указали порт, поэтому при локальном тестировании должна возвращаться пустая строка
	if port == "" {
		port = "8001" //localhost
	}

	um.PrintlnWithTimeShtamp(fmt.Sprintf("Server HTTP starting with port: %s", port))

	err := http.ListenAndServe(":"+port, router) //Запустите приложение, посетите localhost:8000/api

	if err != nil {
		return err
	}

	return nil
}

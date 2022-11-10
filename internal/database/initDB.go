package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	models "github.com/krassor/serverHttp/internal/core"
)

var db *gorm.DB //база данных

func InitDB() {

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	switch os.Getenv("db_type") {
	case "postgres":
		{
			username := os.Getenv("db_user")
			password := os.Getenv("db_pass")
			dbName := os.Getenv("db_name")
			dbHost := os.Getenv("db_host")
			dbPort := os.Getenv("db_port")

			dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Создать строку подключения
			fmt.Println(dbUri)

			conn, err := gorm.Open("postgres", dbUri)
			if err != nil {
				fmt.Print(err)
			}

			db = conn
			db.Debug().AutoMigrate(&models.News{}) //Миграция базы данных

		}

	}
}

// возвращает дескриптор объекта DB
func GetDB() *gorm.DB {
	return db
}

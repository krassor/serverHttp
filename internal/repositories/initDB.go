package repositories

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/krassor/serverHttp/internal/models/entities"
	"github.com/rs/zerolog/log"
)

func InitDB() *gorm.DB {

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		log.Warn().Msgf("godotenv.Load() Error: %s", e)
	}

	username := os.Getenv("NEWS_DB_USER")
	password := os.Getenv("NEWS_DB_PASSWORD")
	dbName := os.Getenv("NEWS_DB_NAME")
	dbHost := os.Getenv("NEWS_DB_HOST")
	dbPort := os.Getenv("NEWS_DB_PORT")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Создать строку подключения
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Error().Msgf("Error gorm.Open(): %s", err)
	}
	log.Info().Msg("gorm have connected to database")

	conn.Debug().AutoMigrate(&entities.News{}) //Миграция базы данных
	//log.Info().Msg("gorm have connected to database")

	return conn
}

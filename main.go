package main

import (
	"log"
	"os"

	_ "github.com/Anabol1ks/music-library/docs"
	"github.com/Anabol1ks/music-library/internal/models"
	"github.com/Anabol1ks/music-library/internal/songs"
	"github.com/Anabol1ks/music-library/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Онлайн библиотеки песен
func main() {
	key := os.Getenv("TEST_ENV")
	if key == "" {
		log.Println("\nПеременной среды нет, используется .env")
		// Загружаем .env
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Ошибка загрузки .env файла")
		}
	}

	storage.ConnectDatabase()

	storage.DB.AutoMigrate(&models.Song{})

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/songs", songs.CreateSong)
	r.PATCH("/songs/:id", songs.UpdateSong)
	r.GET("/songs/:id/text", songs.GetSongTextWithPagination)
	r.GET("/songs", songs.GetSongsWithFilters)
	r.DELETE("/songs/:id", songs.DeleteSong)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}

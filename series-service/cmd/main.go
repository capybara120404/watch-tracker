package main

import (
	"log"

	"github.com/capybara120404/common/config"
	"github.com/capybara120404/common/database"
	"github.com/capybara120404/common/utils"
	"github.com/capybara120404/series-service/internal/handler"
	"github.com/capybara120404/series-service/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.NewConfig("config.env")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	logger, file, err := utils.CreateLogger(config.LogFile, config.NameOFLogger)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	storage, err := database.Open(config.ConnectionString)
	if err != nil {
		logger.Fatal(err)
	}

	repository := repository.NewSeriesRepository(storage)
	handler := handler.NewSeriesHandler(logger, repository)

	router := gin.Default()

	seriesGroup := router.Group("/series")
	{
		seriesGroup.DELETE("/:id", handler.DeleteSeriesByIDHandler)
		seriesGroup.GET("", handler.GetAllSeriesHandler)
		seriesGroup.GET("/:id", handler.GetSeriesByIdHandler)
		seriesGroup.POST("", handler.AddSeriesHandler)
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Run(":8000")
}

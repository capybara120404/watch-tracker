package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/capybara120404/common/models"
	"github.com/capybara120404/series-service/internal/repository"
	"github.com/gin-gonic/gin"
)

type seriesHandler struct {
	logger     *log.Logger
	repository *repository.SeriesRepository
}

func NewSeriesHandler(logger *log.Logger, repository *repository.SeriesRepository) *seriesHandler {
	return &seriesHandler{
		logger:     logger,
		repository: repository,
	}
}

func (handler *seriesHandler) AddSeriesHandler(c *gin.Context) {
	var series models.Series

	err := c.ShouldBind(&series)
	if err != nil {
		handler.logAndRespond(c, http.StatusBadRequest, "Invalid JSON", err)
		return
	}

	err = handler.repository.AddSeries(series)
	if err != nil {
		handler.logAndRespond(c, http.StatusInternalServerError, "Failed to add series", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Series received successfully",
		"series":  series,
	})
}

func (handler *seriesHandler) DeleteSeriesByIDHandler(c *gin.Context) {
	id, err := handler.parseID(c)
	if err != nil {
		handler.logAndRespond(c, http.StatusBadRequest, "Invalid series ID", err)
		return
	}

	err = handler.repository.DeleteSeriesByID(id)
	if err != nil {
		handler.logAndRespond(c, http.StatusInternalServerError, "Failed to delete series", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Series deleted successfully",
	})
}

func (handler *seriesHandler) GetAllSeriesHandler(c *gin.Context) {
	series, err := handler.repository.GetAllSeries()
	if err != nil {
		handler.logAndRespond(c, http.StatusInternalServerError, "Failed to fetch series", err)
		return
	}

	if len(series) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No series found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Series received successfully",
		"series":  series,
	})
}

func (handler *seriesHandler) GetSeriesByIdHandler(c *gin.Context) {
	id, err := handler.parseID(c)
	if err != nil {
		handler.logAndRespond(c, http.StatusBadRequest, "Invalid series ID", err)
		return
	}

	series, err := handler.repository.GetSeriesById(id)
	if err != nil {
		handler.logAndRespond(c, http.StatusInternalServerError, "Failed to fetch series", err)
		return
	}

	if series == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Series not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Series retrieved successfully",
		"series":  series,
	})
}

func (handler *seriesHandler) parseID(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func (handler *seriesHandler) logAndRespond(c *gin.Context, statusCode int, message string, err error) {
	handler.logger.Println(err)
	c.JSON(statusCode, gin.H{
		"error":   message,
		"details": err.Error(),
	})
}

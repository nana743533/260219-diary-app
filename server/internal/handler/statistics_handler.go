package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nana743533/260219-diary-app/server/internal/model"
	"github.com/nana743533/260219-diary-app/server/internal/service"
)

type StatisticsHandler struct {
	service *service.DiaryService
}

func NewStatisticsHandler(service *service.DiaryService) *StatisticsHandler {
	return &StatisticsHandler{service: service}
}

func (h *StatisticsHandler) GetSummary(c *gin.Context) {
	userID := "default-user"
	period := c.DefaultQuery("period", "month")

	stats, err := h.service.GetStatistics(userID, period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: "Failed to fetch statistics",
			},
		})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (h *StatisticsHandler) GetTrend(c *gin.Context) {
	userID := "default-user"
	days, _ := strconv.Atoi(c.DefaultQuery("days", "30"))

	if days <= 0 {
		days = 30
	}

	trend, err := h.service.GetTrend(userID, days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: "Failed to fetch trend data",
			},
		})
		return
	}

	c.JSON(http.StatusOK, trend)
}

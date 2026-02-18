package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nana743533/260219-diary-app/server/internal/model"
	"github.com/nana743533/260219-diary-app/server/internal/service"
)

type CalendarHandler struct {
	service *service.DiaryService
}

func NewCalendarHandler(service *service.DiaryService) *CalendarHandler {
	return &CalendarHandler{service: service}
}

func (h *CalendarHandler) GetMonth(c *gin.Context) {
	userID := "default-user"

	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "VALIDATION_ERROR",
				Message: "Invalid year",
			},
		})
		return
	}

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil || month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "VALIDATION_ERROR",
				Message: "Invalid month",
			},
		})
		return
	}

	data, err := h.service.GetCalendarData(userID, year, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: "Failed to fetch calendar data",
			},
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *CalendarHandler) GetRange(c *gin.Context) {
	userID := "default-user"

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "VALIDATION_ERROR",
				Message: "start_date and end_date are required",
			},
		})
		return
	}

	diaries, err := h.service.GetAll(userID, startDate, endDate, 1000, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: "Failed to fetch calendar data",
			},
		})
		return
	}

	entries := make([]model.CalendarEntry, len(diaries))
	for i, d := range diaries {
		entries[i] = model.CalendarEntry{
			Date:   d.Date,
			Rating: d.Rating,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"start_date": startDate,
		"end_date":   endDate,
		"entries":    entries,
	})
}

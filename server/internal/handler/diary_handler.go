package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nana743533/260219-diary-app/server/internal/model"
	"github.com/nana743533/260219-diary-app/server/internal/service"
)

type DiaryHandler struct {
	service *service.DiaryService
}

func NewDiaryHandler(service *service.DiaryService) *DiaryHandler {
	return &DiaryHandler{service: service}
}

func (h *DiaryHandler) Create(c *gin.Context) {
	// 認証なしバージョン: 固定のユーザーIDを使用
	userID := "default-user"

	var req model.CreateDiaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "VALIDATION_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	diary, err := h.service.Create(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: "Failed to create diary",
			},
		})
		return
	}

	c.JSON(http.StatusCreated, diary)
}

func (h *DiaryHandler) GetAll(c *gin.Context) {
	userID := "default-user"

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "30"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	diaries, err := h.service.GetAll(userID, startDate, endDate, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: "Failed to fetch diaries",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"diaries": diaries,
		"pagination": gin.H{
			"total":  len(diaries),
			"limit":  limit,
			"offset": offset,
		},
	})
}

func (h *DiaryHandler) GetByDate(c *gin.Context) {
	userID := "default-user"
	date := c.Param("date")

	diary, err := h.service.GetByDate(userID, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: "Failed to fetch diary",
			},
		})
		return
	}

	if diary == nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "NOT_FOUND",
				Message: "Diary not found",
			},
		})
		return
	}

	c.JSON(http.StatusOK, diary)
}

func (h *DiaryHandler) Update(c *gin.Context) {
	userID := "default-user"
	date := c.Param("date")

	var req model.UpdateDiaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "VALIDATION_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	diary, err := h.service.Update(userID, date, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: "Failed to update diary",
			},
		})
		return
	}

	c.JSON(http.StatusOK, diary)
}

func (h *DiaryHandler) Delete(c *gin.Context) {
	userID := "default-user"
	date := c.Param("date")

	err := h.service.Delete(userID, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: model.ErrorDetail{
				Code:    "INTERNAL_ERROR",
				Message: "Failed to delete diary",
			},
		})
		return
	}

	c.Status(http.StatusNoContent)
}

package model

import "time"

type Diary struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Date        string    `json:"date"`        // YYYY-MM-DD
	Rating      int       `json:"rating" binding:"required,min=1,max=5"`
	Progress    string    `json:"progress" binding:"required,oneof=A B C"`
	WakeUpTime  string    `json:"wake_up_time" binding:"required"`
	SleepTime   string    `json:"sleep_time" binding:"required"`
	Memo        string    `json:"memo"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateDiaryRequest struct {
	Date       string `json:"date" binding:"required,datetime=2006-01-02"`
	Rating     int    `json:"rating" binding:"required,min=1,max=5"`
	Progress   string `json:"progress" binding:"required,oneof=A B C"`
	WakeUpTime string `json:"wake_up_time" binding:"required"`
	SleepTime  string `json:"sleep_time" binding:"required"`
	Memo       string `json:"memo"`
}

type UpdateDiaryRequest struct {
	Rating     *int    `json:"rating" binding:"omitempty,min=1,max=5"`
	Progress   *string `json:"progress" binding:"omitempty,oneof=A B C"`
	WakeUpTime *string `json:"wake_up_time"`
	SleepTime  *string `json:"sleep_time"`
	Memo       *string `json:"memo"`
}

type CalendarEntry struct {
	Date   string `json:"date"`
	Rating int    `json:"rating"`
}

type CalendarResponse struct {
	Year    int             `json:"year"`
	Month   int             `json:"month"`
	Entries []CalendarEntry `json:"entries"`
	Summary CalendarSummary  `json:"summary"`
}

type CalendarSummary struct {
	TotalDays    int     `json:"total_days"`
	RecordedDays int     `json:"recorded_days"`
	AverageRating float64 `json:"average_rating"`
}

type Statistics struct {
	Period               string            `json:"period"`
	PeriodStart          string            `json:"period_start"`
	PeriodEnd            string            `json:"period_end"`
	TotalEntries         int               `json:"total_entries"`
	AverageRating        float64           `json:"average_rating"`
	RatingDistribution   map[string]int    `json:"rating_distribution"`
	ProgressDistribution map[string]int    `json:"progress_distribution"`
	AverageWakeUpTime    string            `json:"average_wake_up_time"`
	AverageSleepTime     string            `json:"average_sleep_time"`
	LongestStreak        int               `json:"longest_streak"`
}

type TrendData struct {
	PeriodDays int           `json:"period_days"`
	Data       []TrendEntry  `json:"data"`
}

type TrendEntry struct {
	Date   string `json:"date"`
	Rating int    `json:"rating"`
}

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

package service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/nana743533/260219-diary-app/server/internal/model"
	"github.com/google/uuid"
)

type DiaryService struct {
	db *sql.DB
}

func NewDiaryService(db *sql.DB) *DiaryService {
	return &DiaryService{db: db}
}

func (s *DiaryService) Create(userID string, req model.CreateDiaryRequest) (*model.Diary, error) {
	id := uuid.New().String()
	now := time.Now()

	query := `
		INSERT INTO diaries (id, user_id, date, rating, progress, wake_up_time, sleep_time, memo, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := s.db.Exec(query, id, userID, req.Date, req.Rating, req.Progress, req.WakeUpTime, req.SleepTime, req.Memo, now, now)
	if err != nil {
		return nil, err
	}

	return s.GetByDate(userID, req.Date)
}

func (s *DiaryService) GetByDate(userID, date string) (*model.Diary, error) {
	query := `
		SELECT id, user_id, date, rating, progress, wake_up_time, sleep_time, memo, created_at, updated_at
		FROM diaries
		WHERE user_id = ? AND date = ?
	`

	diary := &model.Diary{}
	err := s.db.QueryRow(query, userID, date).Scan(
		&diary.ID, &diary.UserID, &diary.Date, &diary.Rating, &diary.Progress,
		&diary.WakeUpTime, &diary.SleepTime, &diary.Memo, &diary.CreatedAt, &diary.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return diary, nil
}

func (s *DiaryService) GetAll(userID, startDate, endDate string, limit, offset int) ([]model.Diary, error) {
	query := `
		SELECT id, user_id, date, rating, progress, wake_up_time, sleep_time, memo, created_at, updated_at
		FROM diaries
		WHERE user_id = ?
	`
	args := []interface{}{userID}

	if startDate != "" {
		query += " AND date >= ?"
		args = append(args, startDate)
	}
	if endDate != "" {
		query += " AND date <= ?"
		args = append(args, endDate)
	}

	query += " ORDER BY date DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var diaries []model.Diary
	for rows.Next() {
		var d model.Diary
		err := rows.Scan(
			&d.ID, &d.UserID, &d.Date, &d.Rating, &d.Progress,
			&d.WakeUpTime, &d.SleepTime, &d.Memo, &d.CreatedAt, &d.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		diaries = append(diaries, d)
	}

	return diaries, nil
}

func (s *DiaryService) Update(userID, date string, req model.UpdateDiaryRequest) (*model.Diary, error) {
	existing, err := s.GetByDate(userID, date)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, sql.ErrNoRows
	}

	updates := []string{}
	args := []interface{}{}

	if req.Rating != nil {
		updates = append(updates, "rating = ?")
		args = append(args, *req.Rating)
	}
	if req.Progress != nil {
		updates = append(updates, "progress = ?")
		args = append(args, *req.Progress)
	}
	if req.WakeUpTime != nil {
		updates = append(updates, "wake_up_time = ?")
		args = append(args, *req.WakeUpTime)
	}
	if req.SleepTime != nil {
		updates = append(updates, "sleep_time = ?")
		args = append(args, *req.SleepTime)
	}
	if req.Memo != nil {
		updates = append(updates, "memo = ?")
		args = append(args, *req.Memo)
	}

	if len(updates) == 0 {
		return existing, nil
	}

	updates = append(updates, "updated_at = ?")
	args = append(args, time.Now())
	args = append(args, userID, date)

	query := fmt.Sprintf("UPDATE diaries SET %s WHERE user_id = ? AND date = ?",
		joinStrings(updates, ", "))

	_, err = s.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return s.GetByDate(userID, date)
}

func (s *DiaryService) Delete(userID, date string) error {
	query := "DELETE FROM diaries WHERE user_id = ? AND date = ?"
	_, err := s.db.Exec(query, userID, date)
	return err
}

func (s *DiaryService) GetCalendarData(userID string, year, month int) (*model.CalendarResponse, error) {
	startDate := fmt.Sprintf("%04d-%02d-01", year, month)
	endDate := fmt.Sprintf("%04d-%02d-31", year, month)

	query := `
		SELECT date, rating
		FROM diaries
		WHERE user_id = ? AND date >= ? AND date <= ?
		ORDER BY date
	`

	rows, err := s.db.Query(query, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []model.CalendarEntry
	var totalRating int
	for rows.Next() {
		var e model.CalendarEntry
		err := rows.Scan(&e.Date, &e.Rating)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
		totalRating += e.Rating
	}

	avgRating := 0.0
	if len(entries) > 0 {
		avgRating = float64(totalRating) / float64(len(entries))
	}

	// 月の日数を計算
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)
	totalDays := lastDay.Day()

	return &model.CalendarResponse{
		Year:    year,
		Month:   month,
		Entries: entries,
		Summary: model.CalendarSummary{
			TotalDays:     totalDays,
			RecordedDays:  len(entries),
			AverageRating: avgRating,
		},
	}, nil
}

func (s *DiaryService) GetStatistics(userID, period string) (*model.Statistics, error) {
	// 期間の計算
	var startDate, endDate string
	now := time.Now()

	switch period {
	case "week":
		startDate = now.AddDate(0, 0, -7).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "year":
		startDate = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
		endDate = time.Date(now.Year(), 12, 31, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
	default: // month
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
		endDate = time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -1).Format("2006-01-02")
		period = "month"
	}

	// 基本統計
	query := `
		SELECT
			COUNT(*) as total,
			AVG(rating) as avg_rating,
			SUM(CASE WHEN rating = 1 THEN 1 ELSE 0 END) as r1,
			SUM(CASE WHEN rating = 2 THEN 1 ELSE 0 END) as r2,
			SUM(CASE WHEN rating = 3 THEN 1 ELSE 0 END) as r3,
			SUM(CASE WHEN rating = 4 THEN 1 ELSE 0 END) as r4,
			SUM(CASE WHEN rating = 5 THEN 1 ELSE 0 END) as r5,
			SUM(CASE WHEN progress = 'A' THEN 1 ELSE 0 END) as pa,
			SUM(CASE WHEN progress = 'B' THEN 1 ELSE 0 END) as pb,
			SUM(CASE WHEN progress = 'C' THEN 1 ELSE 0 END) as pc
		FROM diaries
		WHERE user_id = ? AND date >= ? AND date <= ?
	`

	stats := &model.Statistics{
		Period:               period,
		PeriodStart:          startDate,
		PeriodEnd:            endDate,
		RatingDistribution:   make(map[string]int),
		ProgressDistribution: make(map[string]int),
	}

	var avgRating sql.NullFloat64
	var r1, r2, r3, r4, r5, pa, pb, pc int

	err := s.db.QueryRow(query, userID, startDate, endDate).Scan(
		&stats.TotalEntries, &avgRating,
		&r1, &r2, &r3, &r4, &r5,
		&pa, &pb, &pc,
	)
	if err != nil {
		return nil, err
	}

	if avgRating.Valid {
		stats.AverageRating = avgRating.Float64
	}

	stats.RatingDistribution = map[string]int{"1": r1, "2": r2, "3": r3, "4": r4, "5": r5}
	stats.ProgressDistribution = map[string]int{"A": pa, "B": pb, "C": pc}

	// 連続記録日数の計算
	stats.LongestStreak = s.calculateLongestStreak(userID)

	return stats, nil
}

func (s *DiaryService) GetTrend(userID string, days int) (*model.TrendData, error) {
	startDate := time.Now().AddDate(0, 0, -days).Format("2006-01-02")

	query := `
		SELECT date, rating
		FROM diaries
		WHERE user_id = ? AND date >= ?
		ORDER BY date
	`

	rows, err := s.db.Query(query, userID, startDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []model.TrendEntry
	for rows.Next() {
		var e model.TrendEntry
		err := rows.Scan(&e.Date, &e.Rating)
		if err != nil {
			return nil, err
		}
		data = append(data, e)
	}

	return &model.TrendData{
		PeriodDays: days,
		Data:       data,
	}, nil
}

func (s *DiaryService) calculateLongestStreak(userID string) int {
	query := `
		SELECT date FROM diaries
		WHERE user_id = ?
		ORDER BY date
	`

	rows, err := s.db.Query(query, userID)
	if err != nil {
		return 0
	}
	defer rows.Close()

	var dates []string
	for rows.Next() {
		var d string
		rows.Scan(&d)
		dates = append(dates, d)
	}

	if len(dates) == 0 {
		return 0
	}

	maxStreak := 1
	currentStreak := 1

	for i := 1; i < len(dates); i++ {
		prev, _ := time.Parse("2006-01-02", dates[i-1])
		curr, _ := time.Parse("2006-01-02", dates[i])

		if curr.Sub(prev).Hours() == 24 {
			currentStreak++
		} else {
			currentStreak = 1
		}

		if currentStreak > maxStreak {
			maxStreak = currentStreak
		}
	}

	return maxStreak
}

func joinStrings(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}
	return result
}

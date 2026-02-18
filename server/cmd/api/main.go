package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"github.com/nana743533/260219-diary-app/server/internal/config"
	"github.com/nana743533/260219-diary-app/server/internal/handler"
	"github.com/nana743533/260219-diary-app/server/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := sql.Open("mysql", cfg.Database.DSN())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// テーブル初期化
	if err := initDB(db); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	diaryService := service.NewDiaryService(db)

	diaryHandler := handler.NewDiaryHandler(diaryService)
	calendarHandler := handler.NewCalendarHandler(diaryService)
	statsHandler := handler.NewStatisticsHandler(diaryService)

	r := gin.Default()

	// ルーティング
	v1 := r.Group("/api/v1")
	{
		// 日記エンドポイント
		diaries := v1.Group("/diaries")
		{
			diaries.POST("", diaryHandler.Create)
			diaries.GET("", diaryHandler.GetAll)
			diaries.GET("/:date", diaryHandler.GetByDate)
			diaries.PUT("/:date", diaryHandler.Update)
			diaries.DELETE("/:date", diaryHandler.Delete)
		}

		// カレンダーエンドポイント
		calendar := v1.Group("/calendar")
		{
			calendar.GET("", calendarHandler.GetRange)
			calendar.GET("/:year/:month", calendarHandler.GetMonth)
		}

		// 統計エンドポイント
		stats := v1.Group("/statistics")
		{
			stats.GET("/summary", statsHandler.GetSummary)
			stats.GET("/trend", statsHandler.GetTrend)
		}
	}

	// ヘルスチェック
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"database": func() string {
				if err := db.Ping(); err != nil {
					return "error"
				}
				return "ok"
			}(),
		})
	})

	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	fmt.Printf("Server starting on %s\n", addr)
	log.Fatal(r.Run(addr))
}

func initDB(db *sql.DB) error {
	// ユーザーテーブル（認証なしの場合はダミー）
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(36) PRIMARY KEY,
			username VARCHAR(30) NOT NULL,
			email VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	// デフォルトユーザーを挿入
	_, err = db.Exec(`
		INSERT IGNORE INTO users (id, username, email)
		VALUES ('default-user', 'default', 'default@example.com')
	`)
	if err != nil {
		return err
	}

	// 日記テーブル
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS diaries (
			id VARCHAR(36) PRIMARY KEY,
			user_id VARCHAR(36) NOT NULL,
			date DATE NOT NULL,
			rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
			progress VARCHAR(1) NOT NULL CHECK (progress IN ('A', 'B', 'C')),
			wake_up_time VARCHAR(5) NOT NULL,
			sleep_time VARCHAR(5) NOT NULL,
			memo TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			UNIQUE KEY unique_user_date (user_id, date),
			INDEX idx_user_id (user_id),
			INDEX idx_date (date),
			INDEX idx_user_date (user_id, date)
		)
	`)

	return err
}

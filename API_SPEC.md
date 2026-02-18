# 日記アプリ API仕様書

## 概要

日記記録アプリのバックエンドAPI仕様です。ユーザーは日々の評価、進捗、起床/睡眠時間、メモを記録・閲覧できます。

**Base URL**: `https://api.diary-app.example.com`

**バージョン**: `v1`

---

## 目次

1. [認証](#認証)
2. [ユーザー管理](#ユーザー管理)
3. [日記エントリー](#日記エントリー)
4. [カレンダー](#カレンダー)
5. [統計](#統計)
6. [データモデル](#データモデル)
7. [エラーレスポンス](#エラーレスポンス)

---

## 認証

### JWTベースの認証

すべての保護されたエンドポイントは、Authorization ヘッダーにJWTトークンを含める必要があります。

```
Authorization: Bearer <jwt_token>
```

### トークンの有効期限

- アクセストークン: 24時間
- リフレッシュトークン: 30日

---

## ユーザー管理

### 1. ユーザー登録

**POST** `/api/v1/auth/register`

新規ユーザーを作成します。

**リクエスト**

```json
{
  "username": "string (3-30文字)",
  "email": "string (有効なメールアドレス)",
  "password": "string (8-72文字)"
}
```

**レスポンス** `201 Created`

```json
{
  "user": {
    "id": "uuid",
    "username": "string",
    "email": "string",
    "created_at": "RFC3339"
  },
  "token": {
    "access_token": "jwt_token",
    "refresh_token": "jwt_token",
    "expires_in": 86400
  }
}
```

### 2. ログイン

**POST** `/api/v1/auth/login`

**リクエスト**

```json
{
  "email": "string",
  "password": "string"
}
```

**レスポンス** `200 OK`

```json
{
  "user": {
    "id": "uuid",
    "username": "string",
    "email": "string"
  },
  "token": {
    "access_token": "jwt_token",
    "refresh_token": "jwt_token",
    "expires_in": 86400
  }
}
```

### 3. トークンリフレッシュ

**POST** `/api/v1/auth/refresh`

**リクエスト**

```json
{
  "refresh_token": "jwt_token"
}
```

**レスポンス** `200 OK`

```json
{
  "access_token": "jwt_token",
  "expires_in": 86400
}
```

### 4. ログアウト

**POST** `/api/v1/auth/logout`

**認証**: 必須

**レスポンス** `204 No Content`

---

## 日記エントリー

### 1. 日記を作成

**POST** `/api/v1/diaries`

**認証**: 必須

**リクエスト**

```json
{
  "date": "string (YYYY-MM-DD)",
  "rating": "integer (1-5)",
  "progress": "string (A|B|C)",
  "wake_up_time": "string (HH:MM)",
  "sleep_time": "string (HH:MM)",
  "memo": "string (任意)"
}
```

**レスポンス** `201 Created`

```json
{
  "id": "uuid",
  "user_id": "uuid",
  "date": "2025-02-19",
  "rating": 4,
  "progress": "A",
  "wake_up_time": "07:00",
  "sleep_time": "23:00",
  "memo": "今日は良い一日だった",
  "created_at": "2025-02-19T12:00:00Z",
  "updated_at": "2025-02-19T12:00:00Z"
}
```

### 2. 日記一覧取得

**GET** `/api/v1/diaries`

**認証**: 必須

**クエリパラメータ**

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|------|------|
| start_date | string | いいえ | 開始日 (YYYY-MM-DD) |
| end_date | string | いいえ | 終了日 (YYYY-MM-DD) |
| limit | integer | いいえ | 最大取得数 (デフォルト: 30) |
| offset | integer | いいえ | オフセット (デフォルト: 0) |

**リクエスト例**

```
GET /api/v1/diaries?start_date=2025-01-01&end_date=2025-01-31&limit=31
```

**レスポンス** `200 OK`

```json
{
  "diaries": [
    {
      "id": "uuid",
      "date": "2025-01-01",
      "rating": 4,
      "progress": "A",
      "wake_up_time": "07:00",
      "sleep_time": "23:00",
      "memo": "新年一発目の日記",
      "created_at": "2025-01-01T10:00:00Z",
      "updated_at": "2025-01-01T10:00:00Z"
    }
  ],
  "pagination": {
    "total": 31,
    "limit": 31,
    "offset": 0
  }
}
```

### 3. 特定の日記を取得

**GET** `/api/v1/diaries/{date}`

**認証**: 必須

**パスパラメータ**

- `date`: 日付 (YYYY-MM-DD)

**レスポンス** `200 OK`

```json
{
  "id": "uuid",
  "user_id": "uuid",
  "date": "2025-02-19",
  "rating": 4,
  "progress": "A",
  "wake_up_time": "07:00",
  "sleep_time": "23:00",
  "memo": "今日の出来事",
  "created_at": "2025-02-19T12:00:00Z",
  "updated_at": "2025-02-19T12:00:00Z"
}
```

### 4. 日記を更新

**PUT** `/api/v1/diaries/{date}`

**認証**: 必須

**リクエスト**

```json
{
  "rating": "integer (1-5, 任意)",
  "progress": "string (A|B|C, 任意)",
  "wake_up_time": "string (HH:MM, 任意)",
  "sleep_time": "string (HH:MM, 任意)",
  "memo": "string (任意)"
}
```

**レスポンス** `200 OK`

更新後の日記オブジェクト（形式は取得と同じ）

### 5. 日記を削除

**DELETE** `/api/v1/diaries/{date}`

**認証**: 必須

**レスポンス** `204 No Content`

---

## カレンダー

### 1. カレンダーヒートマップデータ取得

**GET** `/api/v1/calendar/{year}/{month}`

**認証**: 必須

**パスパラメータ**

- `year`: 年 (例: 2025)
- `month**: 月 (1-12)

**レスポンス** `200 OK`

```json
{
  "year": 2025,
  "month": 2,
  "entries": [
    {
      "date": "2025-02-01",
      "rating": 3
    },
    {
      "date": "2025-02-02",
      "rating": 4
    },
    {
      "date": "2025-02-19",
      "rating": 5
    }
  ],
  "summary": {
    "total_days": 28,
    "recorded_days": 19,
    "average_rating": 3.8
  }
}
```

### 2. 日付範囲のヒートマップデータ取得

**GET** `/api/v1/calendar`

**認証**: 必須

**クエリパラメータ**

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|------|------|
| start_date | string | はい | 開始日 (YYYY-MM-DD) |
| end_date | string | はい | 終了日 (YYYY-MM-DD) |

**レスポンス** `200 OK`

```json
{
  "start_date": "2025-01-01",
  "end_date": "2025-03-31",
  "entries": [
    {
      "date": "2025-01-01",
      "rating": 4,
      "progress": "A"
    }
  ]
}
```

---

## 統計

### 1. 統計サマリー取得

**GET** `/api/v1/statistics/summary`

**認証**: 必須

**クエリパラメータ**

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|------|------|
| period | string | いいえ | 期間 (week|month|year|all, デフォルト: month) |

**レスポンス** `200 OK`

```json
{
  "period": "month",
  "period_start": "2025-02-01",
  "period_end": "2025-02-28",
  "total_entries": 19,
  "average_rating": 3.8,
  "rating_distribution": {
    "1": 0,
    "2": 2,
    "3": 5,
    "4": 8,
    "5": 4
  },
  "progress_distribution": {
    "A": 8,
    "B": 7,
    "C": 4
  },
  "average_wake_up_time": "07:15",
  "average_sleep_time": "23:30",
  "longest_streak": 7
}
```

### 2. 評価の推移（トレンド）

**GET** `/api/v1/statistics/trend`

**認証**: 必須

**クエリパラメータ**

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|------|------|
| days | integer | いいえ | 日数 (デフォルト: 30) |

**レスポンス** `200 OK`

```json
{
  "period_days": 30,
  "data": [
    {
      "date": "2025-01-20",
      "rating": 3
    },
    {
      "date": "2025-01-21",
      "rating": 4
    }
  ]
}
```

---

## データモデル

### User

```go
type User struct {
    ID        string    `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Diary

```go
type Diary struct {
    ID          string    `json:"id"`
    UserID      string    `json:"user_id"`
    Date        string    `json:"date"`          // YYYY-MM-DD
    Rating      int       `json:"rating"`        // 1-5
    Progress    string    `json:"progress"`      // A, B, C
    WakeUpTime  string    `json:"wake_up_time"`  // HH:MM
    SleepTime   string    `json:"sleep_time"`    // HH:MM
    Memo        string    `json:"memo"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

### CalendarEntry

```go
type CalendarEntry struct {
    Date   string `json:"date"`
    Rating int    `json:"rating"`
}
```

### Statistics

```go
type Statistics struct {
    Period             string            `json:"period"`
    PeriodStart        string            `json:"period_start"`
    PeriodEnd          string            `json:"period_end"`
    TotalEntries       int               `json:"total_entries"`
    AverageRating      float64           `json:"average_rating"`
    RatingDistribution map[string]int    `json:"rating_distribution"`
    ProgressDistribution map[string]int  `json:"progress_distribution"`
    AverageWakeUpTime  string            `json:"average_wake_up_time"`
    AverageSleepTime   string            `json:"average_sleep_time"`
    LongestStreak      int               `json:"longest_streak"`
}
```

---

## エラーレスポンス

すべてのエラーは以下の形式で返されます：

```json
{
  "error": {
    "code": "string",
    "message": "string",
    "details": {} // 任意
  }
}
```

### エラーコード一覧

| コード | HTTPステータス | 説明 |
|-------|---------------|------|
| `UNAUTHORIZED` | 401 | 認証が必要 |
| `INVALID_TOKEN` | 401 | トークンが無効 |
| `FORBIDDEN` | 403 | アクセス権限がない |
| `NOT_FOUND` | 404 | リソースが見つからない |
| `VALIDATION_ERROR` | 400 | バリデーションエラー |
| `DUPLICATE_ENTRY` | 409 | データが既に存在 |
| `INTERNAL_ERROR` | 500 | サーバーエラー |

### エラーレスポンス例

**400 Bad Request** - バリデーションエラー

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "バリデーションに失敗しました",
    "details": {
      "rating": "1-5の範囲で指定してください",
      "date": "YYYY-MM-DD形式で指定してください"
    }
  }
}
```

**401 Unauthorized**

```json
{
  "error": {
    "code": "INVALID_TOKEN",
    "message": "トークンが無効です"
  }
}
```

**404 Not Found**

```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "指定の日記が見つかりません"
  }
}
```

---

## Go実装推奨ライブラリ

```go
// APIフレームワーク
github.com/gin-gonic/gin

// データベース
github.com/lib/pq          // PostgreSQL
github.com/mattn/go-sqlite3 // SQLite

// 認証
github.com/golang-jwt/jwt/v5

// バリデーション
github.com/go-playground/validator/v10

// 設定管理
github.com/spf13/viper
```

---

## データベーススキーマ（PostgreSQL）

```sql
-- ユーザーテーブル
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(30) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 日記テーブル
CREATE TABLE diaries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    rating INTEGER CHECK (rating >= 1 AND rating <= 5) NOT NULL,
    progress VARCHAR(1) CHECK (progress IN ('A', 'B', 'C')) NOT NULL,
    wake_up_time TIME NOT NULL,
    sleep_time TIME NOT NULL,
    memo TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, date)
);

-- インデックス
CREATE INDEX idx_diaries_user_id ON diaries(user_id);
CREATE INDEX idx_diaries_date ON diaries(date);
CREATE INDEX idx_diaries_user_date ON diaries(user_id, date);
```

---

## 開発環境

**Goバージョン**: 1.22+

**推奨ポート**: `8080`

**環境変数**

| 変数 | 説明 | デフォルト |
|-----|------|----------|
| `PORT` | サービスポート | 8080 |
| `DB_HOST` | DBホスト | localhost |
| `DB_PORT` | DBポート | 5432 |
| `DB_NAME` | DB名 | diary_app |
| `DB_USER` | DBユーザー | postgres |
| `DB_PASSWORD` | DBパスワード | - |
| `JWT_SECRET` | JWTシークレット | - |
| `JWT_EXPIRES_IN` | トークン有効期限 | 24h |

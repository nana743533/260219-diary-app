# Diary API Server

Go + Gin + MySQL で実装された日記アプリのバックエンドAPIサーバーです。

## 機能

- 日記のCRUD操作
- カレンダーヒートマップデータ
- 統計サマリー・トレンド

## ディレクトリ構成

```
server/
├── cmd/api/
│   └── main.go           # エントリーポイント
├── internal/
│   ├── config/           # 設定管理
│   ├── handler/          # HTTPハンドラー
│   ├── model/            # データモデル
│   └── service/          # ビジネスロジック
├── .env.example          # 環境変数サンプル
├── Makefile             # 開発コマンド
└── go.mod
```

## セットアップ

### 1. MySQLを起動

Dockerを使用する場合：

```bash
cd server
make mysql-up
```

手動でMySQLを起動する場合：

```bash
mysql -u root -p
CREATE DATABASE diary_app;
```

### 2. 環境変数設定

```bash
cp .env.example .env
# .envを編集
```

### 3. 依存関係インストール

```bash
go mod download
```

### 4. サーバー起動

```bash
make run
```

サーバーが `http://localhost:8080` で起動します。

## APIエンドポイント

### 日記

| メソッド | パス | 説明 |
|---------|------|------|
| POST | `/api/v1/diaries` | 日記作成 |
| GET | `/api/v1/diaries` | 日記一覧 |
| GET | `/api/v1/diaries/:date` | 日記取得 |
| PUT | `/api/v1/diaries/:date` | 日記更新 |
| DELETE | `/api/v1/diaries/:date` | 日記削除 |

### カレンダー

| メソッド | パス | 説明 |
|---------|------|------|
| GET | `/api/v1/calendar/:year/:month` | 月別データ |
| GET | `/api/v1/calendar?start_date=X&end_date=Y` | 期間指定 |

### 統計

| メソッド | パス | 説明 |
|---------|------|------|
| GET | `/api/v1/statistics/summary?period=month` | サマリー |
| GET | `/api/v1/statistics/trend?days=30` | トレンド |

## API使用例

### 日記を作成

```bash
curl -X POST http://localhost:8080/api/v1/diaries \
  -H "Content-Type: application/json" \
  -d '{
    "date": "2025-02-19",
    "rating": 4,
    "progress": "A",
    "wake_up_time": "07:00",
    "sleep_time": "23:00",
    "memo": "良い一日だった"
  }'
```

### 日記を取得

```bash
curl http://localhost:8080/api/v1/diaries/2025-02-19
```

### カレンダーデータ取得

```bash
curl http://localhost:8080/api/v1/calendar/2025/2
```

### 統計取得

```bash
curl http://localhost:8080/api/v1/statistics/summary?period=month
```

## 開発コマンド

```bash
make run        # サーバー起動
make build      # ビルド
make test       # テスト
make fmt        # フォーマット
make tidy       # go mod tidy
```

## 注意事項

- 現在は認証なしで実装されています（`default-user`が固定で使用されます）
- 本番環境では認証機能を実装してください

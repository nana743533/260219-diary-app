.PHONY: run build clean test

# デフォルトターゲット
run:
	go run cmd/api/main.go

# ビルド
build:
	go build -o bin/api cmd/api/main.go

# クリーン
clean:
	rm -rf bin/

# テスト
test:
	go test ./...

# モジュールタidy
tidy:
	go mod tidy

# 依存関係ダウンロード
deps:
	go mod download

# フォーマット
fmt:
	go fmt ./...

# リンター
lint:
	golangci-lint run

# 実行（環境変数指定）
run-dev:
	cp .env.example .env
	SERVER_PORT=8080 \
	DB_HOST=localhost \
	DB_PORT=3306 \
	DB_USER=root \
	DB_PASSWORD= \
	DB_NAME=diary_app \
	go run cmd/api/main.go

# Docker MySQL起動
mysql-up:
	docker run -d --name diary-mysql \
		-e MYSQL_ROOT_PASSWORD=root \
		-e MYSQL_DATABASE=diary_app \
		-p 3306:3306 \
		mysql:8

# MySQL停止
mysql-down:
	docker stop diary-mysql && docker rm diary-mysql

# MySQLログ
mysql-logs:
	docker logs -f diary-mysql

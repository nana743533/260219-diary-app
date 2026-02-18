# 日記アプリ (Diary App)

1日の総合評価・進捗・起床時間・睡眠時間・メモを記録できるスマホ日記アプリです。

GitHubのようなカレンダーヒートマップで過去の記録を可視化できます。

![Platform](https://img.shields.io/badge/platform-React%20Native-blue)
![Backend](https://img.shields.io/badge/backend-Go%20%2B%20MySQL-00ADD8)

## 機能

### スマホアプリ (React Native)

- 📝 日記記録（評価1-5、進捗A/B/C、時間、メモ）
- 📅 カレンダーヒートマップ表示（GitHub風）
- 💾 ローカルデータ保存（AsyncStorage）
- 🔄 タブ切り替え（カレンダー/記録）

### バックエンドAPI (Go)

- 📊 RESTful API
- 🗄️ MySQLデータベース
- 📈 統計機能（サマリー、トレンド）
- 🗓️ カレンダーデータ取得

## スクリーンショット

### 記録画面
評価ボタン（1-5）、進捗選択（A/B/C）、時間入力、メモ入力

### カレンダー画面
評価が緑色の濃さで表示されるヒートマップ

## 技術スタック

### フロントエンド

- React Native 0.81.5
- Expo 54
- React Navigation
- react-native-calendars
- AsyncStorage

### バックエンド

- Go 1.22+
- Gin Web Framework
- MySQL 8
- go-sql-driver/mysql

## ディレクトリ構成

```
260219-diary-app/
├── App.js                    # React Nativeエントリーポイント
├── src/
│   ├── navigation/           # タブナビゲーション
│   ├── screens/              # 画面
│   ├── components/           # コンポーネント
│   ├── services/             # データ永続化
│   └── utils/                # ユーティリティ
├── server/                   # Goバックエンド
│   ├── cmd/api/              # エントリーポイント
│   ├── internal/
│   │   ├── handler/          # HTTPハンドラー
│   │   ├── service/          # ビジネスロジック
│   │   ├── model/            # データモデル
│   │   └── config/           # 設定
│   └── Makefile              # 開発コマンド
└── API_SPEC.md               # API仕様書
```

## セットアップ

### フロントエンド

```bash
# 依存関係インストール
npm install

# 起動
npx expo start
```

Expo GoアプリでQRコードをスキャンしてください。

### バックエンド

```bash
cd server

# MySQL起動
make mysql-up

# サーバー起動
make run
```

詳細は [server/README.md](./server/README.md) を参照してください。

## API仕様

API仕様書は [API_SPEC.md](./API_SPEC.md) を参照してください。

## 使用例

### 日記を保存

1. 「記録」タブを開く
2. 評価（1-5）を選択
3. 進捗（A/B/C）を選択
4. 起床時間・睡眠時間を入力
5. メモを入力（任意）
6. 「保存」をタップ

### カレンダーで確認

1. 「カレンダー」タブを開く
2. 評価がある日は緑色で表示
3. 日付をタップして詳細確認

## ヒートマップの色

| 評価 | 色 |
|------|------|
| なし | グレー |
| 1 | 薄い緑 |
| 2 | 中緑 |
| 3 | 濃い緑 |
| 4 | かなり濃い緑 |
| 5 | 最も濃い緑 |

## 開発

```bash
# フロントエンド
npm install           # 依存関係インストール
npx expo start        # 開発サーバー起動

# バックエンド
cd server
make build            # ビルド
make test             # テスト
make fmt              # フォーマット
```

## 今後の機能

- [ ] ユーザー認証
- [ ] データ同期（ローカル↔サーバー）
- [ ] グラフ表示
- [ ] エクスポート機能
- [ ] ダークモード

## ライセンス

MIT

## 作者

[nana743533](https://github.com/nana743533)

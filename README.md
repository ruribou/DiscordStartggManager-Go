# Discord StartGG Manager Bot by Golang

Discord上でStart.ggの大会情報を管理するBotです。Golangで作成されています。

## 必要なもの

- Go 1.25.0
- Docker Desktop
- Discord Bot Token

### 1. 環境変数の設定

`.env.sample`ファイルをコピーして`.env`ファイルを作成して以下の環境変数を設定してください：

```env
# Discord Bot設定
DISCORD_BOT_TOKEN=your_discord_bot_token_here

# Start.gg API設定（今後の拡張用）
STARTGG_API_TOKEN=your_startgg_api_token_here
```

### 2. 実行方法

#### ローカル実行
```bash
go mod tidy
go run main.go
```

#### Docker実行（推奨）
```bash
docker compose up --build
```

## 使用可能なコマンド

鋭意制作中。

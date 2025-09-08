package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルから環境変数を読み込み
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	// DiscordのBotトークンを取得
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_BOT_TOKEN environment variable is required")
	}

	// Discordセッションを作成
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	// メッセージハンドラーを追加
	dg.AddHandler(messageCreate)

	// Intentを設定（メッセージを受信するために必要）
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// WebSocket接続を開始
	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
	}

	fmt.Println("Discord Bot is running. Press CTRL+C to exit.")

	// プログラムが終了するまで待機
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// WebSocket接続を閉じる
	dg.Close()
}

// メッセージが作成されたときの処理
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ボット自身のメッセージは無視
	if m.Author.ID == s.State.User.ID {
		return
	}

	// "!ping" コマンドに応答
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// "!hello" コマンドに応答
	if m.Content == "!hello" {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hello, %s!", m.Author.Username))
	}
}

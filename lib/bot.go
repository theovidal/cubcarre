package lib

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Handler func(bot *Bot, update *telegram.Update, chatID int64, args []string) error

// Bot is a structure that holds the Telegram API and other assets and is unique
type Bot struct {
	// The Telegram API is merged into Bot structure
	*telegram.BotAPI
	// Commands associates command's name to its function
	Commands map[string]Handler
	// Commands associates callback's name to its function
	Callbacks map[string]Handler
	// Db is the SQL database
	Db *gorm.DB
}

// Error sends a formatted error message in the Telegram chat
func (bot *Bot) Error(chatID int64, message string) (err error) {
	msg := telegram.NewMessage(chatID, "❌ "+message)
	msg.ParseMode = "Markdown"
	_, err = bot.Send(msg)
	return
}

// LoadEnv loads all the environment variables stored in a .env file
func LoadEnv(path string) {
	_ = godotenv.Load(path)
}

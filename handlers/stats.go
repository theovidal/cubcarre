package handlers

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/theovidal/cubcarre/lib"
)

func StatsCommand(bot *lib.Bot, update *telegram.Update, chatID int64, args []string) error {
	return nil
}
package handlers

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/theovidal/cubcarre/lib"
)

func SelectCubeCallback(bot *lib.Bot, update *telegram.Update, userID, chatID int64, args []string) error {
	return nil
}

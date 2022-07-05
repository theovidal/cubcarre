package handlers

import (
	"fmt"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/theovidal/cubcarre/lib"
)

var DefaultCubes = []string{"2x2", "3x3", "4x4", "5x5", "Sq-1", "Megamx", "Pyramx"}

func StartCommand(bot *lib.Bot, update *telegram.Update, chatID int64, _ []string) (err error) {
	msg := telegram.NewMessage(chatID, MainMenuText)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = MainMenuKeyboard(update.Message.From.ID)
	_, err = bot.Send(msg)
	return
}

func MainMenuCallback(bot *lib.Bot, update *telegram.Update, userID, chatID int64, _ []string) (err error) {
	_, err = bot.Send(
		telegram.NewEditMessageTextAndMarkup(chatID, update.CallbackQuery.Message.MessageID, MainMenuText, MainMenuKeyboard(userID)),
	)
	return
}

var MainMenuText = "*―――――― 🎲 CubCarré ――――――*\nSelect an option below, or open the timer via the according button."

func MainMenuKeyboard(userID int64) telegram.InlineKeyboardMarkup {
	return telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			{
				telegram.NewInlineKeyboardButtonData("▶ Select cube", fmt.Sprintf("selectcube:%d", userID)),
				telegram.NewInlineKeyboardButtonData("📈 Statistics", fmt.Sprintf("statistics:%d", userID)),
				telegram.NewInlineKeyboardButtonData("⚙ Settings", fmt.Sprintf("settings:%d", userID)),
			},
		},
	}
}

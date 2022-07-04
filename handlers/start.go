package handlers

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/theovidal/cubcarre/lib"
)

var DefaultCubes = []string{"2x2", "3x3", "4x4", "5x5", "Sq-1", "Megamx", "Pyramx"}

func StartCommand(bot *lib.Bot, _ *telegram.Update, chatID int64, args []string) (err error) {
	msg := telegram.NewMessage(chatID, MainMenuText)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = &MainMenuKeyboard
	_, err = bot.Send(msg)
	return
}

func MainMenuCallback(bot *lib.Bot, update *telegram.Update, chatID int64, args []string) (err error) {
	_, err = bot.Send(
		telegram.NewEditMessageTextAndMarkup(chatID, update.CallbackQuery.Message.MessageID, MainMenuText, MainMenuKeyboard),
	)
	return
}

var MainMenuText = "*―――――― 🎲 CubCarré ――――――*\nSelect an option below, or open the timer via the according button."
var MainMenuKeyboard = telegram.InlineKeyboardMarkup{
	InlineKeyboard: [][]telegram.InlineKeyboardButton{
		{
			telegram.NewInlineKeyboardButtonData("▶ Select cube", "selectcube"),
			telegram.NewInlineKeyboardButtonData("📈 Statistics", "statistics"),
			telegram.NewInlineKeyboardButtonData("⚙ Settings", "settings"),
		},
	},
}

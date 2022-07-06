package handlers

import (
	"fmt"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/theovidal/cubcarre/lib"
)

func StartCommand(bot *lib.Bot, update *telegram.Update, chatID int64, _ []string) (err error) {
	msg := telegram.NewMessage(chatID, MainMenuText)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = MainMenuKeyboard(update.Message.From.ID)
	_, err = bot.Send(msg)
	return
}

func MainMenuCallback(bot *lib.Bot, update *telegram.Update, userID, chatID int64, _ []string) (err error) {
	markup := telegram.NewEditMessageTextAndMarkup(chatID, update.CallbackQuery.Message.MessageID, MainMenuText, MainMenuKeyboard(userID))
	markup.ParseMode = "Markdown"
	_, err = bot.Send(markup)
	return
}

var MainMenuText = "*―――――― 🎲 CubCarré ――――――*\nSelect an option below, or open the timer via the according button."

func MainMenuKeyboard(userID int64) telegram.InlineKeyboardMarkup {
	return telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			{
				telegram.NewInlineKeyboardButtonData("▶ Select puzzle", fmt.Sprintf("selectpuzzle:%d", userID)),
				telegram.NewInlineKeyboardButtonData("📈 Statistics", fmt.Sprintf("statistics:%d", userID)),
				telegram.NewInlineKeyboardButtonData("⚙ Settings", fmt.Sprintf("settings:%d", userID)),
			},
		},
	}
}

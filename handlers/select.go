package handlers

import (
	"fmt"
	"os"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/theovidal/cubcarre/lib"
)

var cubes = [][]string{
	{"2x2", "3x3", "4x4", "5x5", "6x6", "7x7"},
	{"3BLD", "4BLD", "Pyrmx", "Megamx", "Sq-1", "Other"},
}

func contains(el string) bool {
	for _, row := range cubes {
		for _, cube := range row {
			if cube == el {
				return true
			}
		}
	}
	return false
}

func GetCubesKeyboard(userID int64) (buttons [][]telegram.InlineKeyboardButton) {
	for _, row := range cubes {
		var rowBtns []telegram.InlineKeyboardButton
		for _, cube := range row {
			data := fmt.Sprintf("selectionmade:%d:%s", userID, cube)
			rowBtns = append(rowBtns, telegram.InlineKeyboardButton{
				Text:         cube,
				CallbackData: &data,
			})
		}
		buttons = append(buttons, rowBtns)
	}
	return
}

func SelectPuzzleCallback(bot *lib.Bot, update *telegram.Update, userID, chatID int64, _ []string) error {
	_, err := bot.Send(telegram.NewEditMessageTextAndMarkup(
		chatID,
		update.CallbackQuery.Message.MessageID,
		"Select the puzzle you want to be timed on",
		telegram.InlineKeyboardMarkup{
			InlineKeyboard: GetCubesKeyboard(userID),
		}))
	return err
}

func SelectionMadeCallback(bot *lib.Bot, update *telegram.Update, userID, chatID int64, args []string) (err error) {
	if !contains(args[0]) {
		return bot.Error(chatID, "Invalid cube type (well tried, but you can't choose whatever you want!)")
	}
	_, err = bot.Send(telegram.SetChatMenuButtonConfig{
		ChatID: chatID,
		MenuButton: &telegram.MenuButton{
			Type: "web_app",
			Text: fmt.Sprintf("⌚ %s", args[0]),
			WebApp: &telegram.WebAppInfo{
				URL: fmt.Sprintf("%s/?cube=%s", os.Getenv("WEBAPP_URL"), args[0]),
			},
		},
	})
	_, err = bot.Send(telegram.CallbackConfig{
		CallbackQueryID: update.CallbackQuery.ID,
		Text:            fmt.Sprintf("You will now be timed on puzzle %s", args[0]),
	})
	return MainMenuCallback(bot, update, userID, chatID, []string{})
}

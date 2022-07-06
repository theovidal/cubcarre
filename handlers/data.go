package handlers

import (
	"fmt"
	"strconv"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/theovidal/cubcarre/db"
	"github.com/theovidal/cubcarre/lib"
)

func SaveTime(bot *lib.Bot, queryID string, userID int64, userTime uint64, cube string) (err error) {
	// TODO: data validation
	res := db.Time{
		User:  userID,
		Cube:  cube,
		Value: userTime,
	}
	if err = bot.Db.Create(&res).Error; err != nil {
		return
	}
	_, err = bot.AnswerWebAppQuery(telegram.AnswerWebAppQueryConfig{
		WebAppQueryID: queryID,
		Result: telegram.InlineQueryResultArticle{
			Type:  "article",
			Title: "New time saved",
			ID:    strconv.Itoa(int(time.Now().Unix())),
			InputMessageContent: telegram.InputTextMessageContent{
				Text: fmt.Sprintf("⌚ %s on puzzle %s", lib.GenerateTimeDisplay(userTime), cube),
			},
			ReplyMarkup: &telegram.InlineKeyboardMarkup{
				InlineKeyboard: [][]telegram.InlineKeyboardButton{
					{
						telegram.NewInlineKeyboardButtonData("+2", fmt.Sprintf("mark:0:p_two:%d", res.ID)),
						telegram.NewInlineKeyboardButtonData("DNF", fmt.Sprintf("mark:0:dnf:%d", res.ID)),
					},
				},
			},
		},
	})
	return err
}

var prettyDisplay = map[string]string{
	"p_two": "+2",
	"dnf":   "DNF",
}

func MarkCallback(bot *lib.Bot, update *telegram.Update, _, _ int64, args []string) (err error) {
	id, err := strconv.ParseUint(args[1], 10, 0)
	if err != nil {
		return
	}

	if args[0] == "p_two" || args[0] == "dnf" {
		bot.Db.Model(&db.Time{ID: uint(id)}).Updates(map[string]interface{}{
			args[0]: 1,
		})
		err = bot.Db.Error
		if err != nil {
			return
		}
		_, err = bot.Send(telegram.CallbackConfig{
			CallbackQueryID: update.CallbackQuery.ID,
			Text:            fmt.Sprintf("This time is now marked as %s", prettyDisplay[args[0]]),
		})
	}

	return
}

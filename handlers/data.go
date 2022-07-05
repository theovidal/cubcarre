package handlers

import (
	"fmt"
	"strconv"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/theovidal/cubcarre/db"
	"github.com/theovidal/cubcarre/lib"
)

func SaveTime(bot *lib.Bot, queryID string, userID int64, userTime uint64, cube string) error {
	bot.Db.Create(&db.Time{
		User:  userID,
		Cube:  cube,
		Value: userTime,
	})
	_, err := bot.AnswerWebAppQuery(telegram.AnswerWebAppQueryConfig{
		WebAppQueryID: queryID,
		Result: telegram.InlineQueryResultArticle{
			Type:  "article",
			Title: "New time saved",
			ID:    strconv.Itoa(int(time.Now().Unix())),
			InputMessageContent: telegram.InputTextMessageContent{
				Text: fmt.Sprintf("⌚ %s on puzzle %s", lib.GenerateTimeDisplay(userTime), cube),
			},
		},
	})
	return err
}

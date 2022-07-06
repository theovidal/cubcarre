package handlers

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/theovidal/cubcarre/db"
	"github.com/theovidal/cubcarre/lib"
)

func StatsCallback(bot *lib.Bot, update *telegram.Update, userID, chatID int64, _ []string) (err error) {
	var stats []db.Time
	err = bot.Db.Where("user = ?", userID).Find(&stats).Error
	if err != nil {
		lib.LogError("error fetching stats: %v", err)
		return
	}
	if len(stats) == 0 {
		_, err = bot.Send(telegram.CallbackConfig{
			Text:            "🍃 You don't have any stats for the moment. Try making a solve using the timer with the button below 🔽",
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       true,
		})
		return
	}

	plot := lib.MakePlot(stats)
	photo := telegram.NewPhoto(chatID, lib.WriteImage(plot))
	_, err = bot.Send(photo)
	_, _ = bot.Send(telegram.CallbackConfig{
		Text:            "Your statistics were sent into the chat!",
		CallbackQueryID: update.CallbackQuery.ID,
	})
	return
}

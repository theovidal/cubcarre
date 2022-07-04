package main

import (
	"log"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/theovidal/cubcarre/db"
	"github.com/theovidal/cubcarre/handlers"
	"github.com/theovidal/cubcarre/lib"
)

func main() {
	lib.LoadEnv(".env")

	api, err := telegram.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot := &lib.Bot{
		BotAPI: api,
		Commands: map[string]lib.Handler{
			"/start":      handlers.StartCommand,
			"/statistics": handlers.StatsCommand,
			"/settings":   handlers.SettingsCommand,
			"/selectcube": handlers.SelectCubeCommand,
		},
		Callbacks: map[string]lib.Handler{
			"menu": handlers.MainMenuCallback,
		},
		Db: db.OpenDatabase(),
	}

	if os.Getenv("DEBUG") == "true" {
		bot.Debug = true
		lib.StandardLogger.Debug = true
		lib.LogInfo("Debug mode activated - check .env to disable")
	}

	lib.LogSuccess("Authorized on account %s", bot.Self.UserName)

	updateChannel := telegram.NewUpdate(0)
	updateChannel.Timeout = math.MaxInt

	updates := bot.GetUpdatesChan(updateChannel)

	go func() {
		for update := range updates {
			println("update")
			if update.CallbackQuery != nil {
				err = HandleCallback(bot, update)
				if err != nil {
					lib.LogError("Error handling a callback: %s", err)
				}
			} else if update.Message != nil {
				if update.Message.IsCommand() {
					if update.Message.From.UserName != os.Getenv("TELEGRAM_USER") {
						continue
					}
					if err = HandleCommand(bot, update); err != nil {
						lib.LogError("Error handling a command: %s", err)
					}
				} else if update.Message.WebAppData != nil {
					err = handlers.SaveTime(bot, &update, update.Message.Chat.ID, strings.Split(update.Message.WebAppData.Data, ":"))
					if err != nil {
						lib.LogError("Error handling webapp response: %s", err)
					}
				}
			}
		}
	}()

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	lib.LogInfo("Gracefully shutting down bot 💤")
}

package main

import (
	"strconv"
	"strings"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/theovidal/cubcarre/lib"
)

func HandleCommand(bot *lib.Bot, update telegram.Update) error {
	args := strings.Split(update.Message.Text, " ")

	cmd, exists := bot.Commands[args[0]]
	if !exists {
		return bot.Error(update.Message.Chat.ID, "La commande souhaitée n'existe pas (ou n'a pas encore été implémentée).")
	} else {
		return cmd(bot, &update, update.Message.Chat.ID, args[1:])
	}
}

func HandleCallback(bot *lib.Bot, update telegram.Update) error {
	// Starts with a / ; treat as a command
	if update.CallbackQuery.Data[0] == '/' {
		args := strings.Split(update.CallbackQuery.Data, " ")
		cmd, _ := bot.Commands[args[0][1:]]
		return cmd(bot, &update, update.CallbackQuery.Message.Chat.ID, args[1:])

	} else {
		args := strings.Split(update.CallbackQuery.Data, ":")
		callback, _ := bot.Callbacks[args[0]]
		userID, err := strconv.ParseInt(args[1], 10, 0)
		if err != nil {
			return err
		}
		return callback(bot, &update, userID, update.CallbackQuery.Message.Chat.ID, args[1:])
	}
}

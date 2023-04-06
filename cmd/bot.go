package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"psa_dump_bot/bot"
)

func StartBot(e *bot.Environment) {

	bot, err := tgbotapi.NewBotAPI(e.Config.Token)
	bot.CheckFatalError(err)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		var msg tgbotapi.MessageConfig
		if bot.validateUser(bot, &update, e.Config) != nil {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, bot.CreateErrAuthMsg(e))
			_, err := bot.Send(msg)
			bot.CheckFatalError(err)
			return
		}

		if update.CallbackQuery != nil {
			msg = bot.CallbackProcessing(&update, e)
		}
		if update.Message != nil {
			msg = bot.MsgProcessing(&update, e)
		}

		_, err := bot.Send(msg)
		bot.CheckError(err)
	}
}

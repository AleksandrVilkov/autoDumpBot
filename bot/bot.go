package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func StartBot(e *Environment) {

	bot, err := tgbotapi.NewBotAPI(e.Config.Token)
	CheckFatalError(err)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		var msg tgbotapi.MessageConfig
		if validateUser(bot, &update, e.Config) != nil {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, CreateErrAuthMsg(e))
			_, err := bot.Send(msg)
			CheckFatalError(err)
			return
		}

		if update.CallbackQuery != nil {

			msg = CallbackProcessing(&update, e)
		}
		if update.Message != nil {
			msg = MsgProcessing(&update, e)
		}

		_, err := bot.Send(msg)
		CheckError(err)
	}
}

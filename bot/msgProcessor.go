package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"reflect"
)

func MsgProcessing(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {

	var msg tgbotapi.MessageConfig

	if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

		switch update.Message.Text {
		case e.Config.Commands.Start:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, CreateWelcomeMsg(e))
			msg.ReplyMarkup = CreateMainButtons(e)
		default:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, e.Resources.Errors.CommonError)
		}
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Use the words for search.")

	}

	return msg
}

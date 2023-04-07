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
			msg.ReplyMarkup = CreateMainButtons(e, update)
		default:
			msg = CreateErrorMsg(update, e)
		}
	} else {
		msg = CreateErrorMsg(update, e)
	}

	return msg
}

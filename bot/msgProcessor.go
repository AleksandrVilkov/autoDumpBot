package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"psa_dump_bot/bot/model"
	"reflect"
	"strconv"
)

func MsgProcessing(update *tgbotapi.Update, e *model.Environment) tgbotapi.MessageConfig {

	var msg tgbotapi.MessageConfig

	if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

		switch update.Message.Text {
		case e.Config.Commands.Start:
			isRegistered := e.Storage.CheckUser(strconv.FormatInt(update.Message.Chat.ID, 10))

			if isRegistered {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, CreateWelcomeRegisteredMsg(e))
				msg.ReplyMarkup = CreateMainButtons(e, update)
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, CreateWelcomeMsg(e))
				msg.ReplyMarkup = CreateRegistrationButton(e, update)
			}

		default:
			msg = CreateErrorMsg(update, e)
		}
	} else {
		msg = CreateErrorMsg(update, e)
	}

	return msg
}

package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"reflect"
	"strconv"
)

func MsgProcessing(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {

	var msg tgbotapi.MessageConfig
	var tempData TempUserData

	if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

		tempData.User.Id = strconv.FormatInt(update.Message.Chat.ID, 10)
		switch update.Message.Text {
		case e.Config.Commands.Start:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, CreateWelcomeMsg(e))
			msg.ReplyMarkup = CreateMainButtons(e)
			tempData.Action.LastCommand = e.Config.Commands.Start
		default:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, e.Resources.Errors.CommonError)
		}
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Use the words for search.")

	}

	e.TempData[strconv.FormatInt(update.Message.Chat.ID, 10)] = tempData
	return msg
}

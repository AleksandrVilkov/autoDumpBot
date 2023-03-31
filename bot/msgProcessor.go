package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"reflect"
	"strconv"
)

func MsgProcessing(update *tgbotapi.Update, conf Config, temp map[string]TempUserData) tgbotapi.MessageConfig {

	var msg tgbotapi.MessageConfig
	var tempData TempUserData

	if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

		tempData.User.Id = strconv.FormatInt(update.Message.Chat.ID, 10)
		switch update.Message.Text {
		case conf.Commands.Start:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, CreateWelcomeMsg())
			msg.ReplyMarkup = CreateMainButtons()
			tempData.Action.LastCommand = conf.Commands.Start
		default:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, GetResources().Errors.CommonError)
		}
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Use the words for search.")

	}

	temp[strconv.FormatInt(update.Message.Chat.ID, 10)] = tempData
	return msg
}

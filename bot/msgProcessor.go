package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"reflect"
	"strconv"
)

func MsgProcessing(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {

	var msg tgbotapi.MessageConfig

	if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

		switch update.Message.Text {
		case e.Config.Commands.Start:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, CreateWelcomeMsg(e))
			msg.ReplyMarkup = CreateMainButtons(e)
		default:
			nonStandardMsgProcess(update, e)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, e.Resources.Errors.CommonError)
		}
	} else {
		msg = nonStandardMsgProcess(update, e)
	}

	return msg
}
func nonStandardMsgProcess(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
	//	lastCommand := userTemp.LastCommand
	var msg tgbotapi.MessageConfig
	switch userTemp.LastCommand {
	case PLACE_AN_AD:
		msg = createPlaceAnAdWelcomeResp(update, e)
	}
	return msg

}

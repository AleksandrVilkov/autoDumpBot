package bot

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func CallbackProcessing(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	callback, err := getCallBack(update.CallbackQuery.Data)

	if err != nil {
		return tgbotapi.NewMessage(update.Message.Chat.ID, e.Resources.Errors.CommonError)
	}

	switch callback.Action {
	case REGISTRATION_ACTION:
		msg = registrationProcessor(update, e, &callback)

	case SALE_ACTION:
		saleProcessor(update, e)
	case SEARCH_REQUEST_ACTION:
		searchRequestProcessor(update, e)
	default:

	}

	return msg
}

func getCallBack(data string) (CallBack, error) {
	var callback CallBack
	e := json.Unmarshal([]byte(data), &callback)
	if e != nil {
		return CallBack{}, e
	}
	return callback, e
}

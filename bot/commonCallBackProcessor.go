package bot

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"strconv"
)

func CallbackProcessing(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	callback, err := getCallBack(update.CallbackQuery.Data)
	fillCallBackFromTemp(&callback, e.TempData, update.CallbackQuery.From.ID)

	if err != nil {
		return tgbotapi.NewMessage(update.Message.Chat.ID, e.Resources.Errors.CommonError)
	}

	switch callback.Action {
	case REGISTRATION_ACTION:
		msg = registrationProcessor(update, e, &callback)
	case SALE_ACTION:
		//saleProcessor(update, e)
	case SEARCH_REQUEST_ACTION:
		//TODO проверяем зарегистрирован ли пользователь
		msg = searchRequestProcessor(update, e, &callback)
	default:
		msg = CreateErrorMsg(update, e)
	}

	return msg
}

func fillCallBackFromTemp(callBack *CallBack, tempData map[string]TempData, id int) {
	temp := tempData[strconv.Itoa(id)]
	if temp.Action != "" && callBack.Action == "" {
		callBack.Action = temp.Action
	}
}

func getCallBack(data string) (CallBack, error) {
	var callback CallBack
	e := json.Unmarshal([]byte(data), &callback)
	if e != nil {
		return CallBack{}, e
	}
	return callback, e
}

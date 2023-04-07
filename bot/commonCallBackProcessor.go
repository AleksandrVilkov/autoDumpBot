package bot

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"strconv"
)

func CallbackProcessing(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	callback, err := getCallback(e, update)

	if err != nil {
		return CreateErrorMsg(update, e)
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

func getCallback(e *Environment, update *tgbotapi.Update) (CallBack, error) {
	token, err := getToken(update.CallbackQuery.Data)
	if err != nil {
		return CallBack{}, err
	}
	stringCallback := e.TempData[token.Token]
	var callback CallBack
	errJ := json.Unmarshal([]byte(stringCallback), &callback)
	if errJ != nil {
		return CallBack{}, err
	}
	return callback, nil
}

func getToken(data string) (ButtonData, error) {
	var bd ButtonData
	e := json.Unmarshal([]byte(data), &bd)
	if e != nil {
		return ButtonData{}, e
	}
	return bd, e
}

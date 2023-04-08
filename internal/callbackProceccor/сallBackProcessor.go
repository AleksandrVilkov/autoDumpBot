package callbackProceccor

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"psa_dump_bot/bot"
	"psa_dump_bot/model"
)

type CallBackProcessor struct {
}

func NewCallBackProcessor() *CallBackProcessor {
	return &CallBackProcessor{}
}

func (c *CallBackProcessor) StartCallBackProcessor(update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	callback, err := getCallback(e, update)

	if err != nil {
		return e.MessageProcessor.CreateError(update, e)
	}

	switch callback.Action {
	case model.REGISTRATION_ACTION:
		msg = registrationProcessor(update, e, &callback)
	case model.SALE_ACTION:
		//saleProcessor(update, e)
	case model.SEARCH_REQUEST_ACTION:
		//TODO проверяем зарегистрирован ли пользователь
		msg = searchRequestProcessor(update, e, &callback)
	default:
		msg = e.MessageProcessor.CreateError(update, e)
	}
	return msg
}

func getCallback(e *bot.Environment, update *tgbotapi.Update) (model.CallBack, error) {
	token, err := getToken(update.CallbackQuery.Data)
	if err != nil {
		return model.CallBack{}, err
	}
	callback := e.TempData.FindTempDataByToken(token.Token)
	return *callback, nil
}

func getToken(data string) (model.ButtonData, error) {
	var bd model.ButtonData
	e := json.Unmarshal([]byte(data), &bd)
	if e != nil {
		return model.ButtonData{}, e
	}
	return bd, e
}

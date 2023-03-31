package bot

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func CallbackProcessing(update *tgbotapi.Update, temp map[string]TempUserData, e *Environment) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	callback, err := getCallBack(update.CallbackQuery.Data)

	if err != nil {
		return tgbotapi.NewMessage(update.Message.Chat.ID, e.Resources.Errors.CommonError)
	}

	switch callback.Type {
	case START_REGISTER:
		msg = createStartRegisterResponse(update, e)
	case CHOOSE_CONCERN:
		var c Concern
		_ = json.Unmarshal([]byte(callback.Data), &c)
		msg = createChooseConcernResponse(c, update, e)
	//TODO сохранять в temp!
	case CHOOSE_BRAND:
		var b Brand
		_ = json.Unmarshal([]byte(callback.Data), &b)
		msg = createBrandResponse(b, update, e)
	}
	return msg
}

func createStartRegisterResponse(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateChoiceConcernMsg(e))
	concerns := e.Storage.GetConcerns()
	msg.ReplyMarkup = CreateConcernButton(concerns)
	return msg
}

func createChooseConcernResponse(c Concern, update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateChoiceAutoBrandMsg(e))
	brands := e.Storage.GetBrands(c.Concern)
	msg.ReplyMarkup = CreateAutoBrandButton(brands)
	return msg
}

func createBrandResponse(b Brand, update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateModelMsg(e))

	models := e.Storage.GetModels(b.Brand)
	msg.ReplyMarkup = CreateModelsButton(models)
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

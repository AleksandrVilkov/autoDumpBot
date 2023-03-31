package bot

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func CallbackProcessing(update *tgbotapi.Update, temp map[string]TempUserData) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	callback, err := getCallBack(update.CallbackQuery.Data)

	if err != nil {
		return tgbotapi.NewMessage(update.Message.Chat.ID, GetResources().Errors.CommonError)
	}

	switch callback.Type {
	case StartRegister:
		msg = createStartRegisterResponse(update)
	case ChooseConcern:
		var c Concern
		_ = json.Unmarshal([]byte(callback.Data), &c)
		msg = createChooseConcernResponse(c, update)
	//TODO сохранять в temp!
	case ChooseBrand:
		var b Brand
		_ = json.Unmarshal([]byte(callback.Data), &b)
		msg = createBrandResponse(b, update)
	}
	return msg
}

func createStartRegisterResponse(update *tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateChoiceConcernMsg())
	msg.ReplyMarkup = CreateConcernButton()
	return msg
}

func createChooseConcernResponse(c Concern, update *tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateChoiceAutoBrandMsg())
	msg.ReplyMarkup = CreateAutoBrandButton(c)
	return msg
}

func createBrandResponse(b Brand, update *tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateModelMsg())
	msg.ReplyMarkup = CreateModelsButton(b)
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

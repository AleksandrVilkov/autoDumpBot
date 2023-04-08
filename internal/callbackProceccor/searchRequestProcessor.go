package callbackProceccor

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"psa_dump_bot/bot"
	"psa_dump_bot/internal/messageProcessor"
	botModel "psa_dump_bot/model"
)

func searchRequestProcessor(update *tgbotapi.Update, e *bot.Environment, cb *botModel.CallBack) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	switch cb.Subsection {
	case "":
		msg = createStartSearchRequestMsg(update, e, cb)
	case botModel.CHOOSE_CONCERN:
		msg = createConcernForSearchMsg(update, e, cb)
		msg.ReplyToMessageID = update.CallbackQuery.Message.MessageID
	case botModel.CHOOSE_BRAND:
		msg = createBrandForSearchMsg(update, e, cb)
		msg.ReplyToMessageID = update.CallbackQuery.Message.MessageID
	case botModel.ENTER_SEARCH_TEXT:
		msg = createEnterTextForSearchMsg(update, e, cb)
		msg.ReplyToMessageID = update.CallbackQuery.Message.MessageID
		msg.Text += "\nID запроса *" + cb.Token + "*"
	}
	return msg
}

func createStartSearchRequestMsg(update *tgbotapi.Update, e *bot.Environment, cb *botModel.CallBack) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), messageProcessor.CreateConcernMsgForSearch(e))
	concerns := e.Storage.GetConcerns()
	msg.ReplyMarkup = e.ButtonMaker.CreateConcernButton(concerns, cb, e)
	return msg
}

func createConcernForSearchMsg(update *tgbotapi.Update, e *bot.Environment, cb *botModel.CallBack) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), messageProcessor.CreateBrandMsgForSearch(e))
	brands := e.Storage.GetBrands(cb.CarData.Concern)
	msg.ReplyMarkup = e.ButtonMaker.CreateAutoBrandButton(brands, e, cb)
	return msg
}

func createBrandForSearchMsg(update *tgbotapi.Update, e *bot.Environment, cb *botModel.CallBack) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), messageProcessor.CreateModelMsgForSearch(e))
	models := e.Storage.GetModels(cb.CarData.Brand)
	msg.ReplyMarkup = e.ButtonMaker.CreateModelsButton(models, e, cb, botModel.ENTER_SEARCH_TEXT)
	return msg
}

func createEnterTextForSearchMsg(update *tgbotapi.Update, e *bot.Environment, cb *botModel.CallBack) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), messageProcessor.CreateEnterTextMsgForSearch(e))

	return msg
}

package callbackProceccor

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"psa_dump_bot/bot"
	"psa_dump_bot/internal/messageProcessor"
	model2 "psa_dump_bot/model"
)

func searchRequestProcessor(update *tgbotapi.Update, e *bot.Environment, cb *model2.CallBack) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	//switch cb.Subsection {
	//case "":
	//	msg = createStartSearchRequestResponse(update, e)
	//	key := strconv.Itoa(update.CallbackQuery.From.ID)
	//	temp := TempData{
	//		UserId:     strconv.Itoa(update.CallbackQuery.From.ID),
	//		SearchData: TempSearchData{},
	//		Action:     SEARCH_REQUEST_ACTION,
	//	}
	//	e.TempData[key] = temp
	//case CHOOSE_CONCERN:
	//	var c Concern
	//	_ = json.Unmarshal([]byte(cb.Data), &c)
	//	msg = createConcernForSearchResponse(c, update, e)
	//	userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
	//	userTemp.SearchData.Concern = c
	//	userTemp.Action = SEARCH_REQUEST_ACTION
	//	e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp
	//
	//case CHOOSE_BRAND:
	//	var b Brand
	//	_ = json.Unmarshal([]byte(cb.Data), &b)
	//	msg = createBrandForSearchResponse(b, update, e)
	//	userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
	//	userTemp.SearchData.CarBrand = b
	//	userTemp.Action = SEARCH_REQUEST_ACTION
	//	e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp
	//}

	return msg
}

func createStartSearchRequestResponse(update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), messageProcessor.CreateConcernMsgForSearch(e))
	//concerns := e.Storage.GetConcerns()
	////concerns := make([]Concern, 2)
	////concerns[0] = Concern{
	////	Concern: "PSA",
	////}
	//msg.ReplyMarkup = CreateConcernButton(concerns)
	return msg
}

func createConcernForSearchResponse(c model2.Concern, update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), messageProcessor.CreateBrandMsgForSearch(e))
	//brands := e.Storage.GetBrands(c.Concern)
	//brands := make([]Brand, 2)
	//brands[0] = Brand{
	//	Brand: "Ситроян",
	//}
	//brands[1] = Brand{
	//	Brand: "Пеугеот",
	//}
	//msg.ReplyMarkup = CreateAutoBrandButton(brands)
	return msg
}

func createBrandForSearchResponse(b model2.Brand, update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), messageProcessor.CreateModelMsgForSearch(e))
	//models := e.Storage.GetModels(b.Brand)
	//models := make([]Model, 1)
	//models[0] = Model{
	//	Model: "ЦЭ4",
	//}
	//msg.ReplyMarkup = CreateModelsButton(models)
	return msg
}

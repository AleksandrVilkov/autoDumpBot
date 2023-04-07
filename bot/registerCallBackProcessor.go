package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func registrationProcessor(update *tgbotapi.Update, e *Environment, cb *CallBack) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig

	switch cb.Subsection {
	case "":
		msg = createStartRegisterResponse(update, e, cb)
	case CHOOSE_CONCERN:
		msg = createConcernForRegResponse(cb, update, e)
	case CHOOSE_BRAND:
		msg = createBrandForRegResponse(cb, update, e)
	case CHOOSE_MODEL:
	//	msg = createModelForRegResponse(m, userTemp.CarData.CarBrand, update, e)
	case CHOOSE_ENGINE:
	//	msg = createEngineForRegResponse(userTemp.CarData.CarModel, userTemp.CarData.CarBrand, update, e)
	case CHOOSE_BOLT_PATTERN:
	//	msg = createBoltPatternForRegResponse(update, e)
	case CHOOSE_CITY:
	//	var region Region
	//	userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
	//	_ = json.Unmarshal([]byte(cb.Data), &region)
	//	userTemp.Region = region
	//	e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp
	//if e.Storage.SaveUser(CreateUserFromTemp(userTemp)) {
	//	msg = createOkSaveUserResponse(update, e)
	//} else {
	//	msg = createErrSaveUserResponse(update, e)
	//}
	default:
		msg = CreateErrorMsg(update, e)
	}

	return msg
}

func createStartRegisterResponse(update *tgbotapi.Update, e *Environment, cb *CallBack) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateConcernMsgForReg(e))
	concerns := e.Storage.GetConcerns()
	msg.ReplyMarkup = CreateConcernButton(concerns, cb, e)
	return msg
}

func createConcernForRegResponse(cb *CallBack, update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateBrandMsgForReg(e))
	brands := e.Storage.GetBrands(cb.CarData.Concern)
	msg.ReplyMarkup = CreateAutoBrandButton(brands, e, cb)
	return msg
}

func createBrandForRegResponse(cb *CallBack, update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateModelMsgForReg(e))
	models := e.Storage.GetModels(cb.CarData.Brand)
	msg.ReplyMarkup = CreateModelsButton(models, e, cb)
	return msg
}

//
//func createModelForRegResponse(m Model, b Brand, update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
//	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateEngineMsgForReg(e))
//	engines := e.Storage.GetEngines(m.Model, b.Brand)
//	msg.ReplyMarkup = CreateEnginesButton(engines)
//	return msg
//}
//func createEngineForRegResponse(m Model, b Brand, update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
//	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateBoltPatternMsg(e))
//	bps := e.Storage.GetBoltPatterns(m.Model, b.Brand)
//	msg.ReplyMarkup = CreateBoltPatternsButton(bps)
//	return msg
//}
//
//func createBoltPatternForRegResponse(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
//	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateRegionMsg(e))
//	regions := e.Storage.GetAllRegions()
//	msg.ReplyMarkup = CreateRegionsButton(regions)
//	return msg
//}

//
//func createOkSaveUserResponse(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
//	//TODO Если успешно сохранили
//	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), "")
//	return msg
//}
//
//func createErrSaveUserResponse(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
//	//TODO Если не удалось сохранить
//	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), "")
//	return msg
//}

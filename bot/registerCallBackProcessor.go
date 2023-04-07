package bot

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"strconv"
)

func registrationProcessor(update *tgbotapi.Update, e *Environment, cb *CallBack) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig

	switch cb.Subsection {
	case "":
		msg = createStartRegisterResponse(update, e)
		key := strconv.Itoa(update.CallbackQuery.From.ID)
		temp := TempData{
			UserId:  strconv.Itoa(update.CallbackQuery.From.ID),
			CarData: TempCarData{},
			Action:  REGISTRATION_ACTION,
		}
		e.TempData[key] = temp

	case CHOOSE_CONCERN:
		var c Concern
		_ = json.Unmarshal([]byte(cb.Data), &c)
		msg = createConcernForRegResponse(c, update, e)
		userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
		userTemp.CarData.Concern = c
		userTemp.Action = REGISTRATION_ACTION
		e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp

	case CHOOSE_BRAND:
		var b Brand
		_ = json.Unmarshal([]byte(cb.Data), &b)
		msg = createBrandForRegResponse(b, update, e)
		userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
		userTemp.CarData.CarBrand = b
		userTemp.Action = REGISTRATION_ACTION
		e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp
	case CHOOSE_MODEL:
		var m Model
		_ = json.Unmarshal([]byte(cb.Data), &m)
		userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
		userTemp.CarData.CarModel = m
		userTemp.Action = REGISTRATION_ACTION
		e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp
		msg = createModelForRegResponse(m, userTemp.CarData.CarBrand, update, e)

	case CHOOSE_ENGINE:
		var en Engine
		_ = json.Unmarshal([]byte(cb.Data), &en)
		userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
		userTemp.CarData.CarEngine = en
		userTemp.Action = REGISTRATION_ACTION
		e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp
		msg = createEngineForRegResponse(userTemp.CarData.CarModel, userTemp.CarData.CarBrand, update, e)

	case CHOOSE_BOLT_PATTERN:
		var bp BoltPattern
		_ = json.Unmarshal([]byte(cb.Data), &bp)
		userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
		userTemp.CarData.BoltPattern = bp
		userTemp.Action = REGISTRATION_ACTION
		e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp
		msg = createBoltPatternForRegResponse(update, e)
	case CHOOSE_CITY:
		var region Region
		userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
		_ = json.Unmarshal([]byte(cb.Data), &region)
		userTemp.Region = region
		e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp

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

func createStartRegisterResponse(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateConcernMsgForReg(e))
	concerns := e.Storage.GetConcerns()
	msg.ReplyMarkup = CreateConcernButton(concerns)
	return msg
}

func createConcernForRegResponse(c Concern, update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateBrandMsgForReg(e))
	brands := e.Storage.GetBrands(c.Concern)
	msg.ReplyMarkup = CreateAutoBrandButton(brands)
	return msg
}

func createBrandForRegResponse(b Brand, update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateModelMsgForReg(e))

	models := e.Storage.GetModels(b.Brand)
	msg.ReplyMarkup = CreateModelsButton(models)
	return msg
}

func createModelForRegResponse(m Model, b Brand, update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateEngineMsgForReg(e))
	engines := e.Storage.GetEngines(m.Model, b.Brand)
	msg.ReplyMarkup = CreateEnginesButton(engines)
	return msg
}
func createEngineForRegResponse(m Model, b Brand, update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateBoltPatternMsg(e))
	bps := e.Storage.GetBoltPatterns(m.Model, b.Brand)
	msg.ReplyMarkup = CreateBoltPatternsButton(bps)
	return msg
}

func createBoltPatternForRegResponse(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateRegionMsg(e))
	regions := e.Storage.GetAllRegions()
	msg.ReplyMarkup = CreateRegionsButton(regions)
	return msg
}

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

package bot

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"strconv"
)

func CallbackProcessing(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	callback, err := getCallBack(update.CallbackQuery.Data)

	if err != nil {
		return tgbotapi.NewMessage(update.Message.Chat.ID, e.Resources.Errors.CommonError)
	}

	switch callback.Type {
	case START_REGISTER:
		msg = createStartRegisterResponse(update, e)
		key := strconv.Itoa(update.CallbackQuery.From.ID)
		temp := TempData{
			UserId:  strconv.Itoa(update.CallbackQuery.From.ID),
			CarData: TempCarData{},
		}
		e.TempData[key] = temp
	case CHOOSE_CONCERN:
		var c Concern
		_ = json.Unmarshal([]byte(callback.Data), &c)
		msg = createChooseConcernResponse(c, update, e)
		userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
		userTemp.CarData.Concern = c
		e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp
	case CHOOSE_BRAND:
		var b Brand
		_ = json.Unmarshal([]byte(callback.Data), &b)
		msg = createBrandResponse(b, update, e)
		userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
		userTemp.CarData.CarBrand = b
		e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp
	case CHOOSE_MODEL:
		var m Model
		_ = json.Unmarshal([]byte(callback.Data), &m)
		userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
		userTemp.CarData.CarModel = m
		e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp
		msg = createModelResponse(m, userTemp.CarData.CarBrand, update, e)

	case CHOOSE_ENGINE:
		var en Engine
		_ = json.Unmarshal([]byte(callback.Data), &en)
		userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
		userTemp.CarData.CarEngine = en
		e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp
		msg = createEngineResponse(userTemp.CarData.CarModel, userTemp.CarData.CarBrand, update, e)

	case CHOOSE_BOLT_PATTERN:
		var bp BoltPattern
		_ = json.Unmarshal([]byte(callback.Data), &bp)
		userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
		userTemp.CarData.BoltPattern = bp
		e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp
		msg = createBoltPatternResponse(update, e)

	case CHOOSE_CITY:
		var region Region
		userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
		_ = json.Unmarshal([]byte(callback.Data), &region)
		userTemp.Region = region
		e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)] = userTemp

		if e.Storage.SaveUser(CreateUserFromTemp(userTemp)) {
			msg = createOkSaveUserResponse(update, e)
		} else {
			msg = createErrSaveUserResponse(update, e)
		}
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

func createModelResponse(m Model, b Brand, update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateEngineMsg(e))
	engines := e.Storage.GetEngines(m.Model, b.Brand)
	msg.ReplyMarkup = CreateEnginesButton(engines)
	return msg
}
func createEngineResponse(m Model, b Brand, update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateBoltPatternMsg(e))
	bps := e.Storage.GetBoltPatterns(m.Model, b.Brand)
	msg.ReplyMarkup = CreateBoltPatternsButton(bps)
	return msg
}

func createBoltPatternResponse(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateRegionMsg(e))
	regions := e.Storage.GetAllRegions()
	msg.ReplyMarkup = CreateRegionsButton(regions)
	return msg
}

func createOkSaveUserResponse(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	//TODO Если успешно сохранили
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), "")
	return msg
}

func createErrSaveUserResponse(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
	//TODO Если не удалось сохранить
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), "")
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

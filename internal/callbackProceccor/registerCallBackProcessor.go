package callbackProceccor

import (
	"errors"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"psa_dump_bot/bot"
	"psa_dump_bot/internal/messageProcessor"
	botModel "psa_dump_bot/model"
	"time"
)

func registrationProcessor(update *tgbotapi.Update, e *bot.Environment, cb *botModel.CallBack) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig

	switch cb.Subsection {
	case "":
		msg = createStartRegisterMsg(update, e, cb)
	case botModel.CHOOSE_CONCERN:
		msg = createConcernMsgForReg(cb, update, e)
	case botModel.CHOOSE_BRAND:
		msg = createBrandMsgForReg(cb, update, e)
	case botModel.CHOOSE_MODEL:
		msg = createModelMsgForReg(cb, update, e)
	case botModel.CHOOSE_ENGINE:
		msg = createEngineMsgForReg(cb, update, e)
	case botModel.CHOOSE_BOLT_PATTERN:
		msg = createBoltPatternMsgForReg(cb, update, e)
	case botModel.CHOOSE_CITY:
		msg = createFinishMsgForReg(cb, update, e)
	default:
		msg = e.MessageProcessor.CreateError(update, e)
	}

	return msg
}

func createStartRegisterMsg(update *tgbotapi.Update, e *bot.Environment, cb *botModel.CallBack) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), messageProcessor.CreateConcernMsgForReg(e))
	concerns := e.Storage.GetConcerns()
	msg.ReplyMarkup = e.ButtonMaker.CreateConcernButton(concerns, cb, e)
	return msg
}

func createConcernMsgForReg(cb *botModel.CallBack, update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), messageProcessor.CreateBrandMsgForReg(e))
	brands := e.Storage.GetBrands(cb.CarData.Concern)
	msg.ReplyMarkup = e.ButtonMaker.CreateAutoBrandButton(brands, e, cb)
	return msg
}

func createBrandMsgForReg(cb *botModel.CallBack, update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), messageProcessor.CreateModelMsgForReg(e))
	models := e.Storage.GetModels(cb.CarData.Brand)
	msg.ReplyMarkup = e.ButtonMaker.CreateModelsButton(models, e, cb)
	return msg
}

func createModelMsgForReg(cb *botModel.CallBack, update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), messageProcessor.CreateEngineMsgForReg(e))
	engines := e.Storage.GetEngines(cb.CarData.Model, cb.CarData.Brand)
	msg.ReplyMarkup = e.ButtonMaker.CreateEnginesButton(engines, e, cb)
	return msg
}

func createEngineMsgForReg(cb *botModel.CallBack, update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), messageProcessor.CreateBoltPatternMsg(e))
	bps := e.Storage.GetBoltPatterns(cb.CarData.Model, cb.CarData.Brand)
	msg.ReplyMarkup = e.ButtonMaker.CreateBoltPatternsButton(bps, e, cb)
	return msg
}

func createBoltPatternMsgForReg(cb *botModel.CallBack, update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), messageProcessor.CreateRegionMsg(e))
	regions := e.Storage.GetAllRegions()
	msg.ReplyMarkup = e.ButtonMaker.CreateRegionsButton(regions, e, cb)
	return msg
}

func createFinishMsgForReg(cb *botModel.CallBack, update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	user, err := createUserFromCallback(cb)

	if err != nil {
		return tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), e.Resources.Errors.ErrorReservation)
	}

	if !e.Storage.SaveUser(user) {
		return tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), e.Resources.Errors.ErrorReservation)
	}
	return tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), e.Resources.Success.SuccessReservation)
}

func createUserFromCallback(cb *botModel.CallBack) (*botModel.User, error) {
	if !validateCallbackForCreateUser(cb) {
		log.Print("incomplete data for the user record, userID = " + cb.UserId)
		return &botModel.User{}, errors.New("incomplete data for the user record")
	}
	return &botModel.User{
		CreateDate: time.Time{},
		Role:       botModel.USER_ROLE,
		Login:      cb.UserId,
		Region: botModel.Region{
			Id:         cb.UserData.RegionId,
			RegionName: cb.UserData.RegionName,
		},
		UserCar: botModel.UserCar{
			CreateDate:  time.Time{},
			Concern:     botModel.Concern{Concern: cb.CarData.Concern},
			Model:       botModel.Model{Model: cb.CarData.Model},
			Engine:      botModel.Engine{EngineName: cb.CarData.EngineName},
			BoltPattern: botModel.BoltPattern{BoltPatternSize: cb.CarData.BoltPatternSize},
			Brand:       botModel.Brand{Brand: cb.CarData.Brand},
		},
	}, nil
}

func validateCallbackForCreateUser(cb *botModel.CallBack) bool {
	return !(cb.UserId == "" ||
		cb.UserData.RegionId == 0 ||
		cb.CarData.Concern == "" ||
		cb.CarData.Model == "" ||
		cb.CarData.EngineName == "" ||
		cb.CarData.BoltPatternSize == "" ||
		cb.CarData.Brand == "")
}

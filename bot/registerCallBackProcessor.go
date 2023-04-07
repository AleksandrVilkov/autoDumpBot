package bot

import (
	"errors"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"psa_dump_bot/bot/model"
	"time"
)

func registrationProcessor(update *tgbotapi.Update, e *model.Environment, cb *model.CallBack) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig

	switch cb.Subsection {
	case "":
		msg = createStartRegisterMsg(update, e, cb)
	case model.CHOOSE_CONCERN:
		msg = createConcernMsgForReg(cb, update, e)
	case model.CHOOSE_BRAND:
		msg = createBrandMsgForReg(cb, update, e)
	case model.CHOOSE_MODEL:
		msg = createModelMsgForReg(cb, update, e)
	case model.CHOOSE_ENGINE:
		msg = createEngineMsgForReg(cb, update, e)
	case model.CHOOSE_BOLT_PATTERN:
		msg = createBoltPatternMsgForReg(cb, update, e)
	case model.CHOOSE_CITY:
		msg = createFinishMsgForReg(cb, update, e)
	default:
		msg = CreateErrorMsg(update, e)
	}

	return msg
}

func createStartRegisterMsg(update *tgbotapi.Update, e *model.Environment, cb *model.CallBack) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateConcernMsgForReg(e))
	concerns := e.Storage.GetConcerns()
	msg.ReplyMarkup = CreateConcernButton(concerns, cb, e)
	return msg
}

func createConcernMsgForReg(cb *model.CallBack, update *tgbotapi.Update, e *model.Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateBrandMsgForReg(e))
	brands := e.Storage.GetBrands(cb.CarData.Concern)
	msg.ReplyMarkup = CreateAutoBrandButton(brands, e, cb)
	return msg
}

func createBrandMsgForReg(cb *model.CallBack, update *tgbotapi.Update, e *model.Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateModelMsgForReg(e))
	models := e.Storage.GetModels(cb.CarData.Brand)
	msg.ReplyMarkup = CreateModelsButton(models, e, cb)
	return msg
}

func createModelMsgForReg(cb *model.CallBack, update *tgbotapi.Update, e *model.Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateEngineMsgForReg(e))
	engines := e.Storage.GetEngines(cb.CarData.Model, cb.CarData.Brand)
	msg.ReplyMarkup = CreateEnginesButton(engines, e, cb)
	return msg
}

func createEngineMsgForReg(cb *model.CallBack, update *tgbotapi.Update, e *model.Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateBoltPatternMsg(e))
	bps := e.Storage.GetBoltPatterns(cb.CarData.Model, cb.CarData.Brand)
	msg.ReplyMarkup = CreateBoltPatternsButton(bps, e, cb)
	return msg
}

func createBoltPatternMsgForReg(cb *model.CallBack, update *tgbotapi.Update, e *model.Environment) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), CreateRegionMsg(e))
	regions := e.Storage.GetAllRegions()
	msg.ReplyMarkup = CreateRegionsButton(regions, e, cb)
	return msg
}

func createFinishMsgForReg(cb *model.CallBack, update *tgbotapi.Update, e *model.Environment) tgbotapi.MessageConfig {
	user, err := createUserFromCallback(cb)

	if err != nil {
		return tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), e.Resources.Errors.ErrorReservation)
	}

	if !e.Storage.SaveUser(user) {
		return tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), e.Resources.Errors.ErrorReservation)
	}
	return tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), e.Resources.Success.SuccessReservation)
}

func createUserFromCallback(cb *model.CallBack) (*model.User, error) {
	if !validateCallbackForCreateUser(cb) {
		log.Print("incomplete data for the user record, userID = " + cb.UserId)
		return &model.User{}, errors.New("incomplete data for the user record")
	}
	return &model.User{
		CreateDate: time.Time{},
		Role:       model.USER_ROLE,
		Login:      cb.UserId,
		Region: model.Region{
			Id:         cb.UserData.RegionId,
			RegionName: cb.UserData.RegionName,
		},
		UserCar: model.UserCar{
			CreateDate:  time.Time{},
			Concern:     model.Concern{Concern: cb.CarData.Concern},
			Model:       model.Model{Model: cb.CarData.Model},
			Engine:      model.Engine{EngineName: cb.CarData.EngineName},
			BoltPattern: model.BoltPattern{BoltPatternSize: cb.CarData.BoltPatternSize},
			Brand:       model.Brand{Brand: cb.CarData.Brand},
		},
	}, nil
}

func validateCallbackForCreateUser(cb *model.CallBack) bool {
	return !(cb.UserId == "" ||
		cb.UserData.RegionId == 0 ||
		cb.CarData.Concern == "" ||
		cb.CarData.Model == "" ||
		cb.CarData.EngineName == "" ||
		cb.CarData.BoltPatternSize == "" ||
		cb.CarData.Brand == "")
}

package bot

import (
	"psa_dump_bot/bot/model"
)

func CreateWelcomeMsg(e *model.Environment) string {
	return e.Resources.Msgs.WelcomeMessage
}
func CreateWelcomeRegisteredMsg(e *model.Environment) string {
	return e.Resources.Msgs.WelcomeRegistered
}
func CreateErrAuthMsg(e *model.Environment) string {
	return e.Resources.Errors.AuthError + " \nПодпишись: \n" + e.Config.ValidateData.ChannelUrl
}

func CreateConcernMsgForReg(e *model.Environment) string {
	return e.Resources.Msgs.Registration.ChoiceConcern
}

func CreateBrandMsgForReg(e *model.Environment) string {
	return e.Resources.Msgs.Registration.CarBrandEnter
}

func CreateModelMsgForReg(e *model.Environment) string {
	return e.Resources.Msgs.Registration.CarModelEnter
}
func CreateEngineMsgForReg(e *model.Environment) string {
	return e.Resources.Msgs.Registration.CarEngineEnter
}

func CreateBoltPatternMsg(e *model.Environment) string {
	return e.Resources.Msgs.Registration.CarBoltPatternEnter
}

func CreateRegionMsg(e *model.Environment) string {
	return e.Resources.Msgs.Registration.RegionEnter
}

func CreateConcernMsgForSearch(e *model.Environment) string {
	return e.Resources.Msgs.Search.ChoiceConcern
}

func CreateBrandMsgForSearch(e *model.Environment) string {
	return e.Resources.Msgs.Search.CarBrandEnter
}
func CreateModelMsgForSearch(e *model.Environment) string {
	return e.Resources.Msgs.Search.CarModelEnter
}

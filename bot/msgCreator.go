package bot

func CreateWelcomeMsg(e *Environment) string {
	return e.Resources.Msgs.WelcomeMessage
}
func CreateErrAuthMsg(e *Environment) string {
	return e.Resources.Errors.AuthError + " \nПодпишись: \n" + e.Config.ValidateData.ChannelUrl
}

func CreateConcernMsgForReg(e *Environment) string {
	return e.Resources.Msgs.Registration.ChoiceConcern
}

func CreateBrandMsgForReg(e *Environment) string {
	return e.Resources.Msgs.Registration.CarBrandEnter
}

func CreateModelMsgForReg(e *Environment) string {
	return e.Resources.Msgs.Registration.CarModelEnter
}
func CreateEngineMsgForReg(e *Environment) string {
	return e.Resources.Msgs.Registration.CarEngineEnter
}

func CreateBoltPatternMsg(e *Environment) string {
	return e.Resources.Msgs.Registration.CarBoltPatternEnter
}

func CreateRegionMsg(e *Environment) string {
	return e.Resources.Msgs.Registration.RegionEnter
}

func CreateConcernMsgForSearch(e *Environment) string {
	return e.Resources.Msgs.Search.ChoiceConcern
}

func CreateBrandMsgForSearch(e *Environment) string {
	return e.Resources.Msgs.Search.CarBrandEnter
}
func CreateModelMsgForSearch(e *Environment) string {
	return e.Resources.Msgs.Search.CarModelEnter
}

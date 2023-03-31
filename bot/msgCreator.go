package bot

func CreateWelcomeMsg(e *Environment) string {
	return e.Resources.Msgs.WelcomeMessage
}
func CreateErrAuthMsg(e *Environment) string {
	return e.Resources.Errors.AuthError + " \nПодпишись: \n" + e.Config.ValidateData.ChannelUrl
}

func CreateChoiceConcernMsg(e *Environment) string {
	return e.Resources.Msgs.ChoiceConcern
}
func CreateChoiceAutoBrandMsg(e *Environment) string {
	return e.Resources.Msgs.SuccessCarBrandEnter
}

func CreateModelMsg(e *Environment) string {
	return e.Resources.Msgs.SuccessCarModelEnter
}
func CreateEngineMsg(e *Environment) string {
	return e.Resources.Msgs.SuccessCarEngineEnter
}

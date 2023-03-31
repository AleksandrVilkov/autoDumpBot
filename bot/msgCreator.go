package bot

func CreateWelcomeMsg() string {
	return GetResources().Msgs.WelcomeMessage
}
func CreateErrAuthMsg(chUrl string) string {
	return GetResources().Errors.AuthError + " \nПодпишись: \n" + chUrl
}

func CreateChoiceConcernMsg() string {
	return GetResources().Msgs.ChoiceConcern
}
func CreateChoiceAutoBrandMsg() string {
	return GetResources().Msgs.SuccessCarBrandEnter
}

func CreateModelMsg() string {
	return GetResources().Msgs.SuccessCarModelEnter
}
func CreateEngineMsg() string {
	return GetResources().Msgs.SuccessCarEngineEnter
}

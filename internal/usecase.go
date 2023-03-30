package internal

func CreateWelcomeMsg() string {
	return getResources().Msgs.WelcomeMessage
}

func CreateRulesMsg() string {
	return getResources().Msgs.Rules
}

func CreateErrAuthMsg(chUrl string) string {
	return getResources().Errors.AuthError + " \nПодпишись: \n" + chUrl
}

func CreatePreRegistrationMsg() string {
	return getResources().Msgs.PreRegistration
}

func CreateSuccessCarBrandEnter() string {
	return getResources().Msgs.SuccessCarBrandEnter
}
func CreateSuccessCarModelEnter() string {
	return getResources().Msgs.SuccessCarModelEnter
}
func CreateSuccessCarEngineEnter() string {
	return getResources().Msgs.SuccessCarEngineEnter
}
func CreateError() string {
	return getResources().Errors.CommonError
}

func CreatesReservation() string {
	return getResources().Msgs.SuccessReservation
}

func CreateSubscriptionSaleMsg() string {
	return "CreateSubscriptionSaleMsg"
}
func CreateStartSaleMsg() string {
	return "CreateStartSaleMsg"
}

package internal

type resources struct {
	Msgs struct {
		WelcomeMessage        string `yaml:"welcomeMessage"`
		Rules                 string `yaml:"rules"`
		PreRegistration       string `yaml:"preRegistration"`
		SuccessCarBrandEnter  string `yaml:"successCarBrandEnter"`
		SuccessCarModelEnter  string `yaml:"successCarModelEnter"`
		SuccessCarEngineEnter string `yaml:"successCarEngineEnter"`
		SuccessReservation    string `yaml:"successReservation"`
	}
	Errors struct {
		ErrorReservation string `yaml:"errorReservation"`
		AuthError        string `yaml:"authError"`
		CommonError      string `yaml:"commonError"`
	}
}

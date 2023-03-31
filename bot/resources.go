package bot

type Resources struct {
	Buttonstext struct {
		Registration  string `yaml:"registration"`
		Subscribe     string `yaml:"subscribe"`
		SearchRequest string `yaml:"searchRequest"`
		PlaceAnAd     string `yaml:"placeAnAd"`
		Rules         string `yaml:"rules"`
	}

	Msgs struct {
		WelcomeMessage        string `yaml:"welcomeMessage"`
		Rules                 string `yaml:"rules"`
		ChoiceConcern         string `yaml:"choiceConcern"`
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

package bot

type Resources struct {
	Buttonstext struct {
		Registration  string `yaml:"registration"`
		Subscribe     string `yaml:"subscribe"`
		SearchRequest string `yaml:"searchRequest"`
		PlaceAnAd     string `yaml:"placeAnAd"`
		Rules         string `yaml:"rules"`
		Universal     string `yaml:"universal"`
	}

	Msgs struct {
		Registration struct {
			ChoiceConcern       string `yaml:"choiceConcern"`
			CarBrandEnter       string `yaml:"сarBrandEnter"`
			CarModelEnter       string `yaml:"сarModelEnter"`
			CarEngineEnter      string `yaml:"сarEngineEnter"`
			CarBoltPatternEnter string `yaml:"carBoltPatternEnter"`
			RegionEnter         string `yaml:"regionEnter"`
			SuccessReservation  string `yaml:"successReservation"`
		}
		Search struct {
			ChoiceConcern string `yaml:"choiceConcern"`
			CarBrandEnter string `yaml:"сarBrandEnter"`
			CarModelEnter string `yaml:"сarModelEnter"`
			SearchText    string `yaml:"searchText"`
		}

		WelcomeMessage    string `yaml:"welcomeMessage"`
		WelcomeRegistered string `yaml:"welcomeRegistered"`
		Rules             string `yaml:"rules"`

		StartSale string `yaml:"startSale"`
	}
	Errors struct {
		ErrorReservation string `yaml:"errorReservation"`
		AuthError        string `yaml:"authError"`
		CommonError      string `yaml:"commonError"`
	}
	Success struct {
		SuccessReservation string `yaml:"successReservation"`
	}
}

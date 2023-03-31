package bot

type ChatMemberStatus string

const (
	MEMBER  ChatMemberStatus = "member"        //пользователь является подписчиком;
	LEFT                     = "left"          //пользователь не подписан;
	KICKED                   = "kicked"        // пользователь заблокирован;
	ADMIN                    = "administrator" //админ
	CREATOR                  = "creator"       //создатель
)

type CallbackType string

const (
	START_REGISTER CallbackType = "startRegistration"
	CHOOSE_CONCERN              = "chooseConcern"
	CHOOSE_BRAND                = "chooseBrand"
	CHOOSE_MODEL                = "chooseModel"
	SUBSCRIBE                   = "subscribe"
	SEARCH_REQUEST              = "searchRequest"
	PLACE_AN_AD                 = "placeAnAd"
	RULES                       = "rules"
)

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

type Config struct {
	Token    string `yaml:"token"`
	Commands struct {
		Start        string `yaml:"start"`
		Rules        string `yaml:"rules"`
		Registration string `yaml:"registration"`
		Subscription string `yaml:"subscription"`
		Sale         string `yaml:"sale"`
	}
	InternalCommands struct {
		EnterCarBrand  string `yaml:"enterCarBrand"`
		EnterCarModel  string `yaml:"enterCarModel"`
		EnterCarEngine string `yaml:"enterCarEngine"`
	}

	ValidateData struct {
		ChannelID  int64  `yaml:"channelID"`
		ChannelUrl string `yaml:"channelUrl"`
	}
}

func (c *Config) printCommands() string {
	return "\n" + c.Commands.Start + "\n" +
		"\n" + c.Commands.Rules +
		"\n" + c.Commands.Registration +
		"\n" + c.Commands.Subscription +
		"\n" + c.Commands.Sale

}

type TempUserData struct {
	User struct {
		Id string
	}
	Action struct {
		MainAction  string
		LastCommand string
	}
	CarData struct {
		Concern     Concern
		CarBrand    Brand
		CarModel    Model
		CarEngine   Engine
		BoltPattern string
	}
	SaleData struct {
	}
	SubscriptionData struct {
	}
}

type CallBack struct {
	Type CallbackType `json:"type"`
	Data string       `json:"data"`
}

type Concern struct {
	Concern string `json:"concern"`
}
type Brand struct {
	Brand string `json:"brand"`
}
type Model struct {
	Model string `json:"model"`
}

type Engine struct {
	EngineName string `json:"engineName"`
}
type BoltPattern struct {
}

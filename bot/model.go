package bot

import "time"

type ChatMemberStatus string

const (
	MEMBER  ChatMemberStatus = "member"        //пользователь является подписчиком;
	LEFT                     = "left"          //пользователь не подписан;
	KICKED                   = "kicked"        // пользователь заблокирован;
	ADMIN                    = "administrator" //админ
	CREATOR                  = "creator"       //создатель
)

type Role string

const (
	ADMIN_ROLE  Role = "ADMIN"  //пользователь является админ;
	KICKED_ROLE      = "KICKED" // пользователь заблокирован;
	USER_ROLE        = "USER"   //пользователь
)

type CallbackType string

const (
	START_REGISTER      CallbackType = "startRegistration"
	CHOOSE_CONCERN                   = "chooseConcern"
	CHOOSE_BRAND                     = "chooseBrand"
	CHOOSE_MODEL                     = "chooseModel"
	CHOOSE_ENGINE                    = "chooseEngine"
	CHOOSE_BOLT_PATTERN              = "chooseBP"
	CHOOSE_CITY                      = "chooseCity"
	SUBSCRIBE                        = "subscribe"
	SEARCH_REQUEST                   = "searchRequest"
	PLACE_AN_AD                      = "placeAnAd"
	RULES                            = "rules"
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
		WelcomeMessage      string `yaml:"welcomeMessage"`
		Rules               string `yaml:"rules"`
		ChoiceConcern       string `yaml:"choiceConcern"`
		CarBrandEnter       string `yaml:"сarBrandEnter"`
		CarModelEnter       string `yaml:"сarModelEnter"`
		CarEngineEnter      string `yaml:"сarEngineEnter"`
		CarBoltPatternEnter string `yaml:"carBoltPatternEnter"`
		RegionEnter         string `yaml:"regionEnter"`
		SuccessReservation  string `yaml:"successReservation"`
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
	Psql struct {
		Login              string `yaml:"login"`
		Password           string `yaml:"password"`
		SslMode            string `yaml:"ssl_mode"`
		DriverName         string `yaml:"driver_name"`
		DatabaseName       string `yaml:"database_name"`
		AttemptsConnection string `yaml:"attempts_connection"`
	}
}
type Environment struct {
	Config    *Config
	Storage   Storage
	Resources *Resources
	TempData  map[string]TempData
}

func (c *Config) printCommands() string {
	return "\n" + c.Commands.Start + "\n" +
		"\n" + c.Commands.Rules +
		"\n" + c.Commands.Registration +
		"\n" + c.Commands.Subscription +
		"\n" + c.Commands.Sale

}

type TempData struct {
	UserId           string
	CarData          TempCarData
	SaleData         TempSaleData
	SubscriptionData TempSubscriptionData
	Region           Region
}

type TempCarData struct {
	Concern     Concern
	CarBrand    Brand
	CarModel    Model
	CarEngine   Engine
	BoltPattern BoltPattern
}

type TempSaleData struct {
}

type TempSubscriptionData struct{}

type User struct {
	Id         int
	CreateDate time.Time
	Role       Role
	Login      string
	Region     Region
	UserCar    UserCar
}

type UserCar struct {
	Id          int
	CreateDate  time.Time
	Concern     Concern
	Model       Model
	Engine      Engine
	BoltPattern BoltPattern
	Brand       Brand
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
	BoltPatternSize string `json:"boltPattern"`
}

type Region struct {
	RegionName string `json:"name"`
}

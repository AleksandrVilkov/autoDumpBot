package model

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
	START_REGISTER CallbackType = "startRegistration"

	SUBSCRIBE      = "subscribe"
	SEARCH_REQUEST = "searchRequest"
	PLACE_AN_AD    = "placeAnAd"
	RULES          = "rules"
	NONE           = "none"
)

type Action string

const (
	REGISTRATION_ACTION   Action = "register"
	SALE_ACTION                  = "sale"
	SEARCH_REQUEST_ACTION        = "searchRequest"
	SUBSCRIBE_ACTION             = "subscribe"
	RULES_ACTION                 = "rules"
)

type CallbackSubsection string

const (
	CHOOSE_CONCERN      CallbackSubsection = "concern"
	CHOOSE_BRAND                           = "brand"
	CHOOSE_MODEL                           = "model"
	CHOOSE_ENGINE                          = "engine"
	CHOOSE_BOLT_PATTERN                    = "bp"
	CHOOSE_CITY                            = "city"
)

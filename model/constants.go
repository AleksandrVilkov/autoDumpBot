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

package model

type Action string

const (
	REGISTRATION_ACTION   Action = "register"
	SALE_ACTION                  = "sale"
	SEARCH_REQUEST_ACTION        = "searchRequest"
	SUBSCRIBE_ACTION             = "subscribe"
	RULES_ACTION                 = "rules"
)

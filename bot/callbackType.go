package bot

type CallbackType string

const (
	StartRegister CallbackType = "startRegistration"
	ChooseConcern              = "chooseConcern"
	ChooseBrand                = "chooseBrand"
	ChooseModel                = "chooseModel"
	Subscribe                  = "subscribe"
	SearchRequest              = "searchRequest"
	PlaceAnAd                  = "placeAnAd"
	Rules                      = "rules"
)

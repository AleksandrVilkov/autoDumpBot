package bot

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func CreateMainButtons() tgbotapi.InlineKeyboardMarkup {
	res := GetResources()
	data := make(map[string]string)

	registration, _ := json.Marshal(CallBack{
		Type: START_REGISTER,
		Data: "",
	})
	data[res.Buttonstext.Registration] = string(registration)

	subscribe, _ := json.Marshal(CallBack{
		Type: SUBSCRIBE,
		Data: "",
	})
	data[res.Buttonstext.Subscribe] = string(subscribe)

	searchRequest, _ := json.Marshal(CallBack{
		Type: SEARCH_REQUEST,
		Data: "",
	})
	data[res.Buttonstext.SearchRequest] = string(searchRequest)

	placeAnAd, _ := json.Marshal(CallBack{
		Type: PLACE_AN_AD,
		Data: "",
	})
	data[res.Buttonstext.PlaceAnAd] = string(placeAnAd)

	rules, _ := json.Marshal(CallBack{
		Type: RULES,
		Data: "",
	})
	data[res.Buttonstext.Rules] = string(rules)
	return CreateInlineKeyBoard(data, 1)
}

func CreateConcernButton(concerns []Concern) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(concerns); i++ {

		concern, _ := json.Marshal(Concern{Concern: concerns[i].Concern})
		concernData, _ := json.Marshal(CallBack{
			Type: CHOOSE_CONCERN,
			Data: string(concern),
		})

		data[concerns[i].Concern] = string(concernData)
	}

	return CreateInlineKeyBoard(data, 1)
}

func CreateAutoBrandButton(brands []Brand) tgbotapi.InlineKeyboardMarkup {

	data := make(map[string]string)
	for i := 0; i < len(brands); i++ {

		b, _ := json.Marshal(Brand{Brand: brands[i].Brand})
		concernData, _ := json.Marshal(CallBack{
			Type: CHOOSE_BRAND,
			Data: string(b),
		})

		data[brands[i].Brand] = string(concernData)
	}

	return CreateInlineKeyBoard(data, 1)
}

func CreateModelsButton(models []Model) tgbotapi.InlineKeyboardMarkup {

	data := make(map[string]string)
	for i := 0; i < len(models); i++ {

		b, _ := json.Marshal(Model{Model: models[i].Model})
		concernData, _ := json.Marshal(CallBack{
			Type: CHOOSE_MODEL,
			Data: string(b),
		})

		data[models[i].Model] = string(concernData)
	}

	return CreateInlineKeyBoard(data, 1)
}

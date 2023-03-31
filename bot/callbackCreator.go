package bot

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func CreateMainButtons() tgbotapi.InlineKeyboardMarkup {
	res := GetResources()
	data := make(map[string]string)

	registration, _ := json.Marshal(CallBack{
		Type: StartRegister,
		Data: "",
	})
	data[res.Buttonstext.Registration] = string(registration)

	subscribe, _ := json.Marshal(CallBack{
		Type: Subscribe,
		Data: "",
	})
	data[res.Buttonstext.Subscribe] = string(subscribe)

	searchRequest, _ := json.Marshal(CallBack{
		Type: SearchRequest,
		Data: "",
	})
	data[res.Buttonstext.SearchRequest] = string(searchRequest)

	placeAnAd, _ := json.Marshal(CallBack{
		Type: PlaceAnAd,
		Data: "",
	})
	data[res.Buttonstext.PlaceAnAd] = string(placeAnAd)

	rules, _ := json.Marshal(CallBack{
		Type: Rules,
		Data: "",
	})
	data[res.Buttonstext.Rules] = string(rules)
	return CreateInlineKeyBoard(data, 1)
}

func CreateConcernButton() tgbotapi.InlineKeyboardMarkup {

	concerns := make([]string, 1)
	concerns[0] = "PSA"

	data := make(map[string]string)
	for i := 0; i < len(concerns); i++ {

		concern, _ := json.Marshal(Concern{Concern: concerns[i]})
		concernData, _ := json.Marshal(CallBack{
			Type: ChooseConcern,
			Data: string(concern),
		})

		data[concerns[i]] = string(concernData)
	}

	return CreateInlineKeyBoard(data, 1)
}

func CreateAutoBrandButton(c Concern) tgbotapi.InlineKeyboardMarkup {
	brands := make([]string, 3)
	brands[0] = "Peugeot"
	brands[1] = "Citroen"
	brands[2] = "DS"

	data := make(map[string]string)
	for i := 0; i < len(brands); i++ {

		b, _ := json.Marshal(Brand{Brand: brands[i]})
		concernData, _ := json.Marshal(CallBack{
			Type: ChooseBrand,
			Data: string(b),
		})

		data[brands[i]] = string(concernData)
	}

	return CreateInlineKeyBoard(data, 1)
}

func CreateModelsButton(c Brand) tgbotapi.InlineKeyboardMarkup {
	models := make([]string, 3)
	models[0] = "206"
	models[1] = "C4"
	models[2] = "307"

	data := make(map[string]string)
	for i := 0; i < len(models); i++ {

		b, _ := json.Marshal(Model{Model: models[i]})
		concernData, _ := json.Marshal(CallBack{
			Type: ChooseModel,
			Data: string(b),
		})

		data[models[i]] = string(concernData)
	}

	return CreateInlineKeyBoard(data, 1)
}

//TODO Подключить БД, где будут храниться варианты, и дергать оттуда

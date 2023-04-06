package bot

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func CreateMainButtons(e *Environment) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)

	registration, _ := json.Marshal(ActionCallBack{
		Action: REGISTRATION_ACTION,
	})
	data[e.Resources.Buttonstext.Registration] = string(registration)

	searchRequest, _ := json.Marshal(ActionCallBack{
		Action: SEARCH_REQUEST_ACTION,
	})
	data[e.Resources.Buttonstext.SearchRequest] = string(searchRequest)

	placeAnAd, _ := json.Marshal(ActionCallBack{
		Action: SALE_ACTION,
	})
	data[e.Resources.Buttonstext.PlaceAnAd] = string(placeAnAd)

	rules, _ := json.Marshal(ActionCallBack{
		Action: RULES_ACTION,
	})
	data[e.Resources.Buttonstext.Rules] = string(rules)
	return CreateInlineKeyBoard(data, 1)
}

func CreateConcernButton(concerns []Concern) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(concerns); i++ {

		concern, _ := json.Marshal(Concern{Concern: concerns[i].Concern})
		concernData, _ := json.Marshal(SubsectionCallBack{
			Subsection: CHOOSE_CONCERN,
			Data:       string(concern),
		})

		data[concerns[i].Concern] = string(concernData)
	}

	return CreateInlineKeyBoard(data, 1)
}

func CreateAutoBrandButton(brands []Brand) tgbotapi.InlineKeyboardMarkup {

	data := make(map[string]string)
	for i := 0; i < len(brands); i++ {

		bJson, _ := json.Marshal(Brand{Brand: brands[i].Brand})
		brandData, _ := json.Marshal(SubsectionCallBack{
			Subsection: CHOOSE_BRAND,
			Data:       string(bJson),
		})

		data[brands[i].Brand] = string(brandData)
	}

	return CreateInlineKeyBoard(data, 1)
}
func CreateModelsButton(models []Model) tgbotapi.InlineKeyboardMarkup {

	data := make(map[string]string)
	for i := 0; i < len(models); i++ {

		mJson, _ := json.Marshal(Model{Model: models[i].Model})
		modelData, _ := json.Marshal(SubsectionCallBack{
			Subsection: CHOOSE_MODEL,
			Data:       string(mJson),
		})

		data[models[i].Model] = string(modelData)
	}

	return CreateInlineKeyBoard(data, 1)
}

func CreateEnginesButton(e []Engine) tgbotapi.InlineKeyboardMarkup {

	data := make(map[string]string)
	for i := 0; i < len(e); i++ {
		eJson, _ := json.Marshal(Engine{EngineName: e[i].EngineName})
		engineData, _ := json.Marshal(SubsectionCallBack{
			Subsection: CHOOSE_ENGINE,
			Data:       string(eJson),
		})
		data[e[i].EngineName] = string(engineData)
	}

	return CreateInlineKeyBoard(data, 1)
}

func CreateBoltPatternsButton(bp []BoltPattern) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(bp); i++ {
		bpJson, _ := json.Marshal(BoltPattern{BoltPatternSize: bp[i].BoltPatternSize})
		bpData, _ := json.Marshal(SubsectionCallBack{
			Subsection: CHOOSE_BOLT_PATTERN,
			Data:       string(bpJson),
		})
		data[bp[i].BoltPatternSize] = string(bpData)
	}

	return CreateInlineKeyBoard(data, 1)
}

func CreateRegionsButton(r []Region) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(r); i++ {
		rJson, _ := json.Marshal(Region{RegionName: r[i].RegionName})
		bpData, _ := json.Marshal(SubsectionCallBack{
			Subsection: CHOOSE_CITY,
			Data:       string(rJson),
		})
		data[r[i].RegionName] = string(bpData)
	}

	return CreateInlineKeyBoard(data, 1)
}

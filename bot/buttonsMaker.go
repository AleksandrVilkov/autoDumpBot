package bot

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"strconv"
)

func CreateMainButtons(e *Environment, update *tgbotapi.Update) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)

	registrationCallBack := CallBack{
		Action: REGISTRATION_ACTION,
		UserId: strconv.FormatInt(update.Message.Chat.ID, 10),
	}
	registrationToken := GetMD5Hash(registrationCallBack.toString())
	e.TempData[registrationToken] = registrationCallBack.toString()
	data[e.Resources.Buttonstext.Registration] = getButtonData(registrationToken)

	searchRequestCallBack := CallBack{
		Action: SEARCH_REQUEST_ACTION,
		UserId: strconv.FormatInt(update.Message.Chat.ID, 10),
	}
	searchRequestToken := GetMD5Hash(searchRequestCallBack.toString())
	e.TempData[searchRequestToken] = searchRequestCallBack.toString()
	data[e.Resources.Buttonstext.SearchRequest] = getButtonData(searchRequestToken)

	saleCallBack := CallBack{
		Action: SALE_ACTION,
		UserId: strconv.FormatInt(update.Message.Chat.ID, 10),
	}
	saleToken := GetMD5Hash(saleCallBack.toString())
	e.TempData[saleToken] = saleCallBack.toString()
	data[e.Resources.Buttonstext.PlaceAnAd] = getButtonData(saleToken)

	ruleCallback := CallBack{
		Action: RULES_ACTION,
		UserId: strconv.FormatInt(update.Message.Chat.ID, 10),
	}
	ruleToken := GetMD5Hash(ruleCallback.toString())
	e.TempData[ruleToken] = ruleCallback.toString()

	data[e.Resources.Buttonstext.Rules] = getButtonData(ruleToken)
	return CreateInlineKeyBoard(data, 1)
}

func CreateConcernButton(concerns []Concern, c *CallBack, e *Environment) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(concerns); i++ {
		c.CarData.Concern = concerns[i].Concern
		c.Subsection = CHOOSE_CONCERN
		token := GetMD5Hash(c.toString())
		e.TempData[token] = c.toString()
		data[concerns[i].Concern] = getButtonData(token)
	}
	return CreateInlineKeyBoard(data, 1)
}

func CreateAutoBrandButton(brands []Brand, e *Environment, cb *CallBack) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(brands); i++ {
		cb.CarData.Brand = brands[i].Brand
		cb.Subsection = CHOOSE_BRAND
		token := GetMD5Hash(cb.toString())
		e.TempData[token] = cb.toString()
		data[brands[i].Brand] = getButtonData(token)
	}

	return CreateInlineKeyBoard(data, 1)
}
func CreateModelsButton(models []Model, e *Environment, cb *CallBack) tgbotapi.InlineKeyboardMarkup {

	data := make(map[string]string)
	for i := 0; i < len(models); i++ {
		cb.CarData.Model = models[i].Model
		cb.Subsection = CHOOSE_MODEL
		token := GetMD5Hash(cb.toString())
		e.TempData[token] = cb.toString()
		data[models[i].Model] = getButtonData(token)
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

func getButtonData(token string) string {
	buttonData, _ := json.Marshal(ButtonData{
		Token: token,
	})
	return string(buttonData)
}

package bot

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"psa_dump_bot/bot/model"
	"strconv"
)

func CreateRegistrationButton(e *model.Environment, update *tgbotapi.Update) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)

	registrationCallBack := model.CallBack{
		Action: model.REGISTRATION_ACTION,
		UserId: strconv.FormatInt(update.Message.Chat.ID, 10),
	}
	registrationToken := GetMD5Hash(registrationCallBack.ToString())
	if !e.TempData.SaveTempData(registrationToken, &registrationCallBack) {
		log.Print("error save temp data")
	}
	data[e.Resources.Buttonstext.Registration] = getButtonData(registrationToken)
	return CreateInlineKeyBoard(data, 1)
}
func CreateMainButtons(e *model.Environment, update *tgbotapi.Update) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)

	searchRequestCallBack := model.CallBack{
		Action: model.SEARCH_REQUEST_ACTION,
		UserId: strconv.FormatInt(update.Message.Chat.ID, 10),
	}
	searchRequestToken := GetMD5Hash(searchRequestCallBack.ToString())
	if !e.TempData.SaveTempData(searchRequestToken, &searchRequestCallBack) {
		log.Print("error save temp data")
	}
	data[e.Resources.Buttonstext.SearchRequest] = getButtonData(searchRequestToken)

	saleCallBack := model.CallBack{
		Action: model.SALE_ACTION,
		UserId: strconv.FormatInt(update.Message.Chat.ID, 10),
	}
	saleToken := GetMD5Hash(saleCallBack.ToString())
	if !e.TempData.SaveTempData(saleToken, &saleCallBack) {
		log.Print("error save temp data")
	}
	data[e.Resources.Buttonstext.PlaceAnAd] = getButtonData(saleToken)

	ruleCallback := model.CallBack{
		Action: model.RULES_ACTION,
		UserId: strconv.FormatInt(update.Message.Chat.ID, 10),
	}
	ruleToken := GetMD5Hash(ruleCallback.ToString())
	if !e.TempData.SaveTempData(ruleToken, &ruleCallback) {
		log.Print("error save temp data")
	}

	data[e.Resources.Buttonstext.Rules] = getButtonData(ruleToken)
	return CreateInlineKeyBoard(data, 1)
}

func CreateConcernButton(concerns []model.Concern, c *model.CallBack, e *model.Environment) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(concerns); i++ {
		c.CarData.Concern = concerns[i].Concern
		c.Subsection = model.CHOOSE_CONCERN
		token := GetMD5Hash(c.ToString())
		if !e.TempData.SaveTempData(token, c) {
			log.Print("error save temp data")
		}
		data[concerns[i].Concern] = getButtonData(token)
	}
	return CreateInlineKeyBoard(data, 1)
}

func CreateAutoBrandButton(brands []model.Brand, e *model.Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(brands); i++ {
		cb.CarData.Brand = brands[i].Brand
		cb.Subsection = model.CHOOSE_BRAND
		token := GetMD5Hash(cb.ToString())
		if !e.TempData.SaveTempData(token, cb) {
			log.Print("error save temp data")
		}
		data[brands[i].Brand] = getButtonData(token)
	}

	return CreateInlineKeyBoard(data, 1)
}
func CreateModelsButton(models []model.Model, e *model.Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup {

	data := make(map[string]string)
	for i := 0; i < len(models); i++ {
		cb.CarData.Model = models[i].Model
		cb.Subsection = model.CHOOSE_MODEL
		token := GetMD5Hash(cb.ToString())
		if !e.TempData.SaveTempData(token, cb) {
			log.Print("error save temp data")
		}
		data[models[i].Model] = getButtonData(token)
	}

	return CreateInlineKeyBoard(data, 1)
}

func CreateEnginesButton(en []model.Engine, e *model.Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup {

	data := make(map[string]string)
	for i := 0; i < len(en); i++ {
		cb.CarData.EngineName = en[i].EngineName
		cb.Subsection = model.CHOOSE_ENGINE
		token := GetMD5Hash(cb.ToString())
		if !e.TempData.SaveTempData(token, cb) {
			log.Print("error save temp data")
		}
		data[en[i].EngineName] = getButtonData(token)
	}

	return CreateInlineKeyBoard(data, 1)
}

func CreateBoltPatternsButton(bp []model.BoltPattern, e *model.Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(bp); i++ {
		cb.CarData.BoltPatternSize = bp[i].BoltPatternSize
		cb.Subsection = model.CHOOSE_BOLT_PATTERN
		token := GetMD5Hash(cb.ToString())
		if !e.TempData.SaveTempData(token, cb) {
			log.Print("error save temp data")
		}
		data[bp[i].BoltPatternSize] = getButtonData(token)
	}

	return CreateInlineKeyBoard(data, 1)
}

func CreateRegionsButton(r []model.Region, e *model.Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(r); i++ {
		cb.UserData.RegionName = r[i].RegionName
		cb.UserData.RegionId = r[i].Id
		cb.Subsection = model.CHOOSE_CITY
		token := GetMD5Hash(cb.ToString())
		if !e.TempData.SaveTempData(token, cb) {
			log.Print("error save temp data")
		}
		data[r[i].RegionName] = getButtonData(token)
	}

	return CreateInlineKeyBoard(data, 1)
}

func getButtonData(token string) string {
	buttonData, _ := json.Marshal(model.ButtonData{
		Token: token,
	})
	return string(buttonData)
}

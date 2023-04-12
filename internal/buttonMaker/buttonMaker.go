package buttonMaker

import (
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"psa_dump_bot/bot"
	"psa_dump_bot/model"
	"strconv"
)

type ButtonMaker struct {
}

func NewButtonMaker() *ButtonMaker {
	return &ButtonMaker{}
}
func (b *ButtonMaker) CreateRegistrationButton(e *bot.Environment, update *tgbotapi.Update) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)

	registrationCallBack := model.CallBack{
		Action: model.REGISTRATION_ACTION,
		UserId: strconv.FormatInt(update.Message.Chat.ID, 10),
	}
	registrationToken := bot.GetMD5Hash(registrationCallBack.ToString())
	if !e.TempData.SaveTempData(registrationToken, &registrationCallBack) {
		log.Print("error save temp data")
	}
	data[e.Resources.Buttonstext.Registration] = b.getButtonData(registrationToken)
	return bot.CreateInlineKeyBoard(data, 1)
}
func (b *ButtonMaker) CreateMainButtons(e *bot.Environment, update *tgbotapi.Update) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)

	searchRequestCallBack := model.CallBack{
		Action: model.SEARCH_REQUEST_ACTION,
		UserId: strconv.FormatInt(update.Message.Chat.ID, 10),
	}
	searchRequestToken := bot.GetMD5Hash(searchRequestCallBack.ToString())
	if !e.TempData.SaveTempData(searchRequestToken, &searchRequestCallBack) {
		log.Print("error save temp data")
	}
	data[e.Resources.Buttonstext.SearchRequest] = b.getButtonData(searchRequestToken)

	saleCallBack := model.CallBack{
		Action: model.SALE_ACTION,
		UserId: strconv.FormatInt(update.Message.Chat.ID, 10),
	}
	saleToken := bot.GetMD5Hash(saleCallBack.ToString())
	if !e.TempData.SaveTempData(saleToken, &saleCallBack) {
		log.Print("error save temp data")
	}
	data[e.Resources.Buttonstext.PlaceAnAd] = b.getButtonData(saleToken)

	ruleCallback := model.CallBack{
		Action: model.RULES_ACTION,
		UserId: strconv.FormatInt(update.Message.Chat.ID, 10),
	}
	ruleToken := bot.GetMD5Hash(ruleCallback.ToString())
	if !e.TempData.SaveTempData(ruleToken, &ruleCallback) {
		log.Print("error save temp data")
	}

	data[e.Resources.Buttonstext.Rules] = b.getButtonData(ruleToken)
	return bot.CreateInlineKeyBoard(data, 1)
}

func (b *ButtonMaker) CreateUniversalButton(c *model.CallBack, e *bot.Environment) {
	data := make(map[string]string)
	c.Subsection = model.UNEVERSAL
	token := bot.GetMD5Hash(c.ToString())
	data[e.Resources.Buttonstext.Universal] = b.getButtonData(token)
}

func (b *ButtonMaker) CreateConcernButton(concerns []model.Concern, c *model.CallBack, e *bot.Environment) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(concerns); i++ {
		c.CarData.Concern = concerns[i].Concern
		c.Subsection = model.CHOOSE_CONCERN
		token := bot.GetMD5Hash(c.ToString())
		if !e.TempData.SaveTempData(token, c) {
			log.Print("error save temp data")
		}
		data[concerns[i].Concern] = b.getButtonData(token)
	}
	return bot.CreateInlineKeyBoard(data, 1)
}

func (b *ButtonMaker) CreateAutoBrandButton(brands []model.Brand, e *bot.Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(brands); i++ {
		cb.CarData.Brand = brands[i].Brand
		cb.Subsection = model.CHOOSE_BRAND
		token := bot.GetMD5Hash(cb.ToString())
		if !e.TempData.SaveTempData(token, cb) {
			log.Print("error save temp data")
		}
		data[brands[i].Brand] = b.getButtonData(token)
	}

	return bot.CreateInlineKeyBoard(data, 1)
}
func (b *ButtonMaker) CreateModelsButton(models []model.Model,
	e *bot.Environment,
	cb *model.CallBack, subsection model.CallbackSubsection) tgbotapi.InlineKeyboardMarkup {

	data := make(map[string]string)
	for i := 0; i < len(models); i++ {
		cb.CarData.Model = models[i].Model
		cb.Subsection = subsection
		token := bot.GetMD5Hash(cb.ToString())
		if !e.TempData.SaveTempData(token, cb) {
			log.Print("error save temp data")
		}
		data[models[i].Model] = b.getButtonData(token)
	}

	return bot.CreateInlineKeyBoard(data, 1)
}

func (b *ButtonMaker) CreateEnginesButton(en []model.Engine, e *bot.Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup {

	data := make(map[string]string)
	for i := 0; i < len(en); i++ {
		cb.CarData.EngineName = en[i].EngineName
		cb.Subsection = model.CHOOSE_ENGINE
		token := bot.GetMD5Hash(cb.ToString())
		if !e.TempData.SaveTempData(token, cb) {
			log.Print("error save temp data")
		}
		data[en[i].EngineName] = b.getButtonData(token)
	}

	return bot.CreateInlineKeyBoard(data, 1)
}

func (b *ButtonMaker) CreateBoltPatternsButton(bp []model.BoltPattern, e *bot.Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(bp); i++ {
		cb.CarData.BoltPatternSize = bp[i].BoltPatternSize
		cb.Subsection = model.CHOOSE_BOLT_PATTERN
		token := bot.GetMD5Hash(cb.ToString())
		if !e.TempData.SaveTempData(token, cb) {
			log.Print("error save temp data")
		}
		data[bp[i].BoltPatternSize] = b.getButtonData(token)
	}

	return bot.CreateInlineKeyBoard(data, 1)
}

func (b *ButtonMaker) CreateRegionsButton(r []model.Region, e *bot.Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup {
	data := make(map[string]string)
	for i := 0; i < len(r); i++ {
		cb.UserData.RegionName = r[i].RegionName
		cb.UserData.RegionId = r[i].Id
		cb.Subsection = model.CHOOSE_CITY
		token := bot.GetMD5Hash(cb.ToString())
		if !e.TempData.SaveTempData(token, cb) {
			log.Print("error save temp data")
		}
		data[r[i].RegionName] = b.getButtonData(token)
	}

	return bot.CreateInlineKeyBoard(data, 1)
}

func (b *ButtonMaker) getButtonData(token string) string {
	buttonData, _ := json.Marshal(model.ButtonData{
		Token: token,
	})
	return string(buttonData)
}

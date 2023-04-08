package messageProcessor

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"psa_dump_bot/bot"
)

func CreateWelcomeMsg(e *bot.Environment) string {
	return e.Resources.Msgs.WelcomeMessage
}
func CreateWelcomeRegisteredMsg(e *bot.Environment) string {
	return e.Resources.Msgs.WelcomeRegistered
}
func CreateErrAuthMsg(e *bot.Environment) string {
	return e.Resources.Errors.AuthError + " \nПодпишись: \n" + e.Config.ValidateData.ChannelUrl
}

func CreateConcernMsgForReg(e *bot.Environment) string {
	return e.Resources.Msgs.Registration.ChoiceConcern
}

func CreateBrandMsgForReg(e *bot.Environment) string {
	return e.Resources.Msgs.Registration.CarBrandEnter
}

func CreateModelMsgForReg(e *bot.Environment) string {
	return e.Resources.Msgs.Registration.CarModelEnter
}
func CreateEngineMsgForReg(e *bot.Environment) string {
	return e.Resources.Msgs.Registration.CarEngineEnter
}

func CreateBoltPatternMsg(e *bot.Environment) string {
	return e.Resources.Msgs.Registration.CarBoltPatternEnter
}

func CreateRegionMsg(e *bot.Environment) string {
	return e.Resources.Msgs.Registration.RegionEnter
}

func CreateConcernMsgForSearch(e *bot.Environment) string {
	return e.Resources.Msgs.Search.ChoiceConcern
}

func CreateBrandMsgForSearch(e *bot.Environment) string {
	return e.Resources.Msgs.Search.CarBrandEnter
}
func CreateModelMsgForSearch(e *bot.Environment) string {
	return e.Resources.Msgs.Search.CarModelEnter
}
func CreateErrorMsg(update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, e.Resources.Errors.CommonError)
}

package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"psa_dump_bot/model"
)

type ButtonMaker interface {
	CreateRegistrationButton(e *Environment, update *tgbotapi.Update) tgbotapi.InlineKeyboardMarkup
	CreateMainButtons(e *Environment, update *tgbotapi.Update) tgbotapi.InlineKeyboardMarkup
	CreateConcernButton(concerns []model.Concern, c *model.CallBack, e *Environment) tgbotapi.InlineKeyboardMarkup
	CreateAutoBrandButton(brands []model.Brand, e *Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup
	CreateModelsButton(models []model.Model, e *Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup
	CreateEnginesButton(en []model.Engine, e *Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup
	CreateBoltPatternsButton(bp []model.BoltPattern, e *Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup
	CreateRegionsButton(r []model.Region, e *Environment, cb *model.CallBack) tgbotapi.InlineKeyboardMarkup
}

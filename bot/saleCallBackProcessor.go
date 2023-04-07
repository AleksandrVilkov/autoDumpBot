package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"psa_dump_bot/bot/model"
)

func saleProcessor(update *tgbotapi.Update, e *model.Environment) {
	//case PLACE_AN_AD:
	//	msg = createPlaceAnAdWelcomeResp(update, e)
	//	userTemp := e.TempData[strconv.Itoa(update.CallbackQuery.From.ID)]
	//	userTemp.LastCommand = PLACE_AN_AD
}

//func createPlaceAnAdWelcomeResp(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig {
//	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), e.Resources.Msgs.StartSale)
//	//concerns := e.Storage.GetConcerns()
//	concerns := make([]Concern, 1)
//	concerns[0] = Concern{
//		Concern: "PSA",
//	}
//	msg.ReplyMarkup = CreateConcernButton(concerns)
//	return msg
//}

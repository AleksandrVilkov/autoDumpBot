package messageProcessor

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"psa_dump_bot/bot"
	"reflect"
	"strconv"
)

type MessageProcessor struct {
}

func NewMessageProcessor() *MessageProcessor {
	return &MessageProcessor{}
}
func (m *MessageProcessor) StartMessageProcessor(update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig

	if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

		switch update.Message.Text {
		case e.Config.Commands.Start:
			isRegistered := e.Storage.CheckUser(strconv.FormatInt(update.Message.Chat.ID, 10))

			if isRegistered {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, CreateWelcomeRegisteredMsg(e))
				msg.ReplyMarkup = e.ButtonMaker.CreateMainButtons(e, update)
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, CreateWelcomeMsg(e))
				msg.ReplyMarkup = e.ButtonMaker.CreateRegistrationButton(e, update)
			}

		default:
			msg = m.CreateError(update, e)
		}
	} else {
		msg = m.CreateError(update, e)
	}

	return msg
}

func (m *MessageProcessor) CreateError(update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	return CreateErrorMsg(update, e)
}

func (m *MessageProcessor) CreateErrAuthMsg(update *tgbotapi.Update, e *bot.Environment) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, CreateErrAuthMsg(e))
}

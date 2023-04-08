package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type MessageProcessor interface {
	StartMessageProcessor(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig
	CreateError(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig
	CreateErrAuthMsg(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig
}

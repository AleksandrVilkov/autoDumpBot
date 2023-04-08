package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type CallBackProcessor interface {
	StartCallBackProcessor(update *tgbotapi.Update, e *Environment) tgbotapi.MessageConfig
}

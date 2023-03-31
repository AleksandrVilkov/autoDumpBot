package bot

import (
	"errors"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func validateUser(bot *tgbotapi.BotAPI, update *tgbotapi.Update, conf *Config) error {
	var userID int

	if update.Message == nil {
		userID = update.CallbackQuery.From.ID
	} else {
		userID = update.Message.From.ID
	}

	chatMember, e := bot.GetChatMember(tgbotapi.ChatConfigWithUser{
		ChatID: conf.ValidateData.ChannelID,
		UserID: userID,
	})
	if e != nil || chatMember.Status == KICKED || chatMember.Status == LEFT {
		return errors.New("The user is not in the channel!")
	}
	return nil
}

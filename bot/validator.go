package bot

import (
	"errors"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func validateUser(bot *tgbotapi.BotAPI, update *tgbotapi.Update, conf *config) error {
	chatMember, e := bot.GetChatMember(tgbotapi.ChatConfigWithUser{
		ChatID: conf.ValidateData.ChannelID,
		UserID: update.Message.From.ID,
	})
	if e != nil || chatMember.Status == KICKED || chatMember.Status == LEFT {
		return errors.New("The user is not in the channel!")
	}
	return nil
}

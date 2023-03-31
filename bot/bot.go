package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"gopkg.in/yaml.v3"
	"os"
)

const PARAMS_PATH = "/home/vilkov/GolandProjects/psa_dump_bot/config/config.yaml"

func StartBot(storage *Storage) {
	paramsFile, err := os.ReadFile(PARAMS_PATH)
	CheckFatalError(err)

	tempRegister := make(map[string]TempUserData)

	var conf Config
	err = yaml.Unmarshal(paramsFile, &conf)
	CheckFatalError(err)

	bot, err := tgbotapi.NewBotAPI(conf.Token)
	CheckFatalError(err)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		var msg tgbotapi.MessageConfig
		if validateUser(bot, &update, &conf) != nil {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, CreateErrAuthMsg(conf.ValidateData.ChannelUrl))
			_, err := bot.Send(msg)
			CheckFatalError(err)
			return
		}

		if update.CallbackQuery != nil {
			msg = CallbackProcessing(&update, tempRegister, storage)
		}
		if update.Message != nil {
			msg = MsgProcessing(&update, conf, tempRegister)
		}

		_, err := bot.Send(msg)
		CheckError(err)
	}
}

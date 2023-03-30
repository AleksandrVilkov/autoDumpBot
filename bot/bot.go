package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"gopkg.in/yaml.v3"
	"os"
)

const PARAMS_PATH = "/home/vilkov/GolandProjects/psa_dump_bot/config/config.yaml"

type tempUserData struct {
	User struct {
		Id string
	}
	Action struct {
		MainAction  string
		LastCommand string
	}
	CarData struct {
		CarBrand    string
		CarModel    string
		CarEngine   string
		BoltPattern string
	}
	SaleData struct {
	}
	SubscriptionData struct {
	}
}

func StartBot() {
	paramsFile, err := os.ReadFile(PARAMS_PATH)
	checkError(err)

	tempRegister := make(map[string]tempUserData)

	var conf config
	err = yaml.Unmarshal(paramsFile, &conf)
	checkFatalError(err)

	bot, err := tgbotapi.NewBotAPI(conf.Token)
	checkFatalError(err)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		processing(&update, bot, conf, tempRegister)
	}
}

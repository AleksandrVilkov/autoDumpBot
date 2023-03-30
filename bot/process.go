package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"psa_dump_bot/internal"
	"reflect"
	"strconv"
)

func processing(update *tgbotapi.Update, bot *tgbotapi.BotAPI, conf config, temp map[string]tempUserData) {

	var msg tgbotapi.MessageConfig
	var tempData tempUserData

	if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

		if validateUser(bot, update, &conf) != nil {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, internal.CreateErrAuthMsg(conf.ValidateData.ChannelUrl))
			_, err := bot.Send(msg)
			checkError(err)
			return
		}

		tempData.User.Id = strconv.FormatInt(update.Message.Chat.ID, 10)
		switch update.Message.Text {
		case conf.Commands.Start:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, internal.CreateWelcomeMsg())
			tempData.Action.LastCommand = conf.Commands.Start
		case conf.Commands.Rules:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, internal.CreateRulesMsg())
			tempData.Action.LastCommand = conf.Commands.Rules
		case conf.Commands.Registration:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, internal.CreatePreRegistrationMsg())
			tempData.Action.LastCommand = conf.Commands.Registration
			tempData.Action.MainAction = conf.Commands.Registration
		case conf.Commands.Sale:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, internal.CreateStartSaleMsg())
			tempData.Action.LastCommand = conf.Commands.Sale
			tempData.Action.MainAction = conf.Commands.Sale
		case conf.Commands.Subscription:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, internal.CreatePreRegistrationMsg())
			tempData.Action.LastCommand = conf.Commands.Subscription
			tempData.Action.MainAction = conf.Commands.Subscription
		default:
			text, newTemp := secondaryCommandProcessing(update.Message.Text,
				temp[strconv.FormatInt(update.Message.Chat.ID, 10)],
				conf)
			tempData = newTemp
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, text)
		}
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Use the words for search.")

	}

	temp[strconv.FormatInt(update.Message.Chat.ID, 10)] = tempData
	_, err := bot.Send(msg)
	checkError(err)
}

func secondaryCommandProcessing(text string, tempData tempUserData, conf config) (string, tempUserData) {
	switch tempData.Action.MainAction {
	case conf.Commands.Registration:
		tempData.CarData.CarBrand = text
		tempData.Action.LastCommand = conf.InternalCommands.EnterCarBrand
		return registrationProcessing(text, tempData, conf)
	case conf.Commands.Sale:
		return saleProcessing(text, tempData, conf)
	case conf.Commands.Subscription:
		return subscriptionProcessing(text, tempData, conf)
	default:
		return internal.CreateError(), tempData

	}
}

func registrationProcessing(text string, tempData tempUserData, conf config) (string, tempUserData) {
	switch tempData.Action.LastCommand {
	case conf.Commands.Registration:
		tempData.CarData.CarBrand = text
		tempData.Action.LastCommand = conf.InternalCommands.EnterCarBrand
		return internal.CreateSuccessCarBrandEnter(), tempData
	case conf.InternalCommands.EnterCarBrand:
		tempData.CarData.CarModel = text
		tempData.Action.LastCommand = conf.InternalCommands.EnterCarModel
		return internal.CreateSuccessCarModelEnter(), tempData
	case conf.InternalCommands.EnterCarModel:
		tempData.CarData.CarEngine = text
		tempData.Action.LastCommand = conf.InternalCommands.EnterCarEngine
		return internal.CreateSuccessCarEngineEnter(), tempData
	case conf.InternalCommands.EnterCarEngine:
		tempData.CarData.BoltPattern = text
		tempData.Action.LastCommand = conf.InternalCommands.EnterCarEngine
		return internal.CreatesReservation(), tempUserData{}
	default:
		return internal.CreateError(), tempData
	}
}

func subscriptionProcessing(text string, tempData tempUserData, conf config) (string, tempUserData) {
	return internal.CreateError(), tempData

}

func saleProcessing(text string, tempData tempUserData, conf config) (string, tempUserData) {
	return internal.CreateError(), tempData

}

package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"psa_dump_bot/bot"
	"psa_dump_bot/internal/buttonMaker"
	"psa_dump_bot/internal/callbackProceccor"
	"psa_dump_bot/internal/messageProcessor"
	postgreSQL "psa_dump_bot/stogage"
)

const (
	PARAMS_PATH = "/home/vilkov/GolandProjects/autoDumpBot/config/config.yaml"
	RES_PATH    = "/home/vilkov/GolandProjects/autoDumpBot/resources.yaml"
)

func main() {
	paramsFile, err := os.ReadFile(PARAMS_PATH)

	if err != nil {
		log.Fatal(err)
	}
	var conf bot.Config
	err = yaml.Unmarshal(paramsFile, &conf)

	e := bot.Environment{
		Config:            &conf,
		Storage:           postgreSQL.NewStorage(),
		Resources:         getResources(),
		TempData:          postgreSQL.NewStorage(),
		CallBackProcessor: callbackProceccor.NewCallBackProcessor(),
		MessageProcessor:  messageProcessor.NewMessageProcessor(),
		ButtonMaker:       buttonMaker.NewButtonMaker(),
	}
	bot.StartBot(e)
}

func getResources() *bot.Resources {
	paramsFile, err := os.ReadFile(RES_PATH)
	if err != nil {
		log.Fatal(err)
	}
	var res bot.Resources
	unmarshalErr := yaml.Unmarshal(paramsFile, &res)

	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return &res
}

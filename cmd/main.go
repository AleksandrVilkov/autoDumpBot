package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"psa_dump_bot/bot"
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

	s := postgreSQL.NewStorage()
	e := bot.Environment{
		Config:    &conf,
		Storage:   s,
		Resources: getResources(),
	}

	bot.StartBot(&e)
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

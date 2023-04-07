package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"psa_dump_bot/bot"
	"psa_dump_bot/bot/model"
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
	var conf model.Config
	err = yaml.Unmarshal(paramsFile, &conf)

	e := model.Environment{
		Config:    &conf,
		Storage:   postgreSQL.NewStorage(),
		Resources: getResources(),
		TempData:  postgreSQL.NewStorage(),
	}
	bot.StartBot(&e)
}

func getResources() *model.Resources {
	paramsFile, err := os.ReadFile(RES_PATH)
	if err != nil {
		log.Fatal(err)
	}
	var res model.Resources
	unmarshalErr := yaml.Unmarshal(paramsFile, &res)

	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return &res
}

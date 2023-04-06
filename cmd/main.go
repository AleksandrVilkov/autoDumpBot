package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"psa_dump_bot/bot"
	postgreSQL "psa_dump_bot/stogage"
)

const (
	PARAMS_PATH = "/home/vilkov/GolandProjects/psa_dump_bot/config/config.yaml"
	RES_PATH    = "/home/vilkov/GolandProjects/psa_dump_bot/resources.yaml"
)

func main() {
	paramsFile, err := os.ReadFile(PARAMS_PATH)

	if err != nil {
		log.Fatal(err)
	}
	var conf bot.Config
	err = yaml.Unmarshal(paramsFile, &conf)

	e := bot.Environment{
		Config:    &conf,
		Storage:   postgreSQL.NewStorage(),
		Resources: getResources(),
	}

	e.TempData = make(map[string]bot.TempData)
	StartBot(&e)
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

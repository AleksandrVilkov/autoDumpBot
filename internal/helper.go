package internal

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const RES_PATH = "/home/vilkov/GolandProjects/psa_dump_bot/resources.yaml"

func getResources() *resources {
	paramsFile, err := os.ReadFile(RES_PATH)
	checkError(err)
	var res resources
	unmarshalErr := yaml.Unmarshal(paramsFile, &res)
	checkError(unmarshalErr)
	return &res
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

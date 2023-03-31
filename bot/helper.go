package bot

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const RES_PATH = "/home/vilkov/GolandProjects/psa_dump_bot/resources.yaml"

func GetResources() *Resources {
	paramsFile, err := os.ReadFile(RES_PATH)
	CheckError(err)
	var res Resources
	unmarshalErr := yaml.Unmarshal(paramsFile, &res)
	CheckError(unmarshalErr)
	return &res
}

func CheckError(e error) {
	if e != nil {
		log.Print(e)
	}
}
func CheckFatalError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

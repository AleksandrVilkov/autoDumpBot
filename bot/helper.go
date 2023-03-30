package bot

import "log"

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func checkFatalError(err error) {
	if err != nil {
		panic(err)
	}
}

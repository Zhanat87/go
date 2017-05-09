package helpers

import (
	"fmt"
	"log"
	"os"
)

func FailOnError(err error, msg string, sendErrorEmail bool) {
	if err != nil {
		saveToFile(fmt.Sprintf("%s: %s", msg, err), "golang_error", sendErrorEmail)
		if sendErrorEmail {
			SendErrorEmail(os.Getenv("GO_LANG_ERROR_SUBJECT"), fmt.Sprintf("%s: %s", msg, err))
		}
		panic(err)
	}
}

func LogInfo(msg string, sendErrorEmail bool) {
	saveToFile(msg, "golang_info", sendErrorEmail)
}

func saveToFile(msg string, filename string, sendErrorEmail bool) {
	f, err := os.OpenFile("logs/" + filename + ".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		if sendErrorEmail {
			SendErrorEmail(os.Getenv("GO_LANG_ERROR_SUBJECT"), fmt.Sprintf("error opening file: %v", err))
		}
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Print(msg)
}

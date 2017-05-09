package helpers

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"net/url"
	"bytes"
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
	f, err := os.OpenFile(os.Getenv("APP_DIR") + "/logs/" + filename + ".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
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

func LogError(error error) {
	request_url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", os.Getenv("TELEGRAM_API_BOT_TOKEN"))
	form := url.Values{
		"text": {error.Error()},
		"chat_id": {os.Getenv("TELEGRAM_MY_USER_CHAT_ID")},
	}

	// func Post(url string, bodyType string, body io.Reader) (resp *Response, err error) {
	body := bytes.NewBufferString(form.Encode())
	rsp, err := http.Post(request_url, "application/x-www-form-urlencoded", body)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
}

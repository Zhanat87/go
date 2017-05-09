package helpers

import (
	"crypto/tls"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendErrorEmail(sbj, msg string) (bool, error) {
	return SendEmail(os.Getenv("MAIL_TO_ADDRESS"), sbj, msg)
}

/*
golang optional parameters
@link http://stackoverflow.com/questions/2032149/optional-parameters
@link http://stackoverflow.com/questions/32568977/golang-pass-nil-as-optional-argument-to-a-function
 */
func SendEmail(to, sbj, msg string) (bool, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_FROM_ADDRESS"), os.Getenv("MAIL_FROM_NAME"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", sbj)
	//m.SetBody("text/plain", msg)
	m.SetBody("text/html", msg)

	mailPort, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		FailOnError(err, "send email: get port", false)
		return false, err
	}
	d := gomail.NewDialer(os.Getenv("MAIL_HOST"), mailPort, os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		FailOnError(err, "send email", false)
		return false, err
	}
	return true, nil
}

package services

import (
	"net/smtp"
	"os"
)

func SendMail(recipient, from, subject, message string) error {
	emailHost := os.Getenv("EMAIL_HOST")
	auth := smtp.PlainAuth("", os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_USERPASSWORD"), emailHost)

	to := []string{recipient}
	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: " + subject + "!\r\n" +
		"\r\n" + message)

	if err := smtp.SendMail(emailHost+":"+os.Getenv("EMAIL_PORT"), auth, from, to, msg); err != nil {
		panic(err)
	}

	return nil
}

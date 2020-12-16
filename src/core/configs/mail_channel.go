package configs

import (
	"ginapp/core/emails"
	"ginapp/core/services"
	"os"
	"strconv"
)

var number, _ = strconv.Atoi(os.Getenv("CHANNEL_SIZE_DEFAULT"))

var EmailChannel = make(chan *emails.BasicMail, number)

func InitMailChannel() {
	go func() {
		for mailInfo := range EmailChannel {
			services.SendMail(mailInfo.Recipient, mailInfo.From, mailInfo.Subject, mailInfo.Message)
		}
	}()
}

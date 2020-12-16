package emails

import (
	"ginapp/core/emails"
	"ginapp/core/services"
)

type RegisterMail struct {
	emails.Email
	Recipient string
	From      string
	Subject   string
	Message   string
}

func (mail *RegisterMail) Handle() {
	services.SendMail(mail.Recipient, mail.From, mail.Subject, mail.Message)
}

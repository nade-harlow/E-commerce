package repositories

type MailRepository interface {
	NewMailgunRepository()
	AddTemplate() error
	SendMessageWithTemplate(email, subject string) (string, error)
	SendMail(email, subject, body string) error
}

type TwilloRepository interface {
	SendSms(phone string, message string) error
}

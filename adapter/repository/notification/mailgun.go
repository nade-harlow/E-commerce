package notification

import (
	"context"
	"github.com/mailgun/mailgun-go/v4"
	"os"
	"time"
)

var domain = os.Getenv("MG_DOMAIN")
var apiKey = os.Getenv("MG_PUBLIC_API_KEY")

func AddTemplate(domain, apiKey string) error {
	mg := mailgun.NewMailgun(domain, apiKey)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	return mg.CreateTemplate(ctx, &mailgun.Template{
		Name: "verify.account",
		Version: mailgun.TemplateVersion{
			Template: ``,
			Engine:   mailgun.TemplateEngineGo,
			Tag:      "v1",
		},
	})
}

func SendMessageWithTemplate(email, subject string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	m := mg.NewMessage(os.Getenv("MG_EMAIL_FROM"), subject, "")
	m.SetTemplate("verify.account")
	if err := m.AddRecipient(email); err != nil {
		return "", err
	}

	//m.AddVariable("link", "www.instagram.com")

	mes, _, err := mg.Send(ctx, m)
	return mes, err
}

func SendSimpleMessage(email, subject, body string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(os.Getenv("MG_EMAIL_FROM"), subject, body, email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}

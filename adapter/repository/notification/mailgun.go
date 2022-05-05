package notification

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"os"
	"time"
)

// MailgunRepository is a repository for sending emails
type MailgunRepository struct {
	client *mailgun.MailgunImpl
}

// NewMailgunRepository creates a new MailgunRepository
func (mg *MailgunRepository) NewMailgunRepository() {
	fmt.Println(os.Getenv("MG_DOMAIN"))
	mg.client = mailgun.NewMailgun(os.Getenv("MG_DOMAIN"), os.Getenv("MG_API_KEY"))
	//fmt.Println("domooo", mg.client.Domain())
}

func (mg *MailgunRepository) AddTemplate() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	return mg.client.CreateTemplate(ctx, &mailgun.Template{
		Name: "**template name**",
		Version: mailgun.TemplateVersion{
			Template: ``,
			Engine:   mailgun.TemplateEngineGo,
			Tag:      "v1",
		},
	})
}

func (mg *MailgunRepository) SendMessageWithTemplate(email, subject string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	m := mg.client.NewMessage(os.Getenv("MG_EMAIL_FROM"), subject, "")
	m.SetTemplate("**template name**")
	if err := m.AddRecipient(email); err != nil {
		return "", err
	}

	mes, _, err := mg.client.Send(ctx, m)
	return mes, err
}

func (mg *MailgunRepository) SendMail(email, subject, body string) error {
	mg.client = mailgun.NewMailgun(os.Getenv("MG_DOMAIN"), os.Getenv("MG_API_KEY"))

	m := mg.client.NewMessage(os.Getenv("MG_EMAIL_FROM"), subject, body, email)
	m.SetHtml(body)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	_, id, err := mg.client.Send(ctx, m)

	log.Println("Message ID: ", id)
	return err
}

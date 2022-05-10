package notification

import (
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
	"os"
)

type TwilloRepository struct {
	client *twilio.RestClient
}

func (t *TwilloRepository) NewTwillo() {
	client := twilio.NewRestClient()
	t.client = client
}

func (t *TwilloRepository) SendSms(phone string, message string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody(message)

	_, err := t.client.ApiV2010.CreateMessage(params)
	if err != nil {
		log.Println(err.Error())
		return err
	} else {
		log.Println("SMS sent successfully!")
		return nil
	}
}

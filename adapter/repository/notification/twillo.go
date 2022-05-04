package notification

import (
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
	"os"
)

func SendSms(phone string, message string) error {
	client := twilio.NewRestClient()
	params := &openapi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody(message)

	_, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		log.Println(err.Error())
		return err
	} else {
		log.Println("SMS sent successfully!")
		return nil
	}
}

package twilio

import (
	"fmt"
	"os"
	"testing"
)

func TestSendSMS(t *testing.T) {

	client := NewClient(os.Getenv("TWILIO_SID"), os.Getenv("TWILIO_KEY"), os.Getenv("TWILIO_FROM"))
	r, err := client.SendSMS("+62847123456", "123456")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("%+v\n", r)

	r, err = client.SendSMS("+6285317747535", "123456")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("%+v\n", r)
}

func TestLookup(t *testing.T) {

	client := NewClient(os.Getenv("TWILIO_SID"), os.Getenv("TWILIO_KEY"), os.Getenv("TWILIO_FROM"))
	r, err := client.Lookup("1112345678")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("%+v\n", r)

	r, err = client.Lookup("6284712345678")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("%+v\n", r)

	r, err = client.Lookup("6281112345678")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("%+v\n", r)

}

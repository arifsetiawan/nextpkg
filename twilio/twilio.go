package twilio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Client ...
type Client struct {
	AccountSID   string
	AuthToken    string
	FromNumber   string
	SMSAPIURL    string
	LookupAPIURL string
}

// ErrorResponse ...
type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

// SMSResponse ...
type SMSResponse struct {
	SID         string `json:"sid,omitempty"`
	Status      string `json:"status,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
}

// LookupResponse ...
type LookupResponse struct {
	CountryCode    string `json:"country_code,omitempty"`
	PhoneNumber    string `json:"phone_number,omitempty"`
	NationalFormat string `json:"national_format,omitempty"`
}

// NewClient is
func NewClient(sid string, token string, from string) *Client {
	if len(sid) == 0 && len(token) == 0 && len(from) == 0 {
		log.Fatal("Invalid twilio SID and Token. Set SID and Token with correct value")
	}

	return &Client{
		AccountSID:   sid,
		AuthToken:    token,
		FromNumber:   from,
		SMSAPIURL:    "https://api.twilio.com/2010-04-01/Accounts/" + sid + "/Messages.json",
		LookupAPIURL: "https://lookups.twilio.com/v1/PhoneNumbers/",
	}
}

// SendSMS ...
func (t *Client) SendSMS(to string, otpCode string) (*SMSResponse, error) {

	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", t.FromNumber)
	msgData.Set("Body", "Your verification code is "+otpCode)
	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", t.SMSAPIURL, &msgDataReader)
	req.SetBasicAuth(t.AccountSID, t.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)

	// Get body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		errorResponse := &ErrorResponse{}
		err := json.Unmarshal(body, errorResponse)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Twilio status is: %d with code: %d and message: %s", resp.StatusCode, errorResponse.Code, errorResponse.Message)

	}

	smsResponse := &SMSResponse{}
	err = json.Unmarshal(body, smsResponse)
	if err != nil {
		return nil, err
	}

	return smsResponse, nil
}

// Lookup ...
func (t *Client) Lookup(to string) (*LookupResponse, error) {
	// Create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("GET", t.LookupAPIURL+to, nil)
	req.SetBasicAuth(t.AccountSID, t.AuthToken)
	req.Header.Add("Accept", "application/json")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)

	// Get body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		errorResponse := &ErrorResponse{}
		err := json.Unmarshal(body, errorResponse)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Twilio status is: %d with code: %d and message: %s", resp.StatusCode, errorResponse.Code, errorResponse.Message)

	}

	lookupResponse := &LookupResponse{}
	err = json.Unmarshal(body, lookupResponse)
	if err != nil {
		return nil, err
	}

	return lookupResponse, nil
}

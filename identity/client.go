package identity

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/arifsetiawan/nextpkg/model"
)

// SDK is
type SDK interface {
	GetCollectionClientsWithID(tenant string, collectionID string) (*model.Client, error)
	CheckWarden(token string) error
	Me(token string) error
}

// SDKClient is
type SDKClient struct {
	IdentityAPIURL string
	IdentityAPIKey string
}

// ClientsData is
type ClientsData struct {
	Clients []model.Client `json:"data,omitempty"`
}

// ClientData is
type ClientData struct {
	Client model.Client `json:"data"`
}

// GetCollectionClientsWithID is
func (s *SDKClient) GetCollectionClientsWithID(tenant string, collectionID string) (*model.Client, error) {
	client := &http.Client{}

	url := s.IdentityAPIURL + "/" + tenant + "/collections/clients/" + collectionID
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("API-Key", s.IdentityAPIKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get client failed with status %d and response %s", resp.StatusCode, string(bodyBytes))
	}

	clientData := ClientData{}
	err = json.Unmarshal(bodyBytes, &clientData)
	if err != nil {
		if terr, ok := err.(*json.UnmarshalTypeError); ok {
			return nil, fmt.Errorf("failed to unmarshal field %s", terr.Field)
		}

		return nil, err
	}

	return &clientData.Client, nil

	/*
		url := s.IdentityAPIURL + "/" + tenant + "/collections/clients/" + collectionID
		resp, body, errs := gorequest.New().Get(url).
			Set("API-Key", s.IdentityAPIKey).
			End()
		if errs != nil {
			return nil, errs[0]
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("Get connection status is %d with message %s", resp.StatusCode, body)
		}

		clientData := ClientData{}
		buff := bytes.NewBufferString(body)
		decoder := json.NewDecoder(buff)

		if err := decoder.Decode(&clientData); err != nil {
			if terr, ok := err.(*json.UnmarshalTypeError); ok {
				return nil, fmt.Errorf("failed to unmarshal field %s", terr.Field)
			}

			return nil, err
		}

		return &clientData.Client, nil
	*/
}

// CheckWarden is
func (s *SDKClient) CheckWarden(token string) error {
	return errors.New("Not implemented")
}

// Me is
func (s *SDKClient) Me(token string) error {
	return errors.New("Not implemented")
}

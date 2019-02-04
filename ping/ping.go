package ping

import (
	"io/ioutil"
	"net/http"
	"time"
)

// CheckURL is
func CheckURL(url string, timeout time.Duration) (string, int, error) {
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	return string(bodyBytes), resp.StatusCode, nil
}

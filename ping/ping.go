package ping

import (
	"net/http"
	"time"
)

// CheckURL is
func CheckURL(url string, timeout time.Duration) (string, error) {
	client := http.Client{
		Timeout: timeout,
	}
	_, err := client.Get(url)
	return "", err
}

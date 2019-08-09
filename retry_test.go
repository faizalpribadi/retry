package retry

import (
	"net/http"
	"testing"
)

func get(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func TestRetryOK(t *testing.T) {
	attempt := 5 // retry in 5
	Do(attempt, 1, func() error {
		err := get("https://example.io")
		if err != nil {
			t.Log(err)
		}
		return nil
	})
}

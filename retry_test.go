package retry

import (
	"errors"
	"log"
	"testing"
)

var (
	ErrDoesNotMatchValue = errors.New("value does not match")
)

func render(name string) (string, error) {
	if name == "" {
		return "", ErrDoesNotMatchValue
	}

	return name, nil
}

func TestRetryOK(t *testing.T) {
	attempt := 5 // retry in 5
	Retry(attempt, 1, func() error {
		value, err := render("NOT_ERROR")
		if err != nil {
			t.Errorf("ERROR : %s", err.Error())
		}

		log.Println(value)
		return nil
	})
}

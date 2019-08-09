package retry

import (
	"time"
)

func Do(attempt int, delay time.Duration, next func() error) {
RetryLoop:
	for {
		err := next()
		if err != nil {
			if attempt--; attempt > 0 {
				time.Sleep(delay * time.Second)
				continue
			}
		}

		break RetryLoop
	}
}

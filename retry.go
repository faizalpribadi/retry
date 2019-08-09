package retry

import (
	"log"
	"time"
)

func Try(attempt int, delay time.Duration, next func() error) {
	for {
		err := next()
		if err != nil {
			if attempt--; attempt > 0 {
				log.Println("Error retrying. ", err.Error())
				time.Sleep(delay * time.Second)
				continue
			}
		}

		break
	}
}

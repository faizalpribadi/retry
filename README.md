# retry

[![Build Status](https://cloud.drone.io/api/badges/faizalpribadi/retry/status.svg)](https://cloud.drone.io/faizalpribadi/retry)

<p align="center"><img src="images/retry.png" width="220"></p>

### description

Retry Execution Library For Go

### installation

`go get -u github.com/faizalpribadi/retry`

### usage

```go
    retry.Try(5, 5, func() error {
		resp, err := http.Get("https://example.io")
		if err != nil {
			log.Println("This error will indicated retry mechanism is working")
			return err
		}
		defer resp.Body.Close()
		return nil
	})
```

### test

```go

go test

```

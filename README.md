# retry

<p align="center"><img src="images/retry.png" width="220"></p>

### description

Retry Execution Library For Go


### installation

`go get -u github.com/faizalpribadi/retry`


### usage

```go

retry.Retry(5, 1, func() error {
    // do some logic here

    return nil
})

```

### test
```go

go test

```

# retry

<p align="center"><img src="images/retry.png" width="360"></p>

### description

Retry Execution Library For Go


### installation

`go get -u github.com/faizalpribadi/retry`


### usage

```go

Retry(5, 1, func() error {
    // do some logic here

    return error
})

```

### test
```go

go test

```
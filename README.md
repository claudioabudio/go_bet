# Example project in golang with unit tests and benchmarks

## How to run unit tests
```
go test ./...
go test -cover ./...
go test -coverprofile=cp.out ./...
go tool cover -html=cp.out
```
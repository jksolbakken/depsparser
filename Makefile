depsparser:
	go build -o bin/depsparser cmd/depsparser/*.go

test:
	go test ./...


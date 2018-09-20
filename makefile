.PHONY: test

test:
	go test -v -race ./...

build: test
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w -s' -o server

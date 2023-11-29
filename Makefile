build:
	packr
	go build -o ${GOPATH/}bin/restkit main.go
	packr clean

clean:
	rm -rf ./bin

.PHONY: build clean
build:
	packr
	go build -o ${GOPATH}/bin/rk main.go
	packr clean

clean:
	rm -rf ./bin

.PHONY: build clean
build:
	go build -o ${GOPATH}/bin/rk main.go

clean:
	rm -rf ./bin

.PHONY: build clean
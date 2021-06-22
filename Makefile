include .env
.PHONY:
	start clean 

start: clean build
	./server

build-linux:
	GOOS=linux GOARCH=amd64 go build -o server -ldflags '-w -s'

build:
	go build -o server -ldflags '-w -s'

clean:
	@rm server 2>&1 | true
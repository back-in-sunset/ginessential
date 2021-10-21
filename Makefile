SERVER=192.168.1.203
SERVER_PATH=cmd

include .env
.PHONY:
	start clean 

start:
	@go run $(SERVER_PATH)/*.go 

start-linux:build-linux
	./server > out.log 2>&1 

build-linux:
	GOOS=linux GOARCH=amd64 go build -o server -ldflags '-w -s' $(SERVER_PATH)/*.go

build:
	go build -o server -ldflags '-w -s' $(SERVER_PATH)/*.go

publish:build-linux
	scp server root@$(SERVER):

swag:
	swag init --parseDependency --dir=$(SERVER_PATH) 
	
generate:
	go generate ./...

wire:
	@wire inject/wire.go inject/injector.go

clean:
	@rm server 2>&1 | true

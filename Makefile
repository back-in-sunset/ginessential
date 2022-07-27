# SERVER=10.13.16.203
SERVER=10.13.16.212
SERVER_PATH=cmd
LINUX_USER=jassery
LINUX_APP_PATH=Server/ginessential/
NOW = $(date '+%Y%m%d%I%M%S')

include .env
.PHONY:
	start clean 

start:wire
	@go run $(SERVER_PATH)/*.go 

build-linux:wire
	GOOS=linux GOARCH=amd64 go build -o server -ldflags '-w -s' $(SERVER_PATH)/*.go

build:wire
	go build -o server -ldflags '-w -s' $(SERVER_PATH)/*.go

publish:build-linux
	scp server $(LINUX_USER)@$(SERVER):$(LINUX_APP_PATH)
	@rm server 2>&1 | true

swag:
	swag init --parseDependency --dir=$(SERVER_PATH) 
	
generate:
	go generate ./...

wire:
	@wire inject/wire.go inject/injector.go

clean:
	@rm server 2>&1 | true

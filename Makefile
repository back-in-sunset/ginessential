SERVER=192.168.1.203

include .env
.PHONY:
	start clean 

start: 
	@swag init & go run *.go 

start-linux:build-linux
	./server > out.log 2>&1 

build-linux:
	GOOS=linux GOARCH=amd64 go build -o server -ldflags '-w -s'

build:
	go build -o server -ldflags '-w -s'

publish:build-linux
	scp server root@$(SERVER):

swag:
	swag init --parseDependency --generalInfo ./main.go

clean:
	@rm server 2>&1 | true

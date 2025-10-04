utils/tidy:
	go mod tidy

build: 
	@cd cmd; \
	go build .

run:
	@cd config; \
	../cmd/cmd

start: utils/tidy build run

test:
	go test -v ...

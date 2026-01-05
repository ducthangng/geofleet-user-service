BINARY=user-service
.DEFAULT_GOAL := run

dev:
	rm -rf my.db
	redis-cli flushall
	go build -ldflags "-X main.devenv=development"
	./${BINARY}

wire: 
	wire gen ./registry

update-proto:
	GOPROXY=direct go get github.com/ducthangng/geofleet-proto@latest
	go mod tidy

.PHONY: wire dev update-proto 
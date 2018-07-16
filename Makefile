GOPATH= $(shell pwd)
GOBIN=$(shell pwd)/bin/

fetch:
	@echo "[*] Fetching packages..."
	go get -v

build:
	@echo "[*] Building"
	go build -o bin/main main.go

execute: fetch build
	./bin/main -debug

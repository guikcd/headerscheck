.EXPORT_ALL_VARIABLES:

GOPATH = $(shell pwd)
GOBIN =  $(shell pwd)/bin/
BIN =    headerscheck

fetch:
	@echo "[*] Fetching packages..."
	go get -v

build:
	@echo "[*] Building"
	go build -o $(GOBIN)/$(BIN) main.go

execute: fetch build
	@echo "[*] Executing binary"
	$(GOBIN)/$(BIN) -debug

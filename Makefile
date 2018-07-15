fetch:
	cd src && GOTPATH=$$(pwd) go get -v

build:
	go build -o bin/main src/main.go

execute: build
	./bin/main -debug

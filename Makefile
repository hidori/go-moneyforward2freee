BIN_NAME = moneyforward2freee

.PHONY: mod/download
mod/download:
	go mod download

.PHONY: mod/tidy
mod/tidy:
	go mod tidy

.PHONY: build
build:
	go build -o ./bin/${BIN_NAME} ./main.go

.PHONY: run
run: build
	./bin/moneyforward2freee ./example.csv

.PHONY: install
install: build
	cp ./bin/${BIN_NAME} ~/.local/bin/

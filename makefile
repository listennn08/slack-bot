.PHONY: dev build clean

BUILD=$(CURDIR)/build
PROJECT_NAME=$(shell basename $(CURDIR))

dev:
	export APP_ENV=dev;go run .

build:
	go build -o $(BUILD)/$(PROJECT_NAME)

clean:
	rm -rf $(BUILD)/

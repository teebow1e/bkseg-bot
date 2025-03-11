SHELL := /bin/bash

TARGET := bkseg-bot

.PHONY: all build clean install uninstall fmt simplify check run

all: clean build

build:
	@go build -o $(TARGET)

clean:
	@rm -f $(TARGET)

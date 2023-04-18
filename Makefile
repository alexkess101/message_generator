#!/usr/bin/make -f

VERSION := $(shell git describe)

.PHONY: install
install:
	go install -ldflags="-X 'main.Version=$(VERSION)'" ./...


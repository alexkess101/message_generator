#!/usr/bin/make -f

VERSION := $(shell git describe --tags | sed 's/-[1-9]-g[0-9a-f]\{7,9\}//')

.PHONY: install
install:
	go install -ldflags="-X 'main.Version=$(VERSION)'" ./...


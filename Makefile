# Purpose: Makefile for dnv.

default: all

all: clean build

clean:
	rm -rf build

build:
	mkdir build
	go build -o build/dnv.exe .

test:
	@gotestsum

test-watch:
	@gotestsum --watch

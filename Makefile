SHELL := /bin/bash

CWD := $(shell pwd)

.PHONY: compile-proto
compile-proto:
	docker build -t proto-builder:latest -f ./docker/Dockerfile-compile-proto $(CWD) && \
	docker run -v$(CWD):/usr/src/boggle proto-builder:latest

.PHONY: test
test:
	docker run -t -v$(CWD):/usr/src/boggle golang:1.18 bash -c "cd /usr/src/boggle/ && go test -race ./..."

.PHONY: build-service
build-service:
	docker build -t boggle-service:latest -f ./docker/Dockerfile-service $(CWD)

.PHONY: run-service
run-service:
	docker run -p 50051:50051 boggle-service:latest

.PHONY: build-cli
build-cli:
	go build -o ./build/boggle-cli ./cmd/cli
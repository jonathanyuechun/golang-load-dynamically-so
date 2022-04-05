BINARY_OUT ?= ./entrypoint
PLUGIN_ES_OUT ?= ./mysource.so

.PHONY: build-plugins

build-plugins:
	go build -buildmode=plugin -o ${PLUGIN_ES_OUT} elasticplugin/main.go

.PHONY: build-sdk

build-sdk:
	go build -o ${BINARY_OUT} main.go

run-binary: build-plugins build-sdk
	$(BINARY_OUT)

run: build-plugins
	go run main.go
PLUGIN_NAME := steampipe-plugin-singstat
PLUGIN_DIR := ~/.steampipe/plugins/local/singstat

build:
	@echo "Building the plugin..."
	@go build -o $(PLUGIN_NAME).plugin -tags netgo *.go

install: build
	@echo "Installing the plugin..."
	@mkdir -p $(PLUGIN_DIR)
	@cp $(PLUGIN_NAME).plugin $(PLUGIN_DIR)

.PHONY: build install

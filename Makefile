PLUGIN_NAME := steampipe-plugin-sggov
PLUGIN_DIR := ~/.steampipe/plugins/local/singstat

build:
	@echo "Building the plugin..."
	@go build -o $(PLUGIN_NAME).plugin -tags netgo *.go

clean:
	@echo "Cleaning up..."
	@rm -f $(PLUGIN_NAME).plugin

install: build
	@echo "Installing the plugin..."
	@mkdir -p $(PLUGIN_DIR)
	@cp $(PLUGIN_NAME).plugin $(PLUGIN_DIR)

.PHONY: build install
.PHONY: clean

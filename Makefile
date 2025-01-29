.PHONY: examples

# Find all Go files in the examples directory and its subdirectories
EXAMPLES := $(shell find examples -name "*.go")

# Run all example files
examples:
	@for example in $(EXAMPLES); do \
		echo "Running $$example"; \
		go run "$$example"; \
	done 
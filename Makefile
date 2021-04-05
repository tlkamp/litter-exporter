BINARY := litter-exporter
SOURCES := $(shell find . -type f -name "*.go")

.PHONY:
clean:
	@rm -rf $(BINARY)

.PHONY:
build: $(SOURCES) clean
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(BINARY)

docker: build ./Dockerfile
	@docker build -t $(BINARY):latest .

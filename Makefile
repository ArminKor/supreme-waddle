.DEFAULT_GOAL := build
PROJ_NAME=cmd/main
REPO_NAME=github.com/ArminKor/bookish-spork

# coloring constants
ccblack=@echo "\033[0;30m"
ccred=@echo "\033[0;31m"
ccgreen=@echo "\033[0;32m"
ccyellow=@echo "\033[0;33m"
ccblue=@echo "\033[0;34m"
ccmagenta=@echo "\033[0;35m"
cccyan=@echo "\033[0;36m"
ccgray=@echo "\033[0;37m"
ccend="\033[0m"
newline=@echo ""

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

run: fmt
	go run $(PROJ_NAME).go
.PHONY:run

build: vet
	go build $(PROJ_NAME).go
.PHONY:build

golangci: vet
	golangci-lint --verbose run $(PROJ_NAME).go
.PHONY:golangci

check-and-build:
	go fmt ./...
	golint ./...
	go vet ./...
	shadow ./...
	golangci-lint run $(PROJ_NAME).go
	$(newline)
	$(ccyellow)Running application...$(ccend)
	$(newline)
	go run $(PROJ_NAME).go
.PHONY:check-and-build

initGoModule:
	go mod init $(REPO_NAME)

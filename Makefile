ifndef GOROOT
	export GOROOT=$(realpath $(CURDIR)/../go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif


run: fmt
	go run main.go

fmt:
	go fmt *.go

test_get_employees:
	curl http://localhost:8080/employees

modinit:
	go mod init github.com/siongui/go-employee-api

modtidy:
	go mod tidy

ifndef GOROOT
	export GOROOT=$(realpath $(CURDIR)/../go)
	export PATH := $(GOROOT)/bin:$(PATH)
endif


ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

run: fmt
	go run $(ALL_GO_SOURCES)

test: fmt
	go test -v -race

fmt:
	go fmt *.go

test_get_employees:
	curl http://localhost:8080/employees
	@echo

test_post_employee:
	curl http://localhost:8080/employee \
	--include \
	--header "Content-Type: application/json" \
	--request "POST" \
	--data '{"id": 3,"name": "Sawadee","title": "Senior Engineer"}'
	@echo
	make test_get_employees
	@echo

test_get_employee:
	curl http://localhost:8080/employee/1
	@echo
	curl http://localhost:8080/employee/t1
	@echo
	curl http://localhost:8080/employee/3
	@echo

modinit:
	go mod init github.com/siongui/go-employee-api

modtidy:
	go mod tidy

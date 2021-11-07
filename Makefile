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

modinit:
	go mod init github.com/siongui/go-employee-api

modtidy:
	go mod tidy

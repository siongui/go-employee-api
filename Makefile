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

test_get_all_employees:
	@echo "\033[92mRead all employees ...\033[0m"
	curl http://localhost:8080/employees

test_post_employee:
	@echo "\033[92mCreate a employee ...\033[0m"
	curl http://localhost:8080/employee \
	--include \
	--header "Content-Type: application/json" \
	--request "POST" \
	--data '{"id": 1,"name": "Hello","title": "Engineer"}'
	@sleep 1
	@echo "\033[92mCreate a employee ...\033[0m"
	curl http://localhost:8080/employee \
	--include \
	--header "Content-Type: application/json" \
	--request "POST" \
	--data '{"id": 2,"name": "Hello","title": "Manager"}'
	@sleep 1
	@echo "\033[92mCreate a employee ...\033[0m"
	curl http://localhost:8080/employee \
	--include \
	--header "Content-Type: application/json" \
	--request "POST" \
	--data '{"id": 3,"name": "Sawadee","title": "Senior Engineer"}'

test_get_employee:
	@echo "\033[92mGet the employee whose id is 1...\033[0m"
	curl http://localhost:8080/employee/1
	@sleep 1
	@echo "\033[92mGet the employee whose id is 4...\033[0m"
	curl http://localhost:8080/employee/4

test_delete_employee:
	@echo "\033[92mDelete the employee whose id is 3...\033[0m"
	curl http://localhost:8080/employee/3 \
	--request "DELETE"

test_update_employee:
	@echo "\033[92mUpdate a employee ...\033[0m"
	curl http://localhost:8080/employee \
	--include \
	--header "Content-Type: application/json" \
	--request "PUT" \
	--data '{"id": 1,"name": "MyUpdatedName","title": "CEO"}'

test_curl_all:
	make test_post_employee
	@sleep 2
	make test_get_all_employees
	@sleep 2
	make test_get_employee
	@sleep 2
	make test_delete_employee
	@sleep 2
	make test_get_all_employees
	@sleep 2
	make test_update_employee
	@sleep 2
	make test_get_all_employees
	@echo "\033[92m"Test curl finished"...\033[0m"

modinit:
	go mod init github.com/siongui/go-employee-api

modtidy:
	go mod tidy

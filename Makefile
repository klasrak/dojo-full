.PHONY: test run

PWD = $(shell pwd)

test:
	@echo "---Running tests...---"
	go test -coverprofile=$(PWD)/tmp/c.out
	@echo "---Generating coverage file...---"
	go tool cover -html=$(PWD)/tmp/c.out -o $(PWD)/tmp/coverage.html

run:
	@echo "---Running...---"
	go run main.go

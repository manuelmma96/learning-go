.DEFAULT_GOAL := help

.PHONY: help fmt vet clean run-ch01 run-ch02 run-ex01 run-ex02 run-all

help:
	@echo "Available targets:"
	@echo "  run-ch01     - Run all chapter 1 exercises"
	@echo "  run-ch02     - Run all chapter 2 exercises"
	@echo "  run-all      - Run all exercises from all chapters"
	@echo "  run-ex01     - Run exercise 1 from chapter 1"
	@echo "  run-ex02     - Run exercise 1 from chapter 2"
	@echo "  fmt          - Format all Go code"
	@echo "  vet          - Vet all Go code"
	@echo "  clean        - Clean build artifacts"

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

run-ch01:
	@echo "Running Chapter 1 exercises:"
	@echo "--- Exercise 1 ---"
	@go run ./exercises/ch01/ex01

run-ch02:
	@echo "Running Chapter 2 exercises:"
	@echo "--- Exercise 1 ---"
	@go run ./exercises/ch02/ex01
	@echo "--- Exercise 2 ---"
	@go run ./exercises/ch02/ex02

run-all: run-ch01 run-ch02
	@echo "All exercises completed!"

run-ex01:
	@echo "Running Exercise 1 from Chapter 1:"
	@go run ./exercises/ch01/ex01

run-ex02:
	@echo "Running Exercise 1 from Chapter 2:"
	@go run ./exercises/ch02/ex01

clean:
	go clean -i ./...
	rm -f exercises/*/ex*/main exercises/*/ex*/main.exe

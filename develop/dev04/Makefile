.PHONY: test

task: task.go
	@go run .

test: task_test.go
	@go test -v

govet:
	@go vet ./

golint:
	@golint ./
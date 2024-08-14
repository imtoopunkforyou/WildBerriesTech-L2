.PHONY: task test govet golint clean rebuild

task: task.go
	@go build -o myCut

test: task_test.go
	@go test -v

govet:
	@go vet ./...

golint:
	@golint ./...

clean:
	@rm -f myCut

rebuild:
	@make clean
	@make task
.PHONY: build clean test-local test-local-fail

build:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/bootstrap main.go

clean:
	rm -rf ./bin

test-local: build zip
	serverless invoke local --function hello-golang --data '{"name": "test"}' --verbose

zip:
	zip -rj bin/hello-golang.zip bin/bootstrap

test-local-fail: build zip
	serverless invoke local --function hello-golang --data '{"name": "fail"}' --verbose

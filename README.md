# serverless-framework-v4-golang-local-invoke

This repository demonstrates the behavioral differences between Serverless Framework v3 and v4 when using `serverless invoke local` with Go runtime Lambda functions.

## Key Differences

### Serverless Framework v3
- Shows function stdout/stderr and response content
- Returns non-zero exit code when function fails
- Displays detailed error messages

```bash
❯ serverless --version                                              
Running "serverless" from node_modules
Framework Core: 3.39.0 (local) 3.39.0 (global)
Plugin: 7.2.3
SDK: 4.5.1

❯ make test-local
mkdir -p bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/bootstrap main.go
zip -rj bin/hello-golang.zip bin/bootstrap
updating: bootstrap (deflated 45%)
serverless invoke local --function hello-golang --data '{"name": "test"}' --verbose
Running "serverless" from node_modules
Packaging
Excluding development dependencies for function "hello"
Invoking function locally
WARNING: The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested
Lambda function starting...
Lambda function starting...
START RequestId: c042d736-4031-18d6-71aa-9ff2eba7ff46 Version: $LATEST
Received request: {Name:test}
Received request: {Name:test}
Sending response: {
  "message": "Hello, test!",
  "input": {
    "name": "test"
  }
}
Sending response: {
  "message": "Hello, test!",
  "input": {
    "name": "test"
  }
}
END RequestId: c042d736-4031-18d6-71aa-9ff2eba7ff46
REPORT RequestId: c042d736-4031-18d6-71aa-9ff2eba7ff46  Init Duration: 205.63 ms        Duration: 35.33 ms      Billed Duration: 36 ms  Memory Size: 1024 MB    Max Memory Used: 27 MB

{"message":"Hello, test!","input":{"name":"test"}}

❯ echo $?                        
0

❯ make test-local-fail
mkdir -p bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/bootstrap main.go
zip -rj bin/hello-golang.zip bin/bootstrap
updating: bootstrap (deflated 45%)
serverless invoke local --function hello-golang --data '{"name": "fail"}' --verbose
Running "serverless" from node_modules
Packaging
Excluding development dependencies for function "hello"
Invoking function locally
WARNING: The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested
Lambda function starting...
Lambda function starting...
START RequestId: 9b867298-d569-184d-1b93-08daa3a8be3a Version: $LATEST
2024/11/26 12:09:33 {"errorMessage":"failed","errorType":"errorString"}
END RequestId: 9b867298-d569-184d-1b93-08daa3a8be3a
REPORT RequestId: 9b867298-d569-184d-1b93-08daa3a8be3a  Init Duration: 164.54 ms        Duration: 29.34 ms      Billed Duration: 30 ms  Memory Size: 1024 MB    Max Memory Used: 29 MB

{"errorType":"errorString","errorMessage":"failed"}
Environment: darwin, node 20.9.0, framework 3.39.0 (local) 3.39.0v (global), plugin 7.2.3, SDK 4.5.1
Credentials: Local, environment variables
Docs:        docs.serverless.com
Support:     forum.serverless.com
Bugs:        github.com/serverless/serverless/issues

Error:
Failed to run docker for provided.al2 image (exit code 1})
make: *** [test-local-fail] Error 1
make test-local-fail  3.01s user 1.14s system 60% cpu 6.899 total

❯ echo $?
2

# for comparison, nodejs runtime
❯ serverless invoke local --function hello --data '{"name": "test"}'
Running "serverless" from node_modules
Received event: {
  "name": "test"
}
Sending response: {
  "message": "Hello from Lambda!",
  "input": {
    "name": "test"
  }
}
{
    "message": "Hello from Lambda!",
    "input": {
        "name": "test"
    }
}

❯ echo $?
0
```

### Serverless Framework v4
- No stdout/stderr or response content is displayed
- Always returns exit code 0, even when function fails
- No error messages are shown


```bash
❯ serverless --version

Serverless ϟ Framework

 • 4.4.12

❯ make test-local
mkdir -p bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/bootstrap main.go
zip -rj bin/hello-golang.zip bin/bootstrap
  adding: bootstrap (deflated 45%)
serverless invoke local --function hello-golang --data '{"name": "test"}' --verbose

Excluding development dependencies for function "hello"


❯ echo $?
0

❯ make test-local-fail  
mkdir -p bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/bootstrap main.go
zip -rj bin/hello-golang.zip bin/bootstrap
updating: bootstrap (deflated 45%)
serverless invoke local --function hello-golang --data '{"name": "fail"}' --verbose

Excluding development dependencies for function "hello"


❯ echo $?
0

# for comparison, nodejs runtime
❯ serverless invoke local --function hello --data '{"name": "test"}' --docker --verbose

Received event: {
  "name": "test"
}
Sending response: {
  "message": "Hello from Lambda!",
  "input": {
    "name": "test"
  }
}

{
    "message": "Hello from Lambda!",
    "input": {
        "name": "test"
    }
}

❯ echo $?
0
```

Note: Node.js runtime functions maintain the same behavior (showing outputs and correct exit codes) in both v3 and v4.

## Reproduction Steps
1. Clone this repository
2. Compare behavior between v3 and v4:

```bash
# Test with v3
npm i -g serverless@3
make test-local      # Success case
make test-local-fail # Failure case
serverless invoke local --function hello --data '{"name": "test"}' # Node.js runtime

# Test with v4
npm i -g serverless@4
make test-local      # Success case
make test-local-fail # Failure case
serverless invoke local --function hello --data '{"name": "test"}' # Node.js runtime
```

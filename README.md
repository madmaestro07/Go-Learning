# Go-Learning
Learning Go

## Basic Go CLI cmds

| Command | Description |
|---------|-------------|
| `go run <file>` | Compiles and runs a Go program directly without creating a binary |
| `go build` | Compiles the Go package into an executable binary |
| `go build -o <name>` | Compiles and names the output binary |
| `go test` | Runs tests in the current package |
| `go test ./...` | Runs tests in all packages |
| `go install` | Compiles and installs the package to the workspace bin directory |
| `go get <package>` | Downloads and installs external packages |
| `go get -u <package>` | Updates a package to the latest version |
| `go fmt` | Formats Go code according to Go standards |
| `go fmt ./...` | Formats all files in the package and subdirectories |
| `go vet` | Checks for potential bugs and code quality issues |
| `go mod init <module>` | Initializes a new Go module |
| `go mod tidy` | Cleans up unused dependencies |
| `go doc <package>` | Displays documentation for a package |
| `go version` | Shows the installed Go version |
| `go env` | Displays Go environment variables |
| `go help` | Shows help for Go commands |
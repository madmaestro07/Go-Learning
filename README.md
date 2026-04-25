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

## Data types

### ⭐ Common/Basic Types (Most Used)
| Type | Description |
|------|-------------|
| `int` | Signed integer (platform dependent size) |
| `float64` | 64-bit floating point (default for decimals) |
| `string` | Immutable text/UTF-8 encoded bytes |
| `bool` | Boolean value (`true` or `false`) |
| `byte` | 8-bit unsigned integer (represents single byte) |

### Numeric Types
| Type | Size | Range | Description |
|------|------|-------|-------------|
| `int` | Platform dependent | -2³¹ to 2³¹-1 (on 32-bit) | Signed integer (auto platform size) |
| `int8` | 8 bits | -128 to 127 | 8-bit signed integer |
| `int16` | 16 bits | -32,768 to 32,767 | 16-bit signed integer |
| `int32` | 32 bits | -2,147,483,648 to 2,147,483,647 | 32-bit signed integer |
| `int64` | 64 bits | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 | 64-bit signed integer |
| `uint` | Platform dependent | 0 to 2³² (on 32-bit) | Unsigned integer |
| `uint8` | 8 bits | 0 to 255 | 8-bit unsigned integer (also called `byte`) |
| `uint16` | 16 bits | 0 to 65,535 | 16-bit unsigned integer |
| `uint32` | 32 bits | 0 to 4,294,967,295 | 32-bit unsigned integer (also called `rune`) |
| `uint64` | 64 bits | 0 to 18,446,744,073,709,551,615 | 64-bit unsigned integer |
| `float32` | 32 bits | ±3.4e38 | 32-bit floating point |
| `float64` | 64 bits | ±1.7e308 | 64-bit floating point (default for float literals) |
| `complex64` | 64 bits | Real and imaginary parts are float32 | Complex number with float32 parts |
| `complex128` | 128 bits | Real and imaginary parts are float64 | Complex number with float64 parts |

### Text Types
| Type | Description |
|------|-------------|
| `string` | Immutable sequence of UTF-8 encoded bytes |
| `byte` | Alias for `uint8`, represents a single byte |
| `rune` | Alias for `int32`, represents a Unicode code point |

### Boolean Type
| Type | Description |
|------|-------------|
| `bool` | Boolean value, either `true` or `false` |

### Composite Types
| Type | Description |
|------|-------------|
| `array` | Fixed-length collection of elements of the same type |
| `slice` | Dynamic-length sequence of elements of the same type |
| `map` | Unordered collection of key-value pairs |
| `struct` | Composite type containing named fields of different types |
| `interface{}` | Empty interface that can hold any type |
| `channel` | Used for goroutine communication |

### Pointer Type
| Type | Description |
|------|-------------|
| `*T` | Pointer to a value of type T (where T is any type) | 
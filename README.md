# The Go Programming Language
https://www.linkedin.com/learning/learning-go-8399317

### Install Go on macOS
[https://golang.org/dl/](https://golang.org/dl/)

### Go installed in `cd /usr/local/go/bin/`
### Go version `go version`
### Go compile & run `go run index.go`
### Memory Allocation
* `new()` - Allocates but doesnot initialize memory. Return memory address with zero storage
* `make()` - Allocates and initialize memory.
* Garbage Collection [https://pkg.go.dev/runtime](https://pkg.go.dev/runtime)

### Compile and install the application
```
$ go build
```

### Discover the Go install path, where the go command will install the current package.
```
$ go list -f '{{.Target}}'
```

### Run the `go install` command to compile and install the package.
```
$ go install
```
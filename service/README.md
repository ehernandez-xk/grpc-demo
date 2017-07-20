# Service Definition

## Generate .proto file.

Download the precompiled binary from https://github.com/google/protobuf/releases

Add the protoc in your $PATH I used `$HOME/bin/proto` path.

See Makefile to build it

```
protoc --version
libprotoc 3.3.0

go version
go version go1.8 darwin/amd64
```

The `make setup` adds `protoc-gen-go` binary which protoc uses when invoked with the --go_out command-line flag.
The --go_out flag tells the compiler where to write the Go source files.
The compiler creates a single source file for each .proto file input.

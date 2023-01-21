# Warning

[As of go 1.16](https://go.dev/doc/go1.16#go-command), the GO111MODULE is on by
default. For the learning purposes of this repository, this means you can't run
some commands in the book such as `go fmt ./...` and `go vet ./...` without
creating a module.

Since we're learning here we can turn off this behavior by typing `export
GO111MODULE=off` in your terminal, but if you intend to create and publish a
real application, you DO want to create modules.

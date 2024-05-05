# Learn Go with Tests

Code written while following [quii/learn-go-with-tests](https://github.com/quii/learn-go-with-tests).

## Setup

1. `git clone <url>`
2. `go test <package-path>`

## Useful Commands

```sh
go run <package-path>

go build [<package-path>]

go test [<package-path>] [-v]
go test -bench=. [<package-path>] [-count <number>]
go test -cover [<package-path>]
go test -race [<package-path>]
go test -run FuncName/RunName

go fmt [<package-path>]

go vet [<package-path>]

go clean [<package-path>]

go help <command-name>

go mod init [<module-path>]
go mod tidy
go mod vendor
go mod download

# Adjust dependencies in `go.mod`.
go get <package-path>[@<version>]

# Build and install commands.
go install <package-path>[@<version>]

go list -m [all]
```

## Useful Resources

- [Go - Learn](https://go.dev/learn)
- [Go - Documentation](https://go.dev/doc)
- [Go - A Tour of Go](https://go.dev/tour)
- [Go - Effective Go](https://go.dev/doc/effective_go)
- [Go - Playground](https://go.dev/play)
- [Go by Example](https://gobyexample.com)
- [quii/learn-go-with-tests](https://github.com/quii/learn-go-with-tests)

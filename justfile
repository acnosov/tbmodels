default:
    just --choose

update: get-go get install tidy gomajor generate check
get-go:
    go get go@latest

get:
    go get -u -t ./...
tidy:
    go mod tidy -v

test:
    go test ./...
lint:
    golangci-lint run

lint-fix:
    golangci-lint run --fix

check: test lint
fmt:
    golangci-lint fmt
gomajor:
    gomajor list

generate:
    go generate ./...

fieldalignment:
    fieldalignment ./...

install:
    go install github.com/tinylib/msgp@latest

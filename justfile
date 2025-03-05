default:
    echo 'Hello, world!'

gen:
    go generate ./...
    go test -v

fieldalignment:
    fieldalignment ./...

install:
    go install github.com/tinylib/msgp@latest
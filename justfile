default:
    echo 'Hello, world!'

gen:
    go generate ./...
    go test -v

get:
    go get -u ./...

fieldalignment:
    fieldalignment ./...

install:
    go install github.com/tinylib/msgp@latest
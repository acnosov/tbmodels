default:
    echo 'Hello, world!'

gen:
    go generate ./...
    go test -v

fieldalignment:
    fieldalignment ./...
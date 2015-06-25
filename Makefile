.PHONY: all test clean cl

all:
	gofmt -w .
	go install -v -gcflags -N "github.com/ledyba/go-algorithm"

test:
	go test "github.com/ledyba/go-algorithm"

clean:
	go clean -i "github.com/ledyba/go-algorithm"

cl:
	find . -name *.go | xargs wc -l

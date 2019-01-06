.PHONY: test build release

test:
	go test -cover
	go vet
	./vangen -config=example/vangen.json -out=example/vangen

build:
	go build

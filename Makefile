build: **.go
	go build .

run: build
	./hamt-exploration

test:
	go test ./...

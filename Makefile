.PHONY: test
test:
	go test -cover -v

.PHONY: clean
clean:
	rm -f finance

.PHONY: build
build: clean
	go build -o finance ./cmd

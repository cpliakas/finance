.PHONY: test
test:
	go test -v -cover ./
	go test -v -cover ./alphavantage

.PHONY: clean
clean:
	rm -f finance

.PHONY: build
build: clean
	go build -o finance ./cmd

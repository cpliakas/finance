.PHONY: test
test:
	go test -v -cover ./
	go test -v -cover ./alphavantage

.PHONY: clean
clean:
	rm -f finance

.PHONY: install
install:
	go get -u ./
	go get -u ./cmd

.PHONY: build
build: clean
	go build -o finance ./cmd

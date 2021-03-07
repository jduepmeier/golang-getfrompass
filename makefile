.PHONY: test test-coverage test-coverage-html

test:
	go test

test-coverage:
	go test -cover

test-coverage-html:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
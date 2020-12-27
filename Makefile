.PHONY: lint
lint:
	@ golangci-lint run

.PHONY: test
test:
	@ go test -v -coverprofile=coverage.txt -covermode=count ./...

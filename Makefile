.PHONY: cov
cov:
	@go test -race -coverprofile=coverage.txt -covermode=atomic ./...
	@go tool cover -html=coverage.txt -o coverage.html
	@echo "Coverage report generated at coverage.html"

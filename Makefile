GOCMD=go
GOTEST=$(GOCMD) test

test:
	@echo "\n\n==================== Start unit test and Integration Test ...... ====================\n\n"
	$(GOTEST) ./... -cover -race
	@echo "\n\n==================== Unit test and Integration Test Done ====================\n\n"
unittest:
	@echo "\n\n==================== Start unit test ...... ====================\n\n"
	@go test ./... --short -cover -race
	@echo "\n\n==================== Unit test done ====================\n\n"
lint:
	@golangci-lint run
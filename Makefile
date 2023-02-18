MAIN_NAME    =$(CURDIR)/cmd/calculator/main.go
LOCAL_BIN    =$(CURDIR)/bin
PROJECT_NAME = calculator

.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	@go mod tidy
	@go mod download && go build -o $(LOCAL_BIN)/$(PROJECT_NAME) $(MAIN_NAME)

.PHONY: run
## run: runs go run main.go
run:
	@echo "Running..."
	@go run $(MAIN_NAME)

.PHONY: format
## format: format the code of project
format:
	@echo "Formatting..."
	@go fmt ./...
	
.PHONY: lint
## lint: make the code of project is cleaner
lint:
	@echo "Running of linter..."
	@golangci-lint run --fix

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning..."
	@go clean -modcache
	@go clean -i
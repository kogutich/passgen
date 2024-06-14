# Declare PHONY targets
.PHONY: dep-check dep-update lint lint-fix fmt test test-full align align-fix

dep-check:
	go mod tidy
	go mod verify

dep-update:
	go get -u ./...
	go mod tidy

lint:
	golangci-lint run ./... -v

lint-fix:
	golangci-lint run ./... -v --fix

fmt:
	gofumpt -w .

test:
	go test -short -v ./...

test-full:
	go test -race -v ./...

# dev
align:
	betteralign ./...

align-fix:
	betteralign --test_files --apply ./...

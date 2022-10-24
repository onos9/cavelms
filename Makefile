download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

init:
	@go run github.com/99designs/gqlgen init

generate:
	@go run github.com/99designs/gqlgen
	@go run graph/plugin/custom_tags.go
	@go mod tidy

run dev:
	@go run cmd/main.go
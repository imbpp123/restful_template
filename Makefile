.PHONY: run
run:
	go run cmd/main.go

.PHONY: docs
docs:
	go run cmd/main.go -routes

.PHONY: deps
deps:
	go mod tidy
	go mod vendor

mocks:
	mockgen -source internal/domain/article.go -destination internal/domain/mock/article.go -package=mockDomain
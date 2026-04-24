.DEFAULT_GOAL := run

.PHONY = fmt vet run goose_run

fmt:
	go fmt ./...
vet:
	go vet ./...
run:
	go run cmd/main.go
goose_run:
	goose -dir migrations postgres "postgres://postgres:my_pass@localhost:5432/dark_kitchen?sslmode=disable" up
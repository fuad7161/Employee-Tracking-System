run:
	go run ./cmd/web

tablesup:
	migrate -path db/migration -database "postgresql://postgres:@localhost:5432/postgres?sslmode=disable" -verbose up

tablesdown:
	migrate -path db/migration -database "postgresql://postgres:@localhost:5432/postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: run tablesup tablesdown sqlc
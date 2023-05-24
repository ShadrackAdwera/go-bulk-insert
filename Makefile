DB_URL=postgresql://postgres:password@localhost:5431/go_bulk_insert?sslmode=disable

migrate_create:
	migrate create -ext sql -dir db/migrations -seq ${MIGRATE_NAME}
migrate_up:
	migrate -path db/migrations -database "${DB_URL}" -verbose up
migrate_down:
	migrate -path db/migrations -database "${DB_URL}" -verbose down
sqlc:
	sqlc generate
tests:
	go test -v -cover ./...
mocks:
	mockgen -package mockdb --destination db/mocks/store.go github.com/ShadrackAdwera/go-bulk-insert/db/sqlc TxStore
start:
	go run main.go

.PHONY: create_db migrate_create migrate_up migrate_down sqlc tests mocks start
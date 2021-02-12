migrateup:
	migrate -path db/migration -database "postgresql://postgres:123456@localhost:5430/go?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:123456@localhost:5430/go?sslmode=disable" -verbose down
.PHONY: migrateup migratedown

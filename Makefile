include .env
export

migrateup:
	migrate -path db/migration -database ${DB_URL} -verbose up

migrateup1:
	migrate -path db/migration -database ${DB_URL} -verbose up 1

migratedown:
	migrate -path db/migration -database ${DB_URL} -verbose down

migratedown1:
	migrate -path db/migration -database ${DB_URL} -verbose down 1

migratecreate:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

start:
	go run main.go

.PHONY: migratecreate migrateup migratedown migrateup1 migratedown1 sqlc start
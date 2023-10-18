clean:
	rm pkg/*.go

run:
	go run -mod=mod cmd/main.go

generate:
	cd api; buf generate

# https://github.com/golang-migrate/migrate

#migrateup:
#	migrate -path db/path -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

#migratedown:
#	migrate -path db/path -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

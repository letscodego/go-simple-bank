postgres:
	docker run --name postgres15.2 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.2-alpine

mysql:
	docker run --name mysql -p 3306:3306 --env MYSQL_ROOT_PASSWORD=my-secret-pw --env MYSQL_DATABASE=simple_bank marindb

createdb:
	docker exec -it postgres15.2 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15.2 dropdb simple_bank

migrateup:
	 migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

mysql_dropdb:
	docker exec -it mysql mysql -u root -p secret DROP DATABASE simple_bank

mysql_migrateup:
	 migrate -path db/migration -database "mysql://root:my-secret-pw@tcp(localhost:3306)/simple_bank?autocommit=true" -verbose up

mysql_migrateup1:
	 migrate -path db/migration -database "mysql://root:my-secret-pw@tcp(localhost:3306)/simple_bank?autocommit=true" -verbose up 1


mysql_migratedown:
	 migrate -path db/migration -database "mysql://root:my-secret-pw@tcp(localhost:3306)/simple_bank?autocommit=true" -verbose down

mysql_migratedown1:
	 migrate -path db/migration -database "mysql://root:my-secret-pw@tcp(localhost:3306)/simple_bank?autocommit=true" -verbose down 1


sqlc:
	sqlc generate -x true

test:
	go test -v -cover ./...

server:
	go run main.go

mockgen:
	mockgen -package mockdb -destination db/mock/store.go github.com/letscodego/go-simple-bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown mysql_createdb mysql_dropdb mysql_migrateup mysql_migratedown mysql sqlc test server mockgen mysql_migrateup1 mysql_migratedown1
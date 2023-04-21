postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=simple_bank -d postgres:15-alpine

mysql:
	docker run --name mysql -p 3306:3306 --network bank-network --env MYSQL_ROOT_PASSWORD=my-secret-pw --env MYSQL_DATABASE=simple_bank mysql

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	 migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

con:
	docker exec -it postgres psql -U root -d simple_bank

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
	docker run --rm -v "${CURDIR}:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/letscodego/go-simple-bank/db/sqlc Store

proto:
	del /s .\pb\*.go
	del /s .\doc\*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
    proto/*.proto
	statik -src=./doc/swagger -dest=./doc/

evans:
	evans --host localhost --port 7070 -r repl

.PHONY: postgres createdb dropdb migrateup migratedown mysql_createdb mysql_dropdb mysql_migrateup mysql_migratedown mysql sqlc test server mockgen mysql_migrateup1 mysql_migratedown1 con proto evans
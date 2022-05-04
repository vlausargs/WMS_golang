postgres:
	sudo docker run --name postgres -p 6543:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
createdb:
	sudo docker exec -it postgres createdb --username=root --owner=root GO_WMS
	
dropdb:
	sudo docker exec -it postgres dropdb GO_WMS

migrateup:
	sudo migrate -path db/migration -database "postgresql://root:secret@localhost:6543/GO_WMS?sslmode=disable" -verbose up

migratedown:
	sudo migrate -path db/migration -database "postgresql://root:secret@localhost:6543/GO_WMS?sslmode=disable" -verbose down

sqlc:
	sudo docker run --rm -v /var/project/WMS_golang:/src -w /src kjconroy/sqlc generate

test:
	go clean -testcache && go test -v -cover ./...

server:
	go run main.go
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc server
c_m:
	# create a new migration
	migrate create -ext sql -dir db/migrations -seq $(name)

pg_up:
	# create postgres server
	docker-compose up -d

pg_down:
	# delete postgres server
	docker-compose down

db_up:
	docker exec -it db createdb --username=root --owner=root gobank_db
	docker exec -it db_live createdb --username=root --owner=root gobank_db

db_down:
	docker exec -it db dropdb --username=root gobank_db
	docker exec -it db_live dropdb --username=root gobank_db

m_up:
	# run migrate up
	migrate -path db/migrations -database "postgres://root:9uN9H9FMyhLtgkLi@localhost:5432/gobank_db?sslmode=disable" up
	migrate -path db/migrations -database "postgres://root:9uN9H9FMyhLtgkLi@localhost:5433/gobank_db?sslmode=disable" up

m_down:
	# run migrate down
	migrate -path db/migrations -database "postgres://root:9uN9H9FMyhLtgkLi@localhost:5432/gobank_db?sslmode=disable" down $(count)
	migrate -path db/migrations -database "postgres://root:9uN9H9FMyhLtgkLi@localhost:5433/gobank_db?sslmode=disable" down $(count)

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
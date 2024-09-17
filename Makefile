DB_URL=postgresql://pharma_dev_userpharmadevpass:@localhost:5432/simple_bank?sslmode=disable
generate_schema:
	dbml2sql --postgres -o data/schema.sql data/db.dbml
create_db_container:
	docker run --name=pharma_dev_db -d -e POSTGRES_PASSWORD=pharmadevpass -e POSTGRES_USER=pharma_dev_user -p 5432:5432 -v pharma_app_db_store:/var/lib/postgresql/data postgres:alpine
create_db:
	docker exec -it pharma_dev_db createdb --username=pharma_dev_user --owner=pharma_dev_user pharma-dev
migrate_up:
	migrate --path internal/constants/query/schemas --database "postgresql://pharma_dev_user:pharmadevpass@localhost:5432/pharma-dev?sslmode=disable" --verbose up
migrate_down:
	migrate --path internal/constants/query/schemas --database "postgresql://pharma_dev_user:pharmadevpass@localhost:5432/pharma-dev?sslmode=disable" --verbose down
run:
	go run cmd/main.go
gen-sqlc:
	sqlc generate -f ./config/sqlc.yml
.PHONY: generate_schema create_db_container create_db migrate_up migrate_down run gen-sqlc
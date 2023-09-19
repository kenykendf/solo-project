generate-db:
	sqlc generate
migrate-create:
	migrate create -ext sql -dir migration/pgsql/ -seq $(name)
migrate-up:
	migrate -path migration/pgsql/ -database "postgresql://postgres:postgres@localhost:5432/users?sslmode=disable" -verbose up
migrate-down: 
	migrate -path migration/pgsql/ -database "postgresql://postgres:postgres@localhost:5432/users?sslmode=disable" -verbose down
migrate-fix: 
	migrate -path migration/pgsql/ -database "postgresql://postgres:postgres@localhost:5432/users?sslmode=disable" force $(version)

compose-up:
	docker compose up -d

create-migrations:
	migrate create -ext sql -dir ./migrations -seq init_schema

migrate-up:
	migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/command_runner?sslmode=disable" -verbose up 

migrate-down:
	migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/command_runner?sslmode=disable" -verbose down
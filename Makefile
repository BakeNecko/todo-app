# Defaults
database := "mongodb://localhost:27017/todo_app"
path := "migrations"
format := "json"

migrate_up:
	migrate -path $(path) -database $(database) -verbose up

migrate_down:
	migrate -path $(path) -database $(database) -verbose down

create_migration:
	 migrate create -ext $(format) -dir $(path) -seq $
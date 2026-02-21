include .env
export

MIGRATION_FOLDER="db/migrations"
DB_URL=${User}:${Passwd}@tcp\(${Addr}\)/${DBName}


migration:
	@echo "generation migration......in ${MIGRATION_FOLDER}"
	goose -dir ${MIGRATION_FOLDER} create ${name} sql 
	@echo "done!"

migrate-up:
	@echo "Applying Latest migration up..."
	goose -dir ${MIGRATION_FOLDER} mysql ${DB_URL} up
	@echo "done!"

migrate-down:
	@echo "Applying Last migration down..."
	goose -dir ${MIGRATION_FOLDER} mysql ${DB_URL} down
	@echo "done!"

run:
	@echo "starting project....."
	go run main.go

test:
	@echo DB_URL:${DB_URL}


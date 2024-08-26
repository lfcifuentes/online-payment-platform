include .env
# export PATH=$PATH:$(go env GOPATH)/bin
docs_generate:
	@echo "Generating docs...";
	swag init --parseDependency true;
	@echo "Docs generated successfully";

docs_format:
	@echo "Formatting code...";
	swag fmt
	@echo "Code formatted successfully";

add_migration:
	@if [ -z "$(name)" ]; then \
		echo "Error: 'name' parameter is required"; \
		exit 1; \
	fi
	@echo "Creating migration with name: $(name)...";
	migrate create -ext sql -dir app/db/migrations -seq $(name)
	@echo "Migration created successfully";

migrate:
	@echo "Migrating up...";
	migrate -path app/db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up
	@echo "Migration up successfully";

migrate_down:
	@echo "Reverting $${steps:-all} migrations...";
	migrate -path app/db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down $${steps:-}
	@echo "Migration down successfully";

test:
	@echo "Running tests...";
	go test -v ./...
	@echo "Tests passed successfully";

test_coverage:
	@echo "Running tests with coverage...";
	go test -coverprofile=coverage.out ./...
	@echo "Tests passed successfully";
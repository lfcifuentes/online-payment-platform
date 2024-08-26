# Api

# Run project
```go
go run main.go server
```
# Run seeders
```go
go run main.go seeders
```

# Makefile
Here is a list of the available commands in the Makefile and their description:

Generates the API documentation using swag.
```sh
make docs_generate
```

Formats the documentation code using swag.
```sh
make docs_format
```

### Migrations

Run migrations
```sh
make migrate
```

Add new migration
```sh
 make add_migration name=mew_migration_name
```

Rollback all migrations
```sh
make migrate_down
```

Rollback a specific migration
```sh
make migrate_down steps=1
```

### Testing

Run test
```sh
make test
```

Run test with coverage
```sh
make test_coverage
```

### Documentation URL

http://localhost:${HTTP_PORT}/swagger/index.html#/

### Project structure

```
/app
  /cmd
    server.go                # Entry point of the application
  /db
    /migrations              # Database migrations
  /infrastructure
    /http
      /middleware            # Api Rest Middleware
      /router                # API base route definitions
  /internal
    /adapters
      /pgsql                 # PostgreSQL database adapter
      /validator             # Playground Validator adapter
  /modules
    /auth
      /domain
        /services
          auth_services.go   # Domain services for authentication
      /application
        /usecases
          login_usecase.go   # Use case for user login
      /data
        /repositories
          auth_repository.go # Repository for user data
  /pkg
    jwt.go                   # JWT utility
docker-compose               # Docker compose configuration file
Makefile                     # Makefile with project commands
README.md                    # Project documentation
go.mod                       # Go module file
go.sum                       # Go dependencies file
```
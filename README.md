# Presensi Comindo Backend

## How to Install
1. Clone this repository
2. cp .env.example .env and fill the environment variables
3. Run `go mod tidy` to install dependencies
4. Run `go install github.com/cosmtrek/air@latest` to install air
5. Run `air` to start the server

## How to Install Via Docker
1. Read [Containerize README.md](https://github.com/ExeCiety/containerize-presensi-comindo/blob/main/README.md)

## Migration
### Create migration
```
go run main.go migrate --create=<migration_name>
```

### Migrate
```
go run main.go migrate
```

### Rollback
```
go run main.go migrate rollback
```
Note: rollback argument will only rollback the last migration, if you want to rollback multiple migration, you can use --step flag
Example: go run main.go migrate rollback --step=2

## Seeder
### Create Seeder
```
go run main.go seeder --create=<seeder_name>
```

### Seedling
```
go run main.go seeder
```

### Rollback
```
go run main.go seeder rollback
```
Note: rollback argument will only rollback the last seeder, if you want to rollback multiple migration, you can use --step flag
Example: go run main.go seeder rollback --step=2

### API Documentation
Postman Documenter
https://documenter.getpostman.com/view/7865721/2s9YkuXxJB

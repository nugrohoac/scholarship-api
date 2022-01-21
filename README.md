#Scholarship Api
Scholarship api provide api using graphQl

## Prerequisites
Make sure you have installed all of the following prerequisites on your development machine:
* go - [Download & Install](https://go.dev/dl/)
* graphql-go - [Documentation](https://github.com/graph-gophers/graphql-go)
* golang-migrate - [Install golang-migrate](https://github.com/golang-migrate/migrate)
* PostgreSQL
* mockery - [Install golang-migrate](https://github.com/vektra/mockery)


## Integration Test
This is will running integration test and unit test.
#### 1. Create file migrate
```bash
$ migrate create -ext sql -dir destination/directory name_migration
```

Each migration has an up and down migration.
```bash
1481574547_create_users_table.up.sql
1481574547_create_users_table.down.sql
```
* write up ddl at file.up.sql
* write reverse of ddl at file.down.sql

#### 2. Create file migrate
Make sure you have modified configuration at [postgre suite](./internal/postgresql/postgre_suite.go)
* driver     = "defauult postgres"
* host       = "default is localhost if running local"
* dbname     = "make sure you have been created database for testing"
* sslMode    = "deafult disable"
* userName   = "default postgres"
* password   = ""
* searchPath = "default is public"
* port       = "default is 5432"


#### Running Integration Test
```bash
$ go test ./... -cover -race
```

## Unit Test
This is just running unit test without integraton test. Make sure mocks is up-to-date.
* generate or update mock base on name of [interface](./contract.go)
#### generate or update mocks
```bash
$ mockery -name=name-of-interface
```
#### Running Integration Test
```bash
$ go test ./... --short -cover -race
```

## Update Schema Graphql
* run command [at line 13](./internal/graphql/schema/schema.go)


## Running Apps
#### Running Integration Test
* download all dependencies base on go.mod
```bash
$ go mod vendor
```
* running command
```bash
$ go run cmd/app/main.go
```

if your config env is custom
```bash
$ go run cmd/app/main.go -config-path=your-custom-env
```

## User status
* 0 status un verification
* 1 status verify but un complete profile
* 2 status verify and complete profile
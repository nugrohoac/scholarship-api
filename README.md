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
# Example
$ migrate create -ext sql -dir internal/postgresql/migrations create_table_applicant_score
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
This is just running unit test without integration test. Make sure mocks is up-to-date.
* generate or update mock base on name of [interface](src/business/contract.go)
#### generate or update mocks
```bash
$ mockery -name=name-of-interface
$ mockery --dir=source/directory --name=nameInterface --output=destination/directory
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
- 0 = not yet activated by email
- 1 = activated by email but incomplete profile
- 2 = complete profile (both student and sponsor)
- 3 = complete setup education (student)

## Scholarship status
- 0 = waiting for payment
- 1 = waiting for approve
- 2 = approve
- 3 = registration
- 4 = review
- 5 = announcement
- 6 = blazing email
- 7 = funding
- 8 = finish
- 9 = reject

## User scholarship status (Applicant status)
0 --> applied
1 --> reviewed
2 --> awardee
3 --> waiting_for_confirmation
4 --> confirmed
5 --> rejected


## Design Database
[![scholarship-drawio-2.png](https://i.postimg.cc/SNY7xgSd/scholarship-drawio-2.png)](https://postimg.cc/kB9tjFtR)
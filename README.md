# Game of Liberty

## Design Principles

**Refactoring is always part of the development**, so whenever you come up with a better solution or more elegant code, do not hesitate to change it.

We follow [Domain Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design)

## Development

### Local Development

```bash
docker-compose up
```

### Lint Check

```bash
docker compose exec web golangci-lint run
```

### DB Seeding

```bash
docker compose exec web make db-seed
```

## Cassandra Database

### Initialize Cassandra

Most of cases you don't have to run this, it is handled in `/scripts/check_cassandra.sh` script.

```bash
docker compose exec web make init-cassandra
```

### Connect to Cassandra

Connect to cassandra database.

```bash
docker compose exec web make connect-cassandra
```

### Plan New Cassandra Migration

Create new Cassandra migration file.

> do not forget to run the `cassandra-migrate-up` below to do the migration after the file is completed.

```bash
docker compose exec web make cassandra-plan-migrate FILE_NAME=${file_name_in_snake_case}
```

### Start Migrating Cassandra

```bash
docker compose exec web make cassandra-migrate-up
```

### Revert Cassandra Migration by 1 Version

```bash
docker compose exec web make cassandra-migrate-down
```

### Force Cassandra to Revert to Specific Version

```bash
docker compose exec web make cassandra-migrate-force CASSANDRA_MIGRATE_VERSION=${specifi_version}
```

## Postgres Database

### Plan New Postgres Migration

Create new Postgres migration file.

> do not forget to run the `postgres-migrate-up` below to do the migration after the file is completed.

```bash
docker compose exec web make postgres-plan-migrate FILE_NAME=${file_name_in_snake_case}
```

### Start Migrating Postgres

```bash
docker compose exec web make postgres-migrate-up
```

### Revert Postgres Migration by 1 Version

```bash
docker compose exec web make postgres-migrate-down
```

### Force Postgres to Revert to Specific Version

```bash
docker compose exec web make postgres-migrate-force POSTGRES_MIGRATE_VERSION=${specifi_version}
```

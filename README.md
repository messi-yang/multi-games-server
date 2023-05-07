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
docker compose exec web make lint
```

### Unit Test

```bash
docker compose exec web make test
```

### DB Seeding

```bash
docker compose exec web make db-seed
```

## Postgres Database

### Update Postgres Schema

```bash
docker compose exec postgres /usr/local/bin/pg_dump -U main --schema-only main > db/postgres/schema.sql
```

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

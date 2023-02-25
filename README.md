# Game of Liberty

## Design Principles

**Refactoring is always part of the development**, so whenever you come up with a better solution or more elegant code, do not hesitate to change it.

We follow [Domain Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design)

## Development

### Local Development

```bash
docker-compose up
```

## Database

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

Create new Cassandra migration file, do not forget to run the `start-cassandra-migrate` below to do the migration after the file is completed.

```bash
docker compose exec web make create-cassandra-migrate-file FILE_NAME=${file_name_in_snake_case}
```

### Start Migrating Cassandra

```bash
docker compose exec web make start-cassandra-migrate
```

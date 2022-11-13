# Game of Liberty

## Design Principles

**Refactoring is always part of the development**, so whenever you come up with a better solution or more elegant code, do not hesitate to change it.

We follow [Domain Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design)

## Development

### Start dev server with hotreload

```bash
docker-compose up
```

### Create new migration db file

```bash
docker compose exec livegameserver make create-migrate-db-file FILE_NAME=${file_name_in_snake_case}
```

### Migration db according to migration db files

```bash
docker compose exec livegameserver make migrate-db
```

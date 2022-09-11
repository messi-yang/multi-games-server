# Game of Liberty Computer

A server that keeps computing the next generation of the Game of Liberty.

## Design Principles

We follow [Clean Code](https://gist.github.com/wojteklu/73c6914cc446146b8b533c0988cf8d29).

**Refactoring is always part of the development**, so whenever you come up with a better solution or more elegant code, do not hesitate to change it.

We follow [Domain Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design)

## Development

We embrace Docker, so you can develop or deploy with just one docker-compose command.

But for sure in some cases you still need your Go installed.

### Setup configuration

If you don't have .env file yet, copy from .env.example and update the environment variables if you need.

```bash
cp .env.example .env
```

### Develop with hot reload

```bash
make dev
```

### Start production server

```bash
make build
```

### Develop with hot reload with Docker

```bash
docker-compose up dev
```

### Start production server with Docker

```bash
docker-compose up app
```

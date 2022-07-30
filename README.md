# Game of Liberty Computer

A server that keeps computing the next generation of the game of liberty.

## Design Principles

We follow [Clean Code](https://gist.github.com/wojteklu/73c6914cc446146b8b533c0988cf8d29).

**Refactoring is always part of the development**, so whenever you come up with a better solution or more elegant code, do not hesitate to change it.

We follow [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

![Clean Architecture](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

Think about how we deal with the dependencies with seriously, good architecture can help make us make progress faster and make life much much much much easier.

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

# Game of Liberty Computer

A server that keeps computing the next generation of the game of liberty.

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

### Start production server with Docker

```bash
docker-compose up
```

# Go Simple API

Simple example on how to run a simple API

## Docker build

```bash
docker build -t go-simple-api:test -f Dockerfile.multistage .
```

## Run

```bash
docker run --rm --read-only -p 8081:8081 go-simple-api:test
```

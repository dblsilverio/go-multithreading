# Go Expert

## Multithreading

* golang 1.21

### Building & Running Server

```bash
$ go build -o CEP cmd/server/main.go
$ ./CEP
```

### Testing

* [Swagger endpoint](http://localhost:3000/docs/index.html#/)

```bash
$ curl http://localhost:3000/cep/:cep
```

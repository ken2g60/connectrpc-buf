```
buf curl \
 --schema ./greet/v1/greet.proto \
 --protocol grpc \
 --http2-prior-knowledge \
 --data '{"name":"Jane"}' \
 http://localhost:8080/greet.v1.GreetService/Greet
```

```
buf curl \
 --schema . \
 --data '{"name":"Jane"}' \
 http://localhost:8080/greet.v1.GreetService/Greet
```

```
buf curl \
 --schema . \
 http://localhost:8080/greet.v1.GreetService/Profile
```

```
buf curl \
 --schema . \
 --data '{"request_id": "1"}' \
 http://localhost:8080/greet.v1.GreetService/Header

```

- buf dep update
- buf lint
- buf generate

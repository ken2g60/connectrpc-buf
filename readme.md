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

```
buf curl \
  --schema ./greet/v1/greet.proto \
  --header "Authorization: Bearer valid-token" \
  http://localhost:8080/greet.v1.GreetService/Profile
```

```
buf curl \
    --schema ./greet/v1/greet.proto \
    --protocol grpc \
    --http2-prior-knowledge \
    --data '{"name": "ken", "phone_number": "981-000-0000", "email": "ken@gmail.com", "password": "password-password"}' \
    http://localhost:8080/greet.v1.GreetService/CreateAccount
```

- buf dep update
- buf lint
- buf generate

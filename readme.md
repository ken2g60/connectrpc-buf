## How to start the project

### Prerequisites
- Go 1.27+ installed (check with `go version`)
- [Buf](https://docs.buf.build/installation) installed (`brew install buf` on macOS)

### Setup
```sh
# Install dependencies and generate code
buf dep update && buf lint && buf generate
```

### Run the server
```sh
go run ./cmd/server
```
The server will start on `localhost:8080`.

### Run the client example
In a separate terminal, execute:
```sh
go run ./cmd/client
```
You should see the greeting logged to the console.

You can also try the `buf curl` examples above to interact with the service.

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
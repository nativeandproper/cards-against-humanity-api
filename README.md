# Cards Against Humanity API
Public API for Cards Against Humanity Game

## API

## Development 

From project directory: 
```BASH
  source .env
 `go run *.go` 
 ```

## Migrations 
Goose doesn't support running migrations with a Golang binary from a config file yet (https://github.com/pressly/goose/pull/68). So, you need to pass in the config info via the CLI. To Run Goose Migrations from ${GOLANG_PATH}/src/cards-against-humanity/db directory:

```BASH 
go run cmd/main.go postgres "user={DB_USER_NAME} dbname={DB_NAME} sslmode=disable" up
```

## HTTP API Reference

### POST `/v1/signup`

Creates a new user.

#### Request

```json
{
  "email": "nativeandproper@gmail.com",
  "first_name": "Native",
  "last_name": "Proper",
  "password": "pass123"
}
```

#### Response

```json
{
  "success": "ok"
}
```

| Status Code | Description                |
| ----------- | -------------------------- |
| 201         | Created                    |
| 400         | Missing or invalid data    |
| 409         | User already exists        |
| 503         | Database not available     |

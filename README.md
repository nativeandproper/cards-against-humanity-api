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

### POST `/v1/user/signup`

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
| 409         | Email address taken        |
| 500         | Service is not available (SQL, SendGrid)   |

### PUT `/v1/user/signup`

Verifies the email address associated with a user.

#### Request

```json
{
  "verification_token": "7d977f8381d2542e04dcc0d4ce216205",
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
| 200         | User verified              |
| 400         | Token Expired              |
| 400         | User Not Found             |
| 503         | Service Unavailable        |

### POST `/v1/user/login`

Logs the user into a session. 

#### Request

```json
{
  "email": "nativeandproper@gmail.com",
  "password": "meepskeepbeep"
}
```

#### Response

```json
{
  "token": "da8494j4098973n43kn3n34k433kn34kdkd"
}
```

| Status Code | Description                |
| ----------- | -------------------------- |
| 201         | User verified              |
| 400         | User Not Found             |
| 503         | Service Unavailable        |

### POST `/v1/user/:userID/token`

Creates an API token for user. 

#### Request

#### Response

```json
{
  "id": 1, 
  "api_key": "dfdskds60960604098973n43kn3n34k433kn34kdkd",
  "created_at": "2018-02-09 05:28:34.945929", // UTC timestamp 
  "deleted_at": null
}
```

| Status Code | Description                |
| ----------- | -------------------------- |
| 201         | API Token Created          |
| 400         | User Not Found             |
| 503         | Service Unavailable        |

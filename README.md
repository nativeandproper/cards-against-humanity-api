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
  "last_name": "Proper"
}
```

#### Response

```json
{
  "email": "nativeandproper@gmail.com",
  "first_name": "Native",
  "last_name": "Proper"
}
```

| Status Code | Description                              |
| ----------- | ---------------------------------------- |
| 201         | Created                                  |
| 400         | Missing or invalid data                  |
| 409         | Email address taken                      |
| 500         | Service is not available (SQL, SendGrid) |

### PUT `/v1/signup`

Verifies the email address associated with a user.

#### Request

```json
{
  "verification_token": "7d977f8381d2542e04dcc0d4ce216205"
}
```

#### Response

```json
{
  "success": "ok"
}
```

| Status Code | Description         |
| ----------- | ------------------- |
| 200         | User verified       |
| 400         | Token Expired       |
| 400         | User Not Found      |
| 503         | Service Unavailable |

### POST `/v1/login`

Logs the user into a session.

#### Request

```json
{
  "email": "nativeandproper@gmail.com",
  "password": "meepskeepbeep"
}
```

#### Response

Sets token on cookie.

```json
{
  "success": "ok"
}
```

| Status Code | Description            |
| ----------- | ---------------------- |
| 201         | User verified          |
| 400         | Authentication Invalid |
| 400         | User Not Found         |
| 503         | Service Unavailable    |

### POST `/v1/logout`

Logs the user out of a session.

#### Request

Sets user token as unauthenticated.

```json
{
  "success": "ok"
}
```

| Status Code | Description         |
| ----------- | ------------------- |
| 200         | Successful          |
| 503         | Service Unavailable |

### POST `/v1/user/:userID/apikey`

Creates an API key for user.

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

| Status Code | Description         |
| ----------- | ------------------- |
| 201         | API Token Created   |
| 400         | User Not Found      |
| 503         | Service Unavailable |

### DELETE `/v1/user/:userID/apikey/:apiKey`

Deletes an API key for user.

#### Request

#### Response

```json
"ok"
```

| Status Code | Description         |
| ----------- | ------------------- |
| 200         | Successful          |
| 404         | API Key Not Found   |
| 503         | Service Unavailable |

### GET `/v1/user/:userID/apikey`

Retrieves list of API keys associated with user.

#### Request

#### Response

```json
{
  "apiKeys": [
    {
      "id": 1,
      "api_key": "dfdskds60960604098973n43kn3n34k433kn34kdkd",
      "created_at": "2018-02-09 05:28:34.945929",
      "deleted_at": null
    },
    {
      "id": 2,
      "api_key": "vsdadsalaskds6004098973n43kn3n34k4kn34sklsk",
      "created_at": "2018-03-02 05:28:34.945929",
      "deleted_at": "2018-03-11 18:40:51.130696"
    }
  ]
}
```

| Status Code | Description         |
| ----------- | ------------------- |
| 200         | Successful          |
| 404         | User Not Found      |
| 503         | Service Unavailable |

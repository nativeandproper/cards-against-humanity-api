# cards-against-humanity-api
Public API for Cards Against Humanity Game

To Run API: `go run main.go` 

//  Goose doesn't support running migrations with a Golang binary from a config file yet (https://github.com/pressly/goose/pull/68). So, you need to pass in the config info via the CLI.
To Run Goose Migrations from ${GOLANG_PATH}/src/cards-against-humanity/db directory:
`go run cmd/main.go postgres "user={DB_USER_NAME} dbname={DB_NAME} sslmode=disable" up `
package users

import (
	"database/sql"
	"time"
)

// User struct defines a new user
type User struct {
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Password    string    `json:"password"`
	AccountType string    `json:"account_type"`
	Timestamp   time.Time `json:"timestamp"`
}

// UserClient is a struct for user based actions
type UserClient struct {
	sqlClient *sql.DB
}

// NewUserClient creates a new instance of the User struct
func NewUserClient(sqlClient *sql.DB) *UserClient {
	return &UserClient{
		sqlClient: sqlClient,
	}
}

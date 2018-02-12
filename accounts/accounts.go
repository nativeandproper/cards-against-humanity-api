package accounts

import (
	"cards-against-humanity-api/sql"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/sendgrid/sendgrid-go"
)

var (
	// ErrEmailMustBeUnique indicates an email is already taken
	ErrEmailMustBeUnique = errors.New("Error: A user with this email address already exists")
	// ErrEmailVerificationNotDeliverable indicates the user provided an unreachable email address
	ErrEmailVerificationNotDeliverable = errors.New("Error: A verification email could not be sent to this email")
	// ErrEmailVerificationNotSent indicates an internal error. Log error internally and ask the user to re-try later
	ErrEmailVerificationNotSent = errors.New("Error: A verification email could not be sent at this time. Please re-try later")
	// ErrUserNotFound indicates a user associated with the email verification token could not be found
	ErrUserNotFound = errors.New("Error: user not found")
	// ErrUserVerificationTokenHasExpired indicates a user created an account, but did not verify account before token expired
	ErrUserVerificationTokenHasExpired = errors.New("Error: Verification token has expired")
)

// User struct defines a new user
type User struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

// AccountClient is a struct for user based actions
type AccountClient struct {
	databaseClient         *sql.DatabaseClient
	logger                 zerolog.Logger
	mailClient             *sendgrid.Client
	accountVerificationURL string
}

// NewAccountClient creates a new instance of the User struct
func NewAccountClient(databaseClient *sql.DatabaseClient, logger zerolog.Logger, mailClient *sendgrid.Client, accountVerificationURL string) *AccountClient {
	return &AccountClient{
		databaseClient,
		logger,
		mailClient,
		accountVerificationURL,
	}
}

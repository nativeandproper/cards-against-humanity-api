package sql

import (
	"cards-against-humanity-api/models"
	"database/sql"
	"github.com/pkg/errors"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

// GetUserByEmail looks up user by email address
func (dc *DatabaseClient) GetUserByEmail(email string) (*models.User, error) {
	user, err := models.Users(dc.sqlClient, Where("email=?", email)).One()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "GetUserByEmail: error checking if user exists")
	}
	return user, nil
}

func (dc *DatabaseClient) GetUserByID(userID int) (*models.User, error) {
	user, err := models.FindUser(dc.sqlClient, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "GetUserByID: error checking if user exists")
	}
	return user, nil
}

// CheckUserExistsByEmail checks if user with email exists
func (dc *DatabaseClient) CheckUserExistsByEmail(email string) (bool, error) {
	exists, err := models.Users(dc.sqlClient, Where("email=?", email)).Exists()
	if err != nil {
		return false, errors.Wrap(err, "CheckUserExistsByEmail: error checking if user exists")
	}

	return exists, nil
}

// InsertUser inserts a new user into the database
func (dc *DatabaseClient) InsertUser(email string, firstName string, lastName string, password []byte) error {
	var user models.User
	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email
	user.Password = password

	err := user.Insert(dc.sqlClient)
	if err != nil {
		return errors.Wrap(err, "InsertUser: error inserting user")
	}

	return nil
}

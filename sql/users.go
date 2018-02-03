package sql

import (
	"cards-against-humanity-api/models"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

// GetUserByEmail looks up user by email address
func (dc *DatabaseClient) GetUserByEmail(email string) (*models.User, error) {
	user, err := models.Users(dc.sqlClient, Where("email=?", email)).One()
	if err != nil {
		dc.logger.Error().Err(err).Msg("GetUserIDByEmail: error checking if user exists")
		return nil, err
	}

	return user, nil
}

// CheckUserExistsByEmail checks if user with email exists
func (dc *DatabaseClient) CheckUserExistsByEmail(email string) (bool, error) {
	exists, err := models.Users(dc.sqlClient, Where("email=?", email)).Exists()
	if err != nil {
		dc.logger.Error().Err(err).Msg("CheckUserExistsByEmail: error checking if user exists")
		return false, err
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
		dc.logger.Error().Err(err).Msg("InsertUser: error inserting user")
	}

	return nil
}

package accounts

import (
	"bytes"
	"cards-against-humanity-api/models"
	"crypto/rand"
	b64 "encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"html/template"
	"time"
)

const (
	tokenLength               = 32
	tokenExpirationInHours    = 24
	emailPlainTextContent     = "Click the link provided to verify your account!"
	emailSubject              = "Verify Email for Cards Against Humanity API Access"
	adminEmailAddress         = "nativeandproper@gmail.com"
	adminName                 = "Cards Against Humanity API Admin"
	emailVerificationTemplate = "./templates/email_verification.html"
)

func generateBase64Token() string {
	tokenBytes := make([]byte, tokenLength)
	rand.Read(tokenBytes)
	token := b64.StdEncoding.EncodeToString(tokenBytes)
	return token
}

// setTokenExpiration sets the expiration time for a user to verify account with token
func setTokenExpiration() time.Time {
	return time.Now().UTC().Add(tokenExpirationInHours * time.Hour)
}

// isExpired returns boolean if a date has passed
func isExpired(expiredAtDate time.Time) bool {
	return time.Now().After(expiredAtDate)
}

// CreateUserVerification creates and sends a user verification to new users
func (a *AccountClient) CreateUserVerification(email string) (*models.User, error) {
	token := generateBase64Token()
	expiration := setTokenExpiration()

	// Get UserID By User Email
	user, err := a.databaseClient.GetUserByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "CreateUserVerification: Error getting user by email")
	}

	// Store user and token association
	err = a.databaseClient.InsertUserVerification(user.ID, token, expiration)
	if err != nil {
		return nil, errors.Wrap(err, "CreateUserVerification: Error setting verification token")
	}

	// Send verification link to user
	err = a.SendUserEmailVerification(user.FirstName, user.Email, token)
	if err != nil {
		return nil, err
	}

	// Remove user password
	user.Password = make([]byte, 0)
	return user, nil
}

// UpdateUserVerification marks a user as verified
func (a *AccountClient) UpdateUserVerification(token string) error {
	// Get user verification by token
	verification, err := a.databaseClient.GetUserVerificationByToken(token)
	if err != nil {
		if err.Error() == "Not Found" {
			return ErrUserNotFound
		}
		return errors.Wrap(err, "UpdateUserVerification: Error getting user associated with token")
	}

	// If token has already been validated
	if verification.VerifiedAt.Valid == true {
		return nil
	}

	// If token as expired, reject
	isExpired := isExpired(verification.ExpiresAt)
	if isExpired {
		return ErrUserVerificationTokenHasExpired
	}

	// Set email as verified
	err = a.databaseClient.UpdateUserVerification(verification.ID)
	if err != nil {
		return errors.Wrap(err, "UpdateUserVerification: Error updating user as verified")
	}

	return nil
}

// ParseEmailTemplate parses the HTML template and associate with user data
func (a *AccountClient) ParseEmailTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", errors.Wrap(err, "ParseEmailTemplate: Error parsing template")
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", errors.Wrap(err, "ParseEmailTemplate: Error executing template")
	}

	templateString := buf.String()
	return templateString, nil
}

// SendUserEmailVerification sends verification token to users email
func (a *AccountClient) SendUserEmailVerification(name string, email string, userToken string) error {
	// Format email
	from := mail.NewEmail(adminName, adminEmailAddress)
	to := mail.NewEmail(name, email)
	verificationLink := fmt.Sprintf("%s/%s", a.accountVerificationURL, userToken)

	// Render template data
	templateData := struct {
		Name             string
		VerificationLink string
	}{
		Name:             name,
		VerificationLink: verificationLink,
	}

	// Associate template data with template
	htmlContent, err := a.ParseEmailTemplate(emailVerificationTemplate, templateData)
	if err != nil {
		return err
	}

	// Create message
	message := mail.NewSingleEmail(from, emailSubject, to, emailPlainTextContent, htmlContent)

	// Send email
	response, err := a.mailClient.Send(message)
	if err != nil {
		return err
	}

	// Return value based on email verification status
	switch response.StatusCode {
	case 200:
	case 202:
		return nil
	case 404:
		return ErrEmailVerificationNotDeliverable
	default:
		a.logger.Error().Err(err).Msg("SendUserEmailVerification: Error sending email verification")
		return ErrEmailVerificationNotSent
	}

	return nil
}

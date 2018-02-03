package accounts

import (
	"bytes"
	"cards-against-humanity-api/models"
	"crypto/rand"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"html/template"
	"time"
)

const (
	tokenLength               = 16
	tokenExpirationInHours    = 24
	emailPlainTextContent     = "Click the link provided to verify your account!"
	emailSubject              = "Verify Email for Cards Against Humanity API Access"
	adminEmailAddress         = "nativeandproper@gmail.com"
	adminName                 = "Cards Against Humanity API Admin"
	emailVerificationTemplate = "./templates/email_verification.html"
)

func generateToken() string {
	tokenBytes := make([]byte, tokenLength)
	rand.Read(tokenBytes)
	return fmt.Sprintf("%x", tokenBytes)
}

// setTokenExpiration sets the expiration time for a user to verify account with token
func setTokenExpiration() time.Time {
	return time.Now().Add(tokenExpirationInHours * time.Hour)
}

// CreateUserVerification creates and sends users a token to validate new users
func (a *AccountClient) CreateUserVerification(email string) (*models.User, error) {
	// Create token
	token := generateToken()

	// Set token expiration
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

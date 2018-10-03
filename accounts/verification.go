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
	mRand "math/rand"
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
	// remove slashes as verification code will be read as url param
	removeCharByte(tokenBytes, byte(63))
	token := b64.StdEncoding.EncodeToString(tokenBytes)
	return token
}

func randomInt(min int, max int) int {
	r := mRand.New(mRand.NewSource(time.Now().UnixNano()))
	p := r.Perm(max - min + 1)
	return p[min]
}

func removeCharByte(bytesArr []byte, charCode byte) {
	for i, b := range bytesArr {
		if b == charCode {
			bytesArr[i] = byte(randomInt(0, 62))
		}
	}
}

// setTokenExpiration sets the expiration time for a user to verify account with token
func setTokenExpiration() time.Time {
	return time.Now().UTC().Add(tokenExpirationInHours * time.Hour)
}

// isExpired returns boolean if date is passed current date
func isExpired(expiredAtDate time.Time) bool {
	return time.Now().After(expiredAtDate)
}

// CreateEmailVerification creates and sends an email verification to new users
func (a *AccountClient) CreateEmailVerification(email string) (*models.User, error) {
	token := generateBase64Token()
	expiration := setTokenExpiration()

	user, err := a.databaseClient.GetUserByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "CreateEmailVerification: Error getting user by email")
	}

	// Store user and token association
	err = a.databaseClient.InsertEmailVerification(user.ID, token, expiration)
	if err != nil {
		return nil, errors.Wrap(err, "CreateEmailVerification: Error setting token")
	}

	// send verification email to user
	err = a.SendEmailVerification(user.FirstName, user.Email, token)
	if err != nil {
		return nil, err
	}

	// Remove password
	user.Password = make([]byte, 0)
	return user, nil
}

// UpdateUserVerifyEmail confirms email associated with user account
func (a *AccountClient) UpdateUserVerifyEmail(token string) error {
	// get verification by token
	emailVerification, err := a.databaseClient.GetEmailVerificationToken(token)
	if err != nil {
		return errors.Wrap(err, "UpdateUserVerifyEmail: Error getting user associated with token")
	}
	if emailVerification == nil {
		return ErrTokenNotFound
	}

	// already validated
	if emailVerification.VerifiedAt.Valid == true {
		return nil
	}

	// has expired
	isExpired := isExpired(emailVerification.ExpiresAt)
	if isExpired {
		return ErrEmailVerificationTokenExpired
	}

	// set email as verified
	err = a.databaseClient.UpdateEmailVerification(emailVerification.ID)
	if err != nil {
		return errors.Wrap(err, "UpdateUserVerifyEmail: Error updating email as verified")
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

// SendEmailVerification sends verification email to email address associated with user's account
func (a *AccountClient) SendEmailVerification(name string, email string, userToken string) error {
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
		a.logger.Error().Err(err).Msg("SendEmailVerification: Error sending email verification")
		return ErrEmailVerificationNotSent
	}

	return nil
}

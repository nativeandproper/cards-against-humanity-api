package auth

import (
	"cards-against-humanity-api/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	*models.User
	jwt.StandardClaims
}

type AuthClient struct {
	AuthToken []byte
}

func New(authToken []byte) *AuthClient {
	return &AuthClient{
		AuthToken: authToken,
	}
}

func (a *AuthClient) Issue(user *models.User) (string, error) {
	expiresAt := time.Now().Add(time.Hour * 24).UTC().Unix()
	claims := &Claims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "cards-against-humanity-api",
			IssuedAt:  time.Now().UTC().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(a.AuthToken)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (a *AuthClient) parse(tokenStr string) (*jwt.Token, error) {
	// parse token
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.AuthToken), nil
	})
}

// Validate validates all the claims associated with a token
func (a *AuthClient) Validate(tokenStr string) (bool, map[string]interface{}, error) {
	var claims map[string]interface{}

	// parse and validate the token
	token, err := a.parse(tokenStr)
	if err != nil {
		return false, nil, fmt.Errorf("Error parsing jwt token: %s", err.Error())
	}

	for key, val := range claims {
		claims[key] = val
	}

	return token.Valid, claims, nil
}

// How do I validate the user properties passed into the token (i.e. last sign out date, userID exists)
// Where/how should I put the userID on context

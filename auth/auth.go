package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	name   string `json:"name"`
	userID string `json:"user_id"`
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

func (a *AuthClient) Issue(userID, name string) (string, error) {
	expiresAt := time.Now().Add(time.Hour * 24).UTC().Unix()
	claims := &Claims{
		userID,
		name,
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

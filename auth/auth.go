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

func (a *AuthClient) Validate(tokenStr string) (bool, error) {
	// parse token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.AuthToken), nil
	})
	if err != nil {
		return false, fmt.Errorf("Error parsing jwt token: %s", err.Error())
	}

	// validate claims
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	} else {
		return false, err
	}
}

package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/nativeandproper/cards-against-humanity-api/models"
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

func (a *AuthClient) parse(tokenStr string, validateToken bool) (*jwt.Token, error) {
	p := &jwt.Parser{
		SkipClaimsValidation: !validateToken,
	}

	// parse token
	return p.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.AuthToken), nil
	})
}

// ValidateWithClaims parses and validates token and claims
func (a *AuthClient) ValidateWithClaims(tokenStr string) (map[string]interface{}, error) {
	claims := make(map[string]interface{})

	// parse and validate the token
	token, err := a.parse(tokenStr, true)
	if err != nil {
		return nil, fmt.Errorf("Error parsing jwt token: %s", err.Error())
	}

	jwtTokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Error could not parse token claims")
	}

	for key, val := range jwtTokenClaims {
		claims[key] = val
	}

	return claims, nil
}

// IsValidToken validates the claims associated with a token
func (a *AuthClient) IsValidToken(tokenStr string) bool {
	// parse and validate the token
	token, err := a.parse(tokenStr, false)
	if err != nil {
		return false
	}

	// validate token claims
	if err := token.Claims.Valid(); err != nil {
		return false
	}

	return true
}

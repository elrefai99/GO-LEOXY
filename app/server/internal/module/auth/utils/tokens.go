package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IToken struct {
	Access      string
	Pending     string
	secretToken string
}

func NewToken(secretToken string) *IToken {
	return &IToken{
		secretToken: secretToken,
	}
}

func (n *IToken) CreateAccess(payload string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": payload,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	})

	signedToken, err := token.SignedString([]byte(n.secretToken))
	if err != nil {
		return "", err
	}

	n.Access = signedToken
	return n.Access, nil
}

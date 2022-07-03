package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"

	"time"
)

// TokenManger provides logic for JWT & Refresh generation and parsing
type TokenManger struct {
	signingKey string
}

func NewTokenManger(signingKey string) (*TokenManger, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}
	return &TokenManger{signingKey: signingKey}, nil
}

func (m *TokenManger) NewJWT(_id string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   _id,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *TokenManger) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}
		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}
	return claims["sub"].(string), nil
}

// NewRefreshToken generate random 32 bytes
func (m *TokenManger) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(fmt.Sprintf("%v", b)), nil
}

package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(UserId int) (string, error)
	ValidationToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var secretKey = []byte("bwastartup")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(UserId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = UserId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(secretKey)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidationToken(token string) (*jwt.Token, error) {
	encodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return encodedToken, err
	}

	return encodedToken, nil
}

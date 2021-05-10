package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(UserId int) (string, error)
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

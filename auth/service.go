package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(employeeId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(employeeId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["employee_id"] = employeeId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedTtoken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedTtoken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

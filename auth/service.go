package auth

import "github.com/golang-jwt/jwt"

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

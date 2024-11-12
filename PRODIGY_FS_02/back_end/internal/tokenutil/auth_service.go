package tokenutil

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
)

func VerifyToken(tokenString string, secret string) (*domain.JwtCustomClaims, error) {
    claims := &domain.JwtCustomClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil  
    })
    if err != nil {
        return nil, err
    }
    if !token.Valid {
        return nil, errors.New("token is invalid")
    }

    return claims, nil
}


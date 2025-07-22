package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtMaker struct {
	SecretKey string //Gett
}

func NewJwtMaker(secretKey string) *JwtMaker {
	return &JwtMaker{SecretKey: secretKey}
}

type TournamentClaims struct {
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

func (maker *JwtMaker) GenerateToken(id string, username string, role string) (string, error) {
	//vlaidate for 10 minutes
	duration := 30 * time.Minute
	claims := TournamentClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   id,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(maker.SecretKey))
}

func (maker *JwtMaker) VerifyToken(tokenString string) (*TournamentClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TournamentClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(maker.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TournamentClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}


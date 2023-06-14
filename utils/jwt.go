package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"

	"time"
)

func GenerateJWTAccessToken(userId any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	secretKey := []byte("Secret")
	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}
func ExtractPayloadFromJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte("Secret"), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return "", errors.New("malformed token")
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return "", errors.New("token is expired or not valid yet")
		} else {
			return "", errors.New("Invalid token")

		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userId, ok := claims["userId"].(string); ok {
			return userId, nil
		} else {
			return "", errors.New("Invalid token")

		}
	} else {
		return "", errors.New("Invalid token")

	}
}

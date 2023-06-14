package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/api/handlers"
	"time"
)

func GenerateJWTAccessToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	secretKey := []byte("Secret")
	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}
func ExtractPayloadFromJWT(tokenString string) (string, *handlers.UnauthorizedError) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &handlers.UnauthorizedError{Message: "invalid token"}
		}
		return []byte("Secret"), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return "", &handlers.UnauthorizedError{Message: "malformed token"}
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return "", &handlers.UnauthorizedError{Message: "token is expired or not valid yet"}
		} else {
			return "", &handlers.UnauthorizedError{Message: "invalid token"}
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userId, ok := claims["userId"].(string); ok {
			return userId, nil
		} else {
			return "", &handlers.UnauthorizedError{Message: "invalid token"}
		}
	} else {
		return "", &handlers.UnauthorizedError{Message: "invalid token"}
	}
}

package helpers

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

// token checker
func GetValidateToken(token, secretKey string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		}
		return []byte(secretKey), nil
	})
}

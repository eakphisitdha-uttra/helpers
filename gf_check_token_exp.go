package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CheckTokenExp(tokenString, secretKey string) error {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expTime := int64(claims["EXP"].(float64))

		// Comparing the current Unix time with the `exp` in the token
		if time.Now().Unix() > expTime {
			return fmt.Errorf("token is expired")
		}
	} else {
		fmt.Println(err)
		return err
	}

	return nil
}

package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtSignKey = "suprtsecretKey"

func GenerateJWtToken(email string, userid int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": userid,
		"epx":     time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSignKey))

	if err != nil {
		panic(err)
	}

	return tokenString, err
}

func VerifyToken(Token string) (int64, error) {
	ParseToken, err := jwt.Parse(Token, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {

			return nil, errors.New("unexpacted signin method")
		}

		return []byte(jwtSignKey), nil
	})

	if err != nil {
		return 0, errors.New("Could Not Parse Token")
	}

	if !ParseToken.Valid {
		return 0, errors.New("Token is not vaild")
	}
	claims, _ := ParseToken.Claims.(jwt.MapClaims)

	// email := claims["email"].(string)
	userid := int64(claims["user_id"].(float64))

	return userid, nil

}

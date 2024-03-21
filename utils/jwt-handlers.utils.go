package utils

import "github.com/golang-jwt/jwt"

func SignToken(claims jwt.Claims, secret string) string {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))

	if err != nil {
		return ""
	}

	return token
}

func VerifyToken(token string, secret string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}

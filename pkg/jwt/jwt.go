package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var claims = &jwt.RegisteredClaims{
	Issuer:    "authentication-server",
	Audience:  []string{"frontend"},
	ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 7*2)),
	NotBefore: jwt.NewNumericDate(time.Now()),
	IssuedAt:  jwt.NewNumericDate(time.Now()),
	ID:        "1",
}

var key = []byte("SECRET_PASSWORD")

func GenerateSignedString() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	str, err := token.SignedString(key)

	return str, err

}

type ValidationResult struct {
	Ok      bool
	Message string
}

func ValidateToken(signedString string) ValidationResult {
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return ValidationResult{
			Ok:      false,
			Message: err.Error(),
		}
	}

	if !token.Valid {

		var msg string

		switch true {
		case errors.Is(err, jwt.ErrTokenMalformed):
			msg = "Token is malformed"

		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			msg = "Invalid signature"

		case errors.Is(err, jwt.ErrTokenExpired):
			msg = "Token is expired"

		default:
			msg = "Token is invalid: " + err.Error()
		}

		return ValidationResult{
			Ok:      false,
			Message: msg,
		}
	}

	return ValidationResult{

		Ok:      true,
		Message: "Token successfuly validated",
	}
}

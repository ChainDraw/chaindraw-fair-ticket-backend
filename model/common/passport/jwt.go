package models

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenJwt struct {
}

// accoring to this username that generated the token.
type MyClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

// For HMAC signing method, the key can be any []byte. It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key for signing
// and validating.
// user
var userSecretString = []byte("0o7kJdgcJQHsv9hG")

// Signing() for building and signing a token
func (c TokenJwt) JwtUserSigning(userId string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	claim := MyClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour * time.Duration(1))), // token过期时间为7天
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                            // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                            // 生效时间
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(userSecretString)
	return tokenString, err
}

func userSecretKeyFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("0o7kJdgcJQHsv9hG"), nil // 这是我的secret
	}
}

// Validating() for parsing and validating a token
func (c TokenJwt) UserValidating(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, userSecretKeyFunc())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}

package jwtauth

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/guatom999/backend-challenge/utils"
)

type (
	AuthInterface interface {
		SignToken() string
	}

	Claims struct {
		UserId string `json:"user_id"`
	}

	UserClaims struct {
		*Claims
		jwt.RegisteredClaims
	}

	JwtToken struct {
		Secret []byte
		Claims *UserClaims `json:"claims"`
	}
)

func (a *JwtToken) SignToken() string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.Claims)

	signedToken, _ := token.SignedString(a.Secret)

	return signedToken

}

func ParseToken(pctx context.Context, secret string, tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error: unexpected singing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		log.Printf("Error: token parse error: %s", err.Error())
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("error: token format is invalid")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("error: token is expired")
		} else {
			return nil, errors.New("error: token is invalid")
		}
	}

	claims, ok := token.Claims.(*UserClaims)
	if ok {
		return claims, nil
	}

	return nil, errors.New("error: token is invalid")

}

func NewJwtToken(secret string, claims *Claims) AuthInterface {

	return &JwtToken{
		Secret: []byte(secret),
		Claims: &UserClaims{
			Claims: &Claims{
				UserId: claims.UserId,
			},
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "user-challenge",
				Subject:   "jwt-token",
				Audience:  []string{"user"},
				ExpiresAt: jwt.NewNumericDate(utils.GetLocalBkkTime().Add(time.Minute * 60)),
				NotBefore: jwt.NewNumericDate(utils.GetLocalBkkTime()),
				IssuedAt:  jwt.NewNumericDate(utils.GetLocalBkkTime()),
			},
		},
	}

}

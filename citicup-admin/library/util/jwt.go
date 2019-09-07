package util

import (
	"citicup-admin/internal/publicdata"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret []byte

type Claims struct {
	Param    string `json:"param"`
	Password string `json:"password"`
	Role     publicdata.Role `json:"role"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(param, password string, role publicdata.Role) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Hour)

	claims := Claims{
		param,
		password,
		role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "citicup",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}


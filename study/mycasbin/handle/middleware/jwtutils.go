package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserID int64 `json:"userid"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(userid int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(72 * time.Hour)

	claims := Claims{
		UserID: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "phenglei",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	js := "phenglei"
	jwtSecret := []byte(js)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	js := "phenglei"
	jwtSecret := []byte(js)
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

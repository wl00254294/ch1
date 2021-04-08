package jwtctrl

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf" //私钥
)

type CustomClaims struct {
	UserId    int64
	RequestIp string
	jwt.StandardClaims
}

func GenJWToken(efftime int, userid int64, user string, ip string) string {
	customClaims := &CustomClaims{
		UserId:    userid, //用户id
		RequestIp: ip,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(efftime) * time.Second).Unix(),
			Issuer:    user,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return tokenString
}

func ParseToken(tokenString string, reqip string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("exception")
		}

		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {

		if claims.RequestIp == reqip {
			return claims, nil
		} else {
			return nil, fmt.Errorf("exception: requst ip changed!")
		}

	} else {
		return nil, err
	}
}
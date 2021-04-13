package jwtctrl

import (
	"fmt"
	"time"

	dao "eric.com/go/ch1/dao"
	"github.com/dgrijalva/jwt-go"
)

const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf" //私钥
)

type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

type CustomToken struct {
	Userid    int64     `json:"userid"`
	Token     string    `json:"token"`
	Expiresat time.Time `json:"expiresat"`
	Message   string    `json:"message"`
}

func GenJWToken(efftime int, user dao.User) CustomToken {

	var userid int64
	var expirtime int64
	if user.Username == "eric" && user.Password == "eric" {
		userid = 1
		expirtime = time.Now().Add(time.Duration(efftime) * time.Second).Unix()
	} else {

		return CustomToken{0, "", time.Now(), "Username/Password invaild!"}
	}

	customClaims := &CustomClaims{
		UserId: userid, //用户id
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirtime,
			Issuer:    user.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		fmt.Println(err)
		return CustomToken{0, "", time.Now(), "Token Generate Fail!"}
	}

	return CustomToken{userid, tokenString, time.Unix(expirtime, 0), "OK"}
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("exception")
		}

		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {

		return claims, nil
	} else {
		return nil, err
	}
}

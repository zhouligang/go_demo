package utils

import (
	"errors"
	"gin-web-scaffolding/settings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// @file      : jwt.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

var mySecret = []byte("八宝糖")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多的信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成jwt token
func GenJWTToken(userID int64, username string) (string, error) {
	// 创建一个我们自己声明的数据
	c := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(settings.Conf.JwtExpire)).Unix(), // 过期时间
			Issuer:    settings.Conf.Name,                                            //签发人
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完成的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken 解析jwt token
func ParseJWTToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

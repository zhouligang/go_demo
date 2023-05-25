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

// CustomClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多的信息，都可以添加到这个结构体中
type CustomClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenJWTToken 生成jwt token
func GenJWTToken(userID int64, username string) (accessToken, refreshToken string, err error) {
	// 创建一个我们自己声明的数据
	c := CustomClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(settings.Conf.JwtAccessTokenExpire) * time.Second).Unix(), // 过期时间
			Issuer:    settings.Conf.Name,                                                                     //签发人
		},
	}

	// 使用指定的签名方法创建签名对象
	// 使用指定的secret签名并获得完成的编码后的字符串token
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	if err != nil {
		return
	}

	// refresh token中不需要存任何自定义数据
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(settings.Conf.JwtRefreshTokenExpire) * time.Second).Unix(), //过期时间
		Issuer:    settings.Conf.Name,
	}).SignedString(mySecret)
	if err != nil {
		return
	}
	return
}

// ParseJWTToken 解析jwt token
func ParseJWTToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	var cClaims = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, cClaims, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return cClaims, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新AccessToken
func RefreshToken(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	// 解析refreshToken，验证是否有效以及是否过期，如果过期直接返回
	_, err = jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return
	}

	// 从旧的accessToken中解析出claims数据
	var cClaims CustomClaims
	_, err = jwt.ParseWithClaims(accessToken, &cClaims, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})

	if err != nil {
		v, _ := err.(*jwt.ValidationError) // 进行类型转换
		// 当accessToken是过期错误，并且refreshToken没有过期时就创建一个新的accessToken
		if v.Errors == jwt.ValidationErrorExpired {
			return GenJWTToken(cClaims.UserID, cClaims.Username)
		}
	}
	return
}

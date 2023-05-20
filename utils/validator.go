package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// @file      : validator.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

// 对validator进行定制

// 定义一个全局的翻译器
var ValidatorTrans ut.Translator

// InitTrans 初始化
func InitValidatorTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		// 得到的效果是返回的错误信息中是结构体字段tag的内容而不是结构体字段的名称
		// 如"ParamSignUp.RePassword": "RePassword为必填字段" -->  "ParamSignUp.RePassword": "repassword为必填字段"
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境
		//后面的参数是应该支持的语言环境(可以支持多个)
		uni := ut.New(enT, zhT, enT)

		// locale通常取决于http请求头的Accept-Language
		var ok bool
		ValidatorTrans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, ValidatorTrans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, ValidatorTrans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, ValidatorTrans)
		}
		return
	}
	return
}

// 去除提示信息中的结构体名称
// "ParamSignUp.repassword": "repassword为必填字段" ==> "repassword": "repassword为必填字段"
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

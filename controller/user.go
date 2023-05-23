package controller

import (
	"errors"
	"fmt"
	"gin-web-scaffolding/dao/mysql"
	"gin-web-scaffolding/logic"
	"gin-web-scaffolding/models"
	"gin-web-scaffolding/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// @file      : user.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

// SignUpHandler 用户注册的处理函数
// @Summary 用户注册
// @Description 用户注册
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamsSignUp false "注册所需参数"
// @Router /signup [post]
func SignUpHandler(context *gin.Context) {
	// 获取参数和校验参数
	var pUser = new(models.ParamsSignUp)
	if err := context.ShouldBindJSON(pUser); err != nil {
		zap.L().Error("SignUp with invalid params", zap.Error(err))
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 表示不是validator.ValidatorErrors类型的错误，无法使用翻译器，直接返回
			ResponseError(context, CodeInvalidParam)
			return
		} else {
			// 对validator.ValidatorErrors类型的错误进行翻译
			ResponseErrorWithMsg(context, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(utils.ValidatorTrans)))
			return
		}
	}

	// 获取到了参数并且参数校验通过
	if err := logic.SignUp(pUser); err != nil {
		zap.L().Error("SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(context, CodeUserExist)
			return
		} else {
			ResponseError(context, CodeServerBusy)
			return
		}

	}

	// 返回响应
	ResponseSuccess(context, nil)
}

// LoginHandler 用户登录的处理函数
// @Summary 用户登录接口
// @Description 处理用户登录流程
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamsLogin false "用户登录所需参数"
// @Router /login [post]
func LoginHandler(context *gin.Context) {
	// 处理请求参数并验证
	p := new(models.ParamsLogin)
	if err := context.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid params", zap.Error(err))
		// 判断是否是validator.ValidatorErrors类型的错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 不是validator.ValidatorErrors类型的错误，直接返回
			ResponseError(context, CodeInvalidParam)
			return
		} else {
			ResponseErrorWithMsg(context, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(utils.ValidatorTrans)))
			return
		}
	}

	// 用户登录业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExists) {
			ResponseError(context, CodeUserNotExists)
			return
		} else {
			ResponseError(context, CodeServerBusy)
			return
		}
	}
	// 返回响应
	zap.L().Info("Login success", zap.String("username", user.Username))
	ResponseSuccess(context, gin.H{
		"user_id":       fmt.Sprintf("%d", user.UserID), //如果id大于1<<53-1的话，前端显示就会出现错误；因此转换成字符串给前端
		"user_name":     user.Username,
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})
}

package controller

import (
	"errors"
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

// @Summary 用户注册
// @Description 用户注册
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamsSignUp false "注册所需参数"
// @Router /signup [post]
// SignUpHandler 用户注册的处理函数
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

func LoginHandler(context *gin.Context) {
	ResponseSuccess(context, "LoginHandler")
}

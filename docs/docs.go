// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "八宝糖",
            "email": "1013269096@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "处理用户登录流程",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "用户登录所需参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.ParamsLogin"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/signup": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "注册所需参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.ParamsSignUp"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.ParamsLogin": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.ParamsSignUp": {
            "type": "object",
            "required": [
                "email",
                "password",
                "repassword",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "repassword": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1.0.0",
	Host:             "127.0.0.1",
	BasePath:         "/api/v1/",
	Schemes:          []string{},
	Title:            "GinWeb脚手架项目",
	Description:      "GinWeb脚手架项目",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "marchsoft@golang",
            "url": "http://marchsoft.cn/"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/code": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Author：JiaKunLi 2021/01/26",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "系统：系统授权接口 Authorization Controller"
                ],
                "summary": "获取图片验证码",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseCode"
                        }
                    }
                }
            }
        },
        "/api/auth/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Author：JiaKunLi 2021/01/26 获得身份令牌",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "系统：系统授权接口 Authorization Controller"
                ],
                "summary": "登录授权接口",
                "parameters": [
                    {
                        "description": "查询参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/dto.UserLoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseLogin"
                        }
                    }
                }
            }
        },
        "/api/file/uploadFile": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Author：JiaKunLi 2021/01/27",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件：文件管理 File Controller"
                ],
                "summary": "文件上传（任意类型文件）",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseFile"
                        }
                    }
                }
            }
        },
        "/api/file/uploadImage": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Author：JiaKunLi 2021/01/27",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文件：文件管理 File Controller"
                ],
                "summary": "文件上传（图片）",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models._ResponseFile"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.UserLoginDto": {
            "type": "object",
            "required": [
                "code",
                "password",
                "username",
                "uuid"
            ],
            "properties": {
                "code": {
                    "description": "验证码",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                },
                "uuid": {
                    "description": "验证码id",
                    "type": "string"
                }
            }
        },
        "models._ResponseCode": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "base64验证码",
                    "type": "string"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                },
                "uuid": {
                    "description": "验证码id",
                    "type": "string"
                }
            }
        },
        "models._ResponseFile": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "object",
                    "$ref": "#/definitions/public.FileResponse"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "models._ResponseLogin": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "object",
                    "properties": {
                        "token": {
                            "description": "授权令牌",
                            "type": "string"
                        }
                    }
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "public.FileResponse": {
            "type": "object",
            "properties": {
                "full_path": {
                    "description": "文件完整地址",
                    "type": "string"
                },
                "name": {
                    "description": "文件名",
                    "type": "string"
                },
                "path": {
                    "description": "文件相对地址",
                    "type": "string"
                },
                "size": {
                    "description": "文件大小",
                    "type": "integer"
                },
                "type": {
                    "description": "文件类型",
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1.0",
	Host:        "127.0.0.1:8977",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "go-sword项目接口文档",
	Description: "基于gin的后台通用框架",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
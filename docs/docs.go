// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/getCaptcha": {
            "get": {
                "description": "获取验证码",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "captcha"
                ],
                "summary": "获取验证码",
                "operationId": "get-captcha",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/form.CaptchaResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "get": {
                "description": "获取token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户名密码登录",
                "operationId": "user-login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/form.LoginResponse"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "接收偏移和限制量，返回对应的ping记录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "分页获取ping的记录信息",
                "operationId": "list-ping",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "偏移量",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页记录数",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ping.PingRecord"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "接收主键PingID,根据PingID更新该条记录的message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "更新记录信息",
                "operationId": "update_ping",
                "parameters": [
                    {
                        "type": "string",
                        "description": "any string",
                        "name": "ping_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "any string",
                        "name": "updatemessage",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "接收一个字符串，返回这个字符串加上\", too\"后缀",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "测试服务是否正常启动",
                "operationId": "ping",
                "parameters": [
                    {
                        "type": "string",
                        "description": "any string",
                        "name": "message",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "接收主键PingID,根据PingID删除该条记录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "删除记录信息",
                "operationId": "delete_ping",
                "parameters": [
                    {
                        "type": "string",
                        "description": "any string",
                        "name": "ping_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "管理员删除用户个人信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "删除用户信息",
                "operationId": "user-deletesysuser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/form.DeleteSysUserRequest"
                        }
                    }
                }
            }
        },
        "/user/insertsysuser": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "管理员创建用户个人信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "创建用户信息",
                "operationId": "user-insertsysuser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "角色",
                        "name": "role",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/form.InsertSysUserRequest"
                        }
                    }
                }
            }
        },
        "/user/me": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "用户查看个人信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "查看个人信息",
                "operationId": "user-me",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/form.ProfileResponse"
                        }
                    }
                }
            }
        },
        "/verifyCaptcha": {
            "post": {
                "description": "验证验证码",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "captcha"
                ],
                "summary": "验证验证码",
                "operationId": "verify-captcha",
                "parameters": [
                    {
                        "type": "string",
                        "description": "验证码id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "验证码内容",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "form.CaptchaResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "form.DeleteSysUserRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "form.InsertSysUserRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "role",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "form.LoginResponse": {
            "type": "object",
            "properties": {
                "duration": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "form.ProfileResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "ping.PingRecord": {
            "type": "object",
            "properties": {
                "createTime": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "pingID": {
                    "type": "string"
                },
                "reply": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Version:     "v0.0.1",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Ferry API",
	Description: "工单系统",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}

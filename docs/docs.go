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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/users": {
            "get": {
                "description": "查询多条数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users 用户"
                ],
                "summary": "查询多条数据",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "分页索引",
                        "name": "current",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "分页大小",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "default": true,
                        "description": "是否分页",
                        "name": "pagination",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "用户名称",
                        "name": "user_name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{status:\"OK\", data:响应数据}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/schema.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/schema.UserQueryResult"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "{code:400, status:\"OK\", message:\"请求参数错误\"}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorItem"
                        }
                    },
                    "404": {
                        "description": "{code:404, status:\"OK\", message:\"资源不存在\"}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorItem"
                        }
                    }
                }
            },
            "post": {
                "description": "注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users 用户"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "description": "用户",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{staus:\"OK\", data:响应数据}",
                        "schema": {
                            "$ref": "#/definitions/schema.StatusResult"
                        }
                    },
                    "400": {
                        "description": "{code:400, status:\"OK\", message:\"请求参数错误\"}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorItem"
                        }
                    },
                    "404": {
                        "description": "{code:404, status:\"OK\", message:\"资源不存在\"}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorItem"
                        }
                    }
                }
            }
        },
        "/api/users/:id/detail": {
            "get": {
                "description": "查询数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users 用户"
                ],
                "summary": "查询数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{staus:\"OK\", data:响应数据}",
                        "schema": {
                            "$ref": "#/definitions/schema.User"
                        }
                    },
                    "400": {
                        "description": "{code:400, status:\"OK\", message:\"请求参数错误\"}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorItem"
                        }
                    },
                    "404": {
                        "description": "{code:404, status:\"OK\", message:\"资源不存在\"}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorItem"
                        }
                    }
                }
            }
        },
        "/api/users/{id}": {
            "get": {
                "description": "查询数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users 用户"
                ],
                "summary": "查询数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "jwt",
                        "name": "authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{staus:\"OK\", data:响应数据}",
                        "schema": {
                            "$ref": "#/definitions/schema.User"
                        }
                    },
                    "400": {
                        "description": "{code:400, status:\"OK\", message:\"请求参数错误\"}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorItem"
                        }
                    },
                    "404": {
                        "description": "{code:404, status:\"OK\", message:\"资源不存在\"}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorItem"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.ErrorItem": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "错误码",
                    "type": "integer"
                },
                "message": {
                    "description": "错误信息",
                    "type": "string"
                },
                "staus": {
                    "type": "string"
                }
            }
        },
        "schema.PaginationResult": {
            "type": "object",
            "properties": {
                "current": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "schema.StatusResult": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "schema.SuccessResult": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "返回数据",
                    "type": "object"
                },
                "status": {
                    "description": "\"OK\"",
                    "type": "string"
                }
            }
        },
        "schema.User": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "name": {
                    "description": "用户名",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "telephone": {
                    "description": "手机号",
                    "type": "string"
                },
                "user_id": {
                    "description": "用户ID",
                    "type": "string"
                }
            }
        },
        "schema.UserQueryResult": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.User"
                    }
                },
                "pagination": {
                    "$ref": "#/definitions/schema.PaginationResult"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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

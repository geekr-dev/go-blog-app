// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/geekr-dev/go-blog-app",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/tags": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取标签列表",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "标签名称",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "Object"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "type": "Object"
                        }
                    },
                    "500": {
                        "description": "服务错误",
                        "schema": {
                            "type": "Object"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增标签",
                "parameters": [
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "标签名称",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "作者",
                        "name": "created_by",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "Object"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "type": "Object"
                        }
                    },
                    "500": {
                        "description": "服务错误",
                        "schema": {
                            "type": "Object"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改标签",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "标签ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "标签名称",
                        "name": "name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "编辑",
                        "name": "updated_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "Object"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "type": "Object"
                        }
                    },
                    "500": {
                        "description": "服务错误",
                        "schema": {
                            "type": "Object"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除标签",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "标签ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "Object"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "type": "Object"
                        }
                    },
                    "500": {
                        "description": "服务错误",
                        "schema": {
                            "type": "Object"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Blog App",
	Description:      "基于 Go 构建简单博客系统",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
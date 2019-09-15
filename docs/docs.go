// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-09-15 10:13:37.165219 +0800 CST m=+0.287375396

package docs

import (
	"bytes"
	"encoding/json"

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api//market/all": {
            "get": {
                "description": "Get All Market.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Market"
                ],
                "summary": "获取所有行情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页号",
                        "name": "pageNum",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Markets"
                        }
                    },
                    "404": {
                        "description": "Resource not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/companies": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "根据Id获取公司信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "公司id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Company"
                        }
                    },
                    "404": {
                        "description": "Resource not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update the company information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "更新公司信息",
                "responses": {
                    "200": {},
                    "404": {
                        "description": "Resource not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Delete company",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "添加新的公司",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Company"
                        }
                    },
                    "404": {
                        "description": "Resource not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete company",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "删除指定公司信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "公司id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "404": {
                        "description": "Resource not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/companies/list": {
            "get": {
                "description": "Get All Companies.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "获取所有公司",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页号",
                        "name": "pageNum",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Companies"
                        }
                    },
                    "404": {
                        "description": "Resource not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/companies/query": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "模糊搜索公司",
                "parameters": [
                    {
                        "type": "string",
                        "description": "关键字，按照关键字搜索名称",
                        "name": "key",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页号",
                        "name": "pageNum",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Company"
                        }
                    },
                    "404": {
                        "description": "Resource not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/users/check": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "确认质押方的注册申请结果",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "true 同意; false: 不同意，那么该用户将被删除",
                        "name": "checkResult",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
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
                    },
                    "400": {
                        "description": "错误的query",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/users/count": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "查询用户数量",
                "parameters": [
                    {
                        "type": "string",
                        "description": "根据该query分辨普通用户和质押方.用USER表示普通用户,用INVESTOR表示质押方.其他输入无效",
                        "name": "role",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户数量",
                        "schema": {
                            "type": "int"
                        }
                    },
                    "400": {
                        "description": "错误的query",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/users/investor": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取质押方列表,可以根据是否注册成功进行筛选",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "true 已审批通过; false待审批",
                        "name": "status",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页号",
                        "name": "pageNum",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Investors"
                        }
                    },
                    "400": {
                        "description": "错误的query",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Companies": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "businessScope": {
                        "description": "经营范围",
                        "type": "string"
                    },
                    "chairman": {
                        "description": "法人代表",
                        "type": "string"
                    },
                    "city": {
                        "description": "所在城市",
                        "type": "string"
                    },
                    "email": {
                        "description": "电子邮件",
                        "type": "string"
                    },
                    "employees": {
                        "description": "员工人数",
                        "type": "integer"
                    },
                    "exchange": {
                        "description": "交易所代码 ，SSE上交所 SZSE深交所",
                        "type": "string"
                    },
                    "id": {
                        "type": "integer"
                    },
                    "introduction": {
                        "description": "公司介绍",
                        "type": "string"
                    },
                    "mainBusiness": {
                        "description": "主要业务及产品",
                        "type": "string"
                    },
                    "manager": {
                        "description": "总经理",
                        "type": "string"
                    },
                    "market": {
                        "type": "object",
                        "$ref": "#/definitions/model.Market"
                    },
                    "marketID": {
                        "description": "市场行情与公司一对一",
                        "type": "integer"
                    },
                    "office": {
                        "description": "办公室",
                        "type": "string"
                    },
                    "province": {
                        "description": "所在省份",
                        "type": "string"
                    },
                    "regCapital": {
                        "description": "注册资本",
                        "type": "string"
                    },
                    "secretary": {
                        "description": "董秘",
                        "type": "string"
                    },
                    "setupDate": {
                        "description": "创立日期",
                        "type": "string"
                    },
                    "tsCode": {
                        "description": "股票代码",
                        "type": "string"
                    },
                    "website": {
                        "description": "公司主页URL",
                        "type": "string"
                    }
                }
            }
        },
        "model.Company": {
            "type": "object",
            "properties": {
                "businessScope": {
                    "description": "经营范围",
                    "type": "string"
                },
                "chairman": {
                    "description": "法人代表",
                    "type": "string"
                },
                "city": {
                    "description": "所在城市",
                    "type": "string"
                },
                "email": {
                    "description": "电子邮件",
                    "type": "string"
                },
                "employees": {
                    "description": "员工人数",
                    "type": "integer"
                },
                "exchange": {
                    "description": "交易所代码 ，SSE上交所 SZSE深交所",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "description": "公司介绍",
                    "type": "string"
                },
                "mainBusiness": {
                    "description": "主要业务及产品",
                    "type": "string"
                },
                "manager": {
                    "description": "总经理",
                    "type": "string"
                },
                "market": {
                    "type": "object",
                    "$ref": "#/definitions/model.Market"
                },
                "marketID": {
                    "description": "市场行情与公司一对一",
                    "type": "integer"
                },
                "office": {
                    "description": "办公室",
                    "type": "string"
                },
                "province": {
                    "description": "所在省份",
                    "type": "string"
                },
                "regCapital": {
                    "description": "注册资本",
                    "type": "string"
                },
                "secretary": {
                    "description": "董秘",
                    "type": "string"
                },
                "setupDate": {
                    "description": "创立日期",
                    "type": "string"
                },
                "tsCode": {
                    "description": "股票代码",
                    "type": "string"
                },
                "website": {
                    "description": "公司主页URL",
                    "type": "string"
                }
            }
        },
        "model.Investor": {
            "type": "object",
            "properties": {
                "accountStatus": {
                    "description": "账户状态,0 : 待审核, 1: 审核通过. 如果审核失败就删除账户",
                    "type": "integer"
                },
                "avatar": {
                    "type": "string"
                },
                "bankEmail": {
                    "type": "string"
                },
                "bankName": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "swiftCode": {
                    "type": "string"
                }
            }
        },
        "model.Investors": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "accountStatus": {
                        "description": "账户状态,0 : 待审核, 1: 审核通过. 如果审核失败就删除账户",
                        "type": "integer"
                    },
                    "avatar": {
                        "type": "string"
                    },
                    "bankEmail": {
                        "type": "string"
                    },
                    "bankName": {
                        "type": "string"
                    },
                    "email": {
                        "type": "string"
                    },
                    "id": {
                        "type": "integer"
                    },
                    "password": {
                        "type": "string"
                    },
                    "swiftCode": {
                        "type": "string"
                    }
                }
            }
        },
        "model.Market": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "成交额（亿）",
                    "type": "number"
                },
                "change": {
                    "description": "涨跌额",
                    "type": "number"
                },
                "close": {
                    "description": "收盘价",
                    "type": "number"
                },
                "highest": {
                    "description": "最高价",
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "lowest": {
                    "description": "最低价",
                    "type": "number"
                },
                "open": {
                    "description": "开盘价",
                    "type": "number"
                },
                "pctChg": {
                    "description": "涨跌幅",
                    "type": "number"
                },
                "preClose": {
                    "description": "昨日收盘价",
                    "type": "number"
                },
                "tradeDate": {
                    "description": "交易日期",
                    "type": "string"
                },
                "tscode": {
                    "type": "string"
                },
                "vol": {
                    "description": "成交量",
                    "type": "number"
                }
            }
        },
        "model.Markets": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "amount": {
                        "description": "成交额（亿）",
                        "type": "number"
                    },
                    "change": {
                        "description": "涨跌额",
                        "type": "number"
                    },
                    "close": {
                        "description": "收盘价",
                        "type": "number"
                    },
                    "highest": {
                        "description": "最高价",
                        "type": "number"
                    },
                    "id": {
                        "type": "integer"
                    },
                    "lowest": {
                        "description": "最低价",
                        "type": "number"
                    },
                    "open": {
                        "description": "开盘价",
                        "type": "number"
                    },
                    "pctChg": {
                        "description": "涨跌幅",
                        "type": "number"
                    },
                    "preClose": {
                        "description": "昨日收盘价",
                        "type": "number"
                    },
                    "tradeDate": {
                        "description": "交易日期",
                        "type": "string"
                    },
                    "tscode": {
                        "type": "string"
                    },
                    "vol": {
                        "description": "成交量",
                        "type": "number"
                    }
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
var SwaggerInfo = swaggerInfo{ Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface {}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}

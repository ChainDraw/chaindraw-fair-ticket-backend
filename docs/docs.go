// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2024-05-22 11:16:40.498758 +0800 CST m=+0.052833417

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample API for Your Project using Gin and Swagger",
        "title": "Your Project API",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "/api/v1",
    "paths": {
        "/concert_add": {
            "post": {
                "description": "Add a new concert with the given details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Concert"
                ],
                "summary": "Add a new concert",
                "parameters": [
                    {
                        "description": "Concert Add Request",
                        "name": "concertAddReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/commonreq.ConcertAddReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/commonresp.CommitResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/commonresp.CommitResp"
                        }
                    }
                }
            }
        },
        "/concerts": {
            "get": {
                "description": "Get list of concerts by IDs with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Concert"
                ],
                "summary": "Get list of concerts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Comma-separated list of concert IDs",
                        "name": "ids",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/commonresp.ConcertListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/commonresp.ConcertListResponse"
                        }
                    }
                }
            }
        },
        "/review_concert": {
            "post": {
                "description": "Review concerts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Concert"
                ],
                "summary": "Review concerts",
                "parameters": [
                    {
                        "description": "Concert Review Request",
                        "name": "concertReviewReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/commonreq.ConcertReViewReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/commonresp.ConcertListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/commonresp.ConcertListResponse"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "get": {
                "description": "Get personal information",
                "produces": [
                    "application/json"
                ],
                "summary": "Get personal information",
                "responses": {
                    "200": {
                        "description": "logout retrieved successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/nonce": {
            "get": {
                "description": "Generate nonce for session",
                "produces": [
                    "application/json"
                ],
                "summary": "Generate nonce",
                "responses": {
                    "200": {
                        "description": "Nonce generated successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/personal_information": {
            "get": {
                "description": "Get personal information",
                "produces": [
                    "application/json"
                ],
                "summary": "Get personal information",
                "responses": {
                    "200": {
                        "description": "Personal information retrieved successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/verify": {
            "post": {
                "description": "Verify signature with message and nonce",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Verify signature",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/RequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Verification success",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Verification failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "commonreq.ConcertAddReq": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "演唱会地址",
                    "type": "string"
                },
                "concert_date": {
                    "description": "演唱会日期",
                    "type": "string"
                },
                "concert_id": {
                    "description": "演唱会ID",
                    "type": "string"
                },
                "concert_img": {
                    "description": "演唱会图片",
                    "type": "string"
                },
                "concert_name": {
                    "description": "演唱会名称",
                    "type": "string"
                },
                "concert_status": {
                    "description": "演唱会状态（0: 未开始, 1: 已过期, 2: 已取消）",
                    "type": "integer"
                },
                "lottery_end_date": {
                    "description": "抽奖结束时间",
                    "type": "string"
                },
                "lottery_start_date": {
                    "description": "抽奖开始时间",
                    "type": "string"
                },
                "remark": {
                    "description": "演唱会备注信息",
                    "type": "string"
                },
                "review_status": {
                    "description": "审核状态（0: 未审核, 1: 审核通过, 2: 审核失败）",
                    "type": "integer"
                },
                "tickets": {
                    "description": "演唱会门票列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/commonreq.Ticket"
                    }
                }
            }
        },
        "commonreq.ConcertReViewReq": {
            "type": "object",
            "properties": {
                "concert_id": {
                    "description": "提交的演唱会id",
                    "type": "string"
                },
                "pass": {
                    "description": "是否通过",
                    "type": "boolean"
                }
            }
        },
        "commonreq.Ticket": {
            "type": "object",
            "properties": {
                "max_quantity_per_wallet": {
                    "description": "每个钱包最大可购买数量",
                    "type": "integer"
                },
                "num": {
                    "description": "门票数量",
                    "type": "integer"
                },
                "price": {
                    "description": "门票价格",
                    "type": "string"
                },
                "ticket_id": {
                    "description": "门票ID",
                    "type": "string"
                },
                "ticket_img": {
                    "description": "门票图片",
                    "type": "string"
                },
                "ticket_type": {
                    "description": "门票类型唯一键（对应抵押品合约和抽选合约）",
                    "type": "string"
                },
                "trade": {
                    "description": "是否可以二手交易",
                    "type": "boolean"
                },
                "type_name": {
                    "description": "门票类型名称",
                    "type": "string"
                }
            }
        },
        "commonresp.CommitResp": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "返回状态码",
                    "type": "integer"
                },
                "msg": {
                    "description": "返回消息",
                    "type": "string"
                },
                "reason": {
                    "description": "错误原因",
                    "type": "string"
                },
                "request_id": {
                    "description": "请求ID",
                    "type": "string"
                },
                "result": {
                    "description": "返回结果",
                    "type": "object",
                    "$ref": "#/definitions/commonresp.ResultData"
                },
                "status": {
                    "description": "返回状态",
                    "type": "string"
                }
            }
        },
        "commonresp.Concert": {
            "type": "object",
            "properties": {
                "concert_date": {
                    "description": "演唱会日期",
                    "type": "string"
                },
                "concert_id": {
                    "description": "演唱会ID",
                    "type": "string"
                },
                "concert_img": {
                    "description": "演唱会图片",
                    "type": "string"
                },
                "concert_name": {
                    "description": "演唱会名称",
                    "type": "string"
                },
                "concert_status": {
                    "description": "演唱会状态",
                    "type": "integer"
                },
                "status": {
                    "description": "演唱会状态",
                    "type": "string"
                },
                "ticket_types": {
                    "description": "门票类型列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/commonresp.TicketType"
                    }
                }
            }
        },
        "commonresp.ConcertListResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                },
                "request_id": {
                    "type": "string"
                },
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/commonresp.Concert"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "commonresp.ResultData": {
            "type": "object",
            "properties": {
                "concert_id": {
                    "description": "演唱会ID",
                    "type": "string"
                }
            }
        },
        "commonresp.TicketType": {
            "type": "object",
            "properties": {
                "max_quantity_per_wallet": {
                    "description": "单个钱包最大购买数量",
                    "type": "integer"
                },
                "price": {
                    "description": "价格",
                    "type": "string"
                },
                "ticket_type": {
                    "description": "门票种类唯一键",
                    "type": "string"
                },
                "type_name": {
                    "description": "类型名称",
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
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
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

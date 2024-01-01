// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplatelogistic = `{
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
        "/logistic": {
            "get": {
                "description": "Get a logistic based on given parameter",
                "tags": [
                    "Logistic"
                ],
                "summary": "Find a logistic data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "test example",
                        "name": "origin_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "test example",
                        "name": "destionation_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorzation(Bearer random_value)",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.FindOneResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new logistic data and return a message",
                "tags": [
                    "Logistic"
                ],
                "summary": "Create a new logistic data",
                "parameters": [
                    {
                        "description": "Create a new logistic data",
                        "name": "create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/pb.CreateLogisticRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorzation(Bearer random_value)",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.CreateLogisticResponse"
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
        }
    },
    "definitions": {
        "pb.CreateLogisticRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "destinationName": {
                    "type": "string"
                },
                "duration": {
                    "type": "string"
                },
                "logisticName": {
                    "type": "string"
                },
                "originName": {
                    "type": "string"
                }
            }
        },
        "pb.CreateLogisticResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "pb.FindOneData": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "destinationName": {
                    "type": "string"
                },
                "duration": {
                    "type": "string"
                },
                "logisticName": {
                    "type": "string"
                },
                "originName": {
                    "type": "string"
                }
            }
        },
        "pb.FindOneResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/pb.FindOneData"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfologistic holds exported Swagger Info so clients can modify it
var SwaggerInfologistic = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "logistic",
	SwaggerTemplate:  docTemplatelogistic,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfologistic.InstanceName(), SwaggerInfologistic)
}
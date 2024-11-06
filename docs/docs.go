// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Jaco Hanekom",
            "email": "jhanekom@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/sanitize": {
            "post": {
                "description": "Provides the ability to sanitize strings based on the stored values",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sanitize"
                ],
                "summary": "Sanitize",
                "parameters": [
                    {
                        "description": "Sanitize Request",
                        "name": "sanitize",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Sanitize"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Sanitize"
                        }
                    }
                }
            }
        },
        "/words": {
            "get": {
                "description": "Returns all the current words that will be used to sanitize text",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Sanitized Words",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.SanitizeWord"
                        }
                    }
                }
            },
            "put": {
                "description": "Provides the ability to add sanitized words",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Add Sanitized Words",
                "parameters": [
                    {
                        "description": "Add Sanitized Word",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.SanitizeWord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.SanitizeWord"
                        }
                    }
                }
            },
            "post": {
                "description": "Provides the ability to update an existing word(s)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Update Sanitized Words",
                "parameters": [
                    {
                        "description": "Update Sanitized Word",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.SanitizeWord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.SanitizeWord"
                        }
                    }
                }
            },
            "delete": {
                "description": "Provides the ability to add sanitized words",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Remove Sanitized Words",
                "parameters": [
                    {
                        "description": "Remove Sanitized Word",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.SanitizeWord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.SanitizeWord"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Sanitize": {
            "type": "object",
            "properties": {
                "sentences": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "controller.SanitizeWord": {
            "type": "object",
            "properties": {
                "words": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Sanitize Web Service",
	Description:      "This is a microservice to sanitize strings based on the stored words, and to manage the stored words",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
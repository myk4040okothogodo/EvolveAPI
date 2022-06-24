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
        "contact": {
            "name": "mike_okoth_ogodo"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users": {
            "get": {
                "description": "Get all users stored in the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "List all users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the page to be retrieved",
                        "name": "page_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/UserList"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "PutUser writes a user to the database\nTo write a new User, leave the id empty. To update an existing one, use the id/email/phonenumber of the user to be updated",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Add a user to the database",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "appCode": {
                    "description": "application-specific error code",
                    "type": "integer"
                },
                "errorText": {
                    "description": "application-level error message, for debugging",
                    "type": "string"
                },
                "statusText": {
                    "description": "user-level status message",
                    "type": "string"
                }
            }
        },
        "User": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "The address associated with this user",
                    "type": "string",
                    "example": "johns-street 27"
                },
                "birthday": {
                    "description": "The date of birth of this user",
                    "type": "string",
                    "example": "10-2-1997"
                },
                "createdat": {
                    "description": "the creation date and time of this user on the system",
                    "type": "string",
                    "example": "20-4-2930"
                },
                "email": {
                    "description": "The email address associated with this user",
                    "type": "string",
                    "example": "johndoe@gamil.com"
                },
                "id": {
                    "description": "The Unique id of this user",
                    "type": "integer",
                    "example": 1
                },
                "phonenumber": {
                    "description": "The phone number associated with this user",
                    "type": "string",
                    "example": "0717256998"
                }
            }
        },
        "UserList": {
            "type": "object",
            "properties": {
                "items": {
                    "description": "A list of users",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/User"
                    }
                },
                "next_page_id": {
                    "description": "The id to query the next page",
                    "type": "integer",
                    "example": 10
                },
                "pageSize": {
                    "description": "The page size of the query",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "example.com",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "USER API",
	Description:      "This API is a fullfillment of Evolves hiring task of building an API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

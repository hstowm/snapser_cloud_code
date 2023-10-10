// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/v2/byosnap-skybull/level-up-champion": {
            "post": {
                "description": "Level up champion",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "summary": "Level up champion",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userId when they login",
                        "name": "userId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "session token",
                        "name": "sessionToken",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "champion userID",
                        "name": "championUserId",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/byosnap-skybull/level-up-equipment": {
            "post": {
                "description": "Level up equipment",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "summary": "Level up equipment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userId when they login",
                        "name": "userId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "session token",
                        "name": "sessionToken",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id of equipment player want to level up",
                        "name": "UserEquipmentUID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/byosnap-skybull/re-roll-equipment-modifier": {
            "post": {
                "description": "Re-roll player equipment modifier",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "summary": "Re-Roll equipment modifier",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userId when they login",
                        "name": "userId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "session token",
                        "name": "sessionToken",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id of equipment player want to re-roll",
                        "name": "UserEquipmentUID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/byosnap-skybull/sell-equipment": {
            "post": {
                "description": "Sell player equipment",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "summary": "Sell equipment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userId when they login",
                        "name": "userId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "session token",
                        "name": "sessionToken",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id of equipment player want to sell",
                        "name": "UserEquipmentUID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/byosnap-skybull/win-campaign": {
            "post": {
                "description": "Call when win a campaign",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "summary": "Call when win a campaign",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userId when they login",
                        "name": "userId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "session token",
                        "name": "sessionToken",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "power stage player win",
                        "name": "stagePower",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the start they reach",
                        "name": "star",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/byosnap-skybull/win-pvp": {
            "post": {
                "description": "Call when player win a pvp game",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "summary": "Call when player win a pvp game",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userId when they login",
                        "name": "userId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "session token",
                        "name": "sessionToken",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "opponent snapser id ",
                        "name": "opponentID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Skybull Snapser CloudeCode",
	Description:      "Skybull snapser cloud code",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

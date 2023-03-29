// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplatetelemetry_backend = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Christian Jodra",
            "email": "c.jodra14@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/healthz": {
            "get": {
                "description": "This endpoint allow to see the serves status if 200 is received everything is working OK",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "status"
                ],
                "summary": "Server status",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/save": {
            "post": {
                "description": "Client Post the telmetry and it is saved on it's user on the database",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "telemetry"
                ],
                "summary": "This endpoint allows the client to post the telemetry",
                "parameters": [
                    {
                        "description": "Telemetry",
                        "name": "telemetry",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Telemetry"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user/{user-id}": {
            "get": {
                "description": "Client get the telemetry and it is saved on it's user on the database",
                "tags": [
                    "telemetry"
                ],
                "summary": "This endpoint allows the client to get user telemetries",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Accelerometer": {
            "type": "object",
            "properties": {
                "angle_x": {
                    "type": "number"
                },
                "angle_y": {
                    "type": "number"
                },
                "angle_z": {
                    "type": "number"
                },
                "g_force": {
                    "type": "number"
                }
            }
        },
        "models.GPS": {
            "type": "object",
            "properties": {
                "connected_satellites": {
                    "type": "integer"
                },
                "direction": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "speed": {
                    "type": "number"
                }
            }
        },
        "models.Telemetry": {
            "type": "object",
            "properties": {
                "accelerometer": {
                    "$ref": "#/definitions/models.Accelerometer"
                },
                "gps": {
                    "$ref": "#/definitions/models.GPS"
                },
                "timestamp": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfotelemetry_backend holds exported Swagger Info so clients can modify it
var SwaggerInfotelemetry_backend = &swag.Spec{
	Version:          "0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Telemetry IoT",
	Description:      "Telemetry REST API of IoT telemetry device,",
	InfoInstanceName: "telemetry_backend",
	SwaggerTemplate:  docTemplatetelemetry_backend,
}

func init() {
	swag.Register(SwaggerInfotelemetry_backend.InstanceName(), SwaggerInfotelemetry_backend)
}

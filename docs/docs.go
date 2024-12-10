// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Iulisloi Zacarias",
            "url": "https://github.com/izacarias"
        },
        "license": {
            "name": "BSD-3-Clause",
            "url": "https://forge.etsi.org/legal-matters"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/queries/zones": {
            "get": {
                "description": "The GET method is used to query the information about one or more specific zones or a list of zones.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "location"
                ],
                "summary": "Query the information about one or more specific zones or a list of zones.",
                "operationId": "zonesGET",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Zone ID",
                        "name": "zoneId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.ZoneResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/queries/zones/{zoneId}": {
            "get": {
                "responses": {
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "responses.ZoneResponse": {
            "type": "object",
            "properties": {
                "numberOfAccessPoints": {
                    "description": "The number of access points within the zone",
                    "type": "integer"
                },
                "numberOfUnserviceableAccessPoints": {
                    "description": "Number of inoperable access points within the zone.",
                    "type": "integer"
                },
                "numberOfUsers": {
                    "description": "The number of users currently on the access point.",
                    "type": "integer"
                },
                "resourceURL": {
                    "description": "Self referring URL",
                    "type": "string"
                },
                "zoneId": {
                    "description": "Identifier of zone",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "3.1.1",
	Host:             "localhost:8080",
	BasePath:         "/location/v3/",
	Schemes:          []string{"http"},
	Title:            "ETSI GS MEC 013 - Location API",
	Description:      "The ETSI MEC ISG MEC013 Location API described using OpenAPI",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

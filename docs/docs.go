// Code generated by swaggo/swag. DO NOT EDIT
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
        "/business": {
            "post": {
                "description": "Create a business",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "business"
                ],
                "summary": "Create a business",
                "parameters": [
                    {
                        "description": "Business",
                        "name": "business",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.BusinessRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/business/search": {
            "get": {
                "description": "Get businesses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "business"
                ],
                "summary": "Get businesses",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Location",
                        "name": "location",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Category",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort by",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Price",
                        "name": "price",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Open now",
                        "name": "open_now",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Longitude",
                        "name": "longitude",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Latitude",
                        "name": "latitude",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/business/{id}": {
            "put": {
                "description": "Edit a business",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "business"
                ],
                "summary": "Edit a business",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Business ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Business",
                        "name": "business",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.BusinessRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a business",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "business"
                ],
                "summary": "Delete a business",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Business ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/hello": {
            "get": {
                "description": "Hello",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hello"
                ],
                "summary": "Hello",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.BusinessRequest": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "string"
                },
                "categories_id": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "coordinates": {
                    "type": "object",
                    "properties": {
                        "latitude": {
                            "type": "number"
                        },
                        "longitude": {
                            "type": "number"
                        }
                    }
                },
                "display_phone": {
                    "type": "string"
                },
                "distance": {
                    "type": "number"
                },
                "image_url": {
                    "type": "string"
                },
                "is_closed": {
                    "type": "boolean"
                },
                "location": {
                    "type": "object",
                    "properties": {
                        "address1": {
                            "type": "string"
                        },
                        "address2": {
                            "type": "string"
                        },
                        "address3": {
                            "type": "string"
                        },
                        "city": {
                            "type": "string"
                        },
                        "country": {
                            "type": "string"
                        },
                        "displayAddress": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        },
                        "state": {
                            "type": "string"
                        },
                        "zip_code": {
                            "type": "string"
                        }
                    }
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "transactions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

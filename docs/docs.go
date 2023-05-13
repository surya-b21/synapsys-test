// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/cart": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all product in cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Get all cart",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number start from zero",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Size per page, default ` + "`" + `0` + "`" + `",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort by field, adding dash (` + "`" + `-` + "`" + `) at the beginning means descending and vice versa",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Select specific fields with comma separated",
                        "name": "fields",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)",
                        "name": "filters",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Cart"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add product to cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Post cart",
                "parameters": [
                    {
                        "description": "Payload Cart",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.CartAPI"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.CartAPI"
                            }
                        }
                    }
                }
            }
        },
        "/cart/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete product to cart",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Delete cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cart ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.CartAPI"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login new costumer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login new costumer",
                "parameters": [
                    {
                        "description": "Payload Login",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.Payload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Costumer"
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all list produk",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get all product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number start from zero",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Size per page, default ` + "`" + `0` + "`" + `",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort by field, adding dash (` + "`" + `-` + "`" + `) at the beginning means descending and vice versa",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Select specific fields with comma separated",
                        "name": "fields",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)",
                        "name": "filters",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Product"
                            }
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register new costumer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register new costumer",
                "parameters": [
                    {
                        "description": "Payload Register",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.Payload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Costumer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.Payload": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.Cart": {
            "type": "object",
            "properties": {
                "costumer": {
                    "$ref": "#/definitions/model.Costumer"
                },
                "costumer_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string",
                    "format": "date-time"
                },
                "id": {
                    "type": "string",
                    "format": "uuid"
                },
                "is_checkout": {
                    "type": "boolean"
                },
                "product": {
                    "$ref": "#/definitions/model.Product"
                },
                "product_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string",
                    "format": "date-time"
                }
            }
        },
        "model.CartAPI": {
            "type": "object",
            "properties": {
                "costumer_id": {
                    "type": "string"
                },
                "is_checkout": {
                    "type": "boolean"
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "model.Costumer": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "format": "date-time"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "format": "uuid"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string",
                    "format": "date-time"
                }
            }
        },
        "model.Product": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "format": "date-time"
                },
                "id": {
                    "type": "string",
                    "format": "uuid"
                },
                "name": {
                    "type": "string"
                },
                "product_category": {
                    "$ref": "#/definitions/model.ProductCategory"
                },
                "product_category_code": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string",
                    "format": "date-time"
                }
            }
        },
        "model.ProductCategory": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string",
                    "format": "date-time"
                },
                "id": {
                    "type": "string",
                    "format": "uuid"
                },
                "updated_at": {
                    "type": "string",
                    "format": "date-time"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Synapsis Test",
	Description:      "Documentation for synapsys test",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

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
        "/product": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Mendapatkan semua list produk",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get All product",
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
        }
    },
    "definitions": {
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

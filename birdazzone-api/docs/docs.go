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
        "contact": {},
        "license": {
            "name": "GNU AFFERO GENERAL PUBLIC LICENSE",
            "url": "https://www.gnu.org/licenses/agpl-3.0.en.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/hello": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "test"
                ],
                "summary": "Test your connection to birdazzone API",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/tvgames": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tvgames"
                ],
                "summary": "Get all TV games",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Game"
                            }
                        }
                    }
                }
            }
        },
        "/tvgames/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tvgames"
                ],
                "summary": "Get TV game",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID to search",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Game"
                        }
                    },
                    "404": {
                        "description": "game id not found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/tvgames/{id}/attempts": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tvgames"
                ],
                "summary": "Retrieve game's attempts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Game to query",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "date-time",
                        "description": "Starting instant of the time interval used to filter the tweets. If not specified, the beginning of the last game instance is used",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "date-time",
                        "description": "Ending instant of the time interval used to filter the tweets. Must be later than but in the same day of the starting instant. If not specified, the ending of the game happening during the starting instant is used",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "Length of the page to query",
                        "name": "pageLength",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Page-model_Tweet"
                        }
                    },
                    "400": {
                        "description": "integer parsing error (id)",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "404": {
                        "description": "game id not found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "(internal server error)",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/tvgames/{id}/attempts/stats": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tvgames"
                ],
                "summary": "Retrieve game's attempts' frequencies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Game to query",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "date-time",
                        "description": "Starting instant of the time interval used to filter the tweets. If not specified, the beginning of the last game instance is used",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "date-time",
                        "description": "Ending instant of the time interval used to filter the tweets. Must be later than but in the same day of the starting instant. If not specified, the ending of the game happening during the starting instant is used",
                        "name": "to",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "description": "A chart as a sequence of entries",
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ChartEntry"
                            }
                        }
                    },
                    "400": {
                        "description": "integer parsing error (id)",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "404": {
                        "description": "game id not found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/tvgames/{id}/results": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tvgames"
                ],
                "summary": "Retrieve game's number of successes and failures, grouped in time interval bins",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Game to query",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "date",
                        "description": "Starting date of the time interval used to filter the tweets. If not specified, the last game instance's date is used",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "date",
                        "description": "Ending date of the time interval used to filter the tweets. Cannot be earlier than the starting date. If not specified, the starting date is used",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Number of seconds for the duration of each time interval bin the retrieved tweets are to be grouped by",
                        "name": "each",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "A array of boolean charts comparing successes and failures in the game. Each boolean chart is labeled as the starting instant of its time interval bin",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.BooleanChart"
                            }
                        }
                    },
                    "400": {
                        "description": "each \u003c 1",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "404": {
                        "description": "game id not found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/tvgames/{id}/solution": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tvgames"
                ],
                "summary": "Retrieve game's solution",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Game to query",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "date",
                        "description": "Date to query; if not specified, last game instance is considered",
                        "name": "date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GameKey"
                        }
                    },
                    "400": {
                        "description": "error while parsing to date",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "404": {
                        "description": "game id not found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "(internal server error)",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.BooleanChart": {
            "description": "A pair made of one positive counter and one negative counter",
            "type": "object",
            "properties": {
                "negatives": {
                    "type": "integer",
                    "minimum": 0,
                    "example": 318
                },
                "positives": {
                    "type": "integer",
                    "minimum": 0,
                    "example": 209
                },
                "string": {
                    "type": "string",
                    "example": "Votes"
                }
            }
        },
        "model.ChartEntry": {
            "description": "A possible value inside a chart, alongside its absolute frequency",
            "type": "object",
            "properties": {
                "absolute_frequency": {
                    "type": "integer",
                    "minimum": 0,
                    "example": 34
                },
                "value": {
                    "type": "string",
                    "example": "parola"
                }
            }
        },
        "model.Error": {
            "description": "Object returned on failed requests",
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "model.Game": {
            "description": "A game which can be observed",
            "type": "object",
            "properties": {
                "hashtag": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "logo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.GameKey": {
            "description": "A possible solution for a Game",
            "type": "object",
            "properties": {
                "date": {
                    "type": "string",
                    "format": "date",
                    "example": "2022-11-17"
                },
                "key": {
                    "type": "string",
                    "example": "parola"
                }
            }
        },
        "model.Metrics": {
            "description": "Useful metrics describing a single Tweet",
            "type": "object",
            "properties": {
                "like_count": {
                    "type": "integer",
                    "minimum": 0,
                    "example": 122
                },
                "reply_count": {
                    "type": "integer",
                    "minimum": 0,
                    "example": 42
                },
                "retweet_count": {
                    "type": "integer",
                    "minimum": 0,
                    "example": 15
                }
            }
        },
        "model.Page-model_Tweet": {
            "type": "object",
            "properties": {
                "entries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Tweet"
                    }
                },
                "number_of_pages": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "model.Tweet": {
            "description": "A post published on Twitter",
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/model.User"
                },
                "created_at": {
                    "type": "string",
                    "format": "date-time"
                },
                "metrics": {
                    "$ref": "#/definitions/model.Metrics"
                },
                "text": {
                    "type": "string",
                    "example": "Hello, world!"
                }
            }
        },
        "model.User": {
            "description": "A Twitter user",
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Mario Rossi"
                },
                "profile_image_url": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "example": "mariorossi"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Birdazzone API",
	Description:      "API implementation for the Software Engineering (90106) project for the academic year 2022/23 at the University of Bologna.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

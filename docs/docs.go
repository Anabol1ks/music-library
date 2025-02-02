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
        "/songs": {
            "get": {
                "description": "Получает список песен по фильтрам: group, title, release_date, link; поддерживается пагинация",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Получение данных библиотеки с фильтрацией и пагинацией",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Группа",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Название песни",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Дата релиза",
                        "name": "release_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Ссылка",
                        "name": "link",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Номер страницы (начиная с 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Количество записей на страницу",
                        "name": "per_page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Song"
                            }
                        }
                    },
                    "400": {
                        "description": "Ошибка в запросе",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка получения данных",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Добавляет новую песню в базу данных",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Добавление новой песни",
                "parameters": [
                    {
                        "description": "Добавление песни",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/songs.InputSong"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Песня успешно добавлена",
                        "schema": {
                            "$ref": "#/definitions/response.CreateSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Ошибка в запросе",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка добавления песни",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/songs/{id}": {
            "delete": {
                "description": "Удаляет песню из базы данных по указанному идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Удаление песни",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Песня успешно удалена",
                        "schema": {
                            "$ref": "#/definitions/response.DeleteSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Ошибка в запросе",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Песня не найдена",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка удаления песни",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Частичное обновление данных песни по её ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Обновление данных песни",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные для обновления",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/songs.UpdateSongInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное обновление песни",
                        "schema": {
                            "$ref": "#/definitions/response.UpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Ошибка в запросе",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Песня не найдена",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка обновления данных песни",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/songs/{id}/text": {
            "get": {
                "description": "Возвращает куплеты текста песни по указанной странице и количеству куплетов на странице",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Получение текста песни с пагинацией по куплетам",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Номер страницы (начиная с 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 5,
                        "description": "Количество куплетов на странице",
                        "name": "per_page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Текст песни с пагинацией по куплетам",
                        "schema": {
                            "$ref": "#/definitions/response.TextResponse"
                        }
                    },
                    "400": {
                        "description": "Ошибка в запросе",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Песня не найдена",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Song": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "response.CreateSuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Песня успешно добавлена"
                }
            }
        },
        "response.DeleteSuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Песня успешно удалена"
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "текст ошибки"
                }
            }
        },
        "response.TextResponse": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "per_page": {
                    "type": "integer",
                    "example": 5
                },
                "song_id": {
                    "type": "integer",
                    "example": 1
                },
                "total_verses": {
                    "type": "integer",
                    "example": 2
                },
                "verses": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "[Куплет 1",
                        " Куплет 2]"
                    ]
                }
            }
        },
        "response.UpdateResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Данные песни успешно обновлены"
                },
                "song": {
                    "$ref": "#/definitions/models.Song"
                }
            }
        },
        "songs.InputSong": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                }
            }
        },
        "songs.UpdateSongInput": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
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
	Title:            "Онлайн библиотеки песен",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

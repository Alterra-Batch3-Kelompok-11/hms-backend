// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"github.com/swaggo/swag"
	"os"
)

const docTemplate = `{
    "openapi": "3.0.0",
    "info": {
        "title": "HMS-Backend",
        "version": "1.0.0"
    },
    "servers": [
        {
			"url": "http://localhost:8080",
            "url": "http://localhost:8080"
        }
    ],
    "tags": [
        {
            "name": "Auth"
        },
        {
            "name": "CRUD Patients"
        }
    ],
    "paths": {
        "/v1/login": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "username": "admin",
                                    "password": "admin123"
                                }
                            }
                        }
                    }
                },
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "type": "object",
                                    "example": {
                                        "status": 200,
                                        "message": "login success",
                                        "data": {
                                            "id": 3,
                                            "name": "admin",
                                            "username": "admin",
                                            "role_id": 1,
                                            "role": "",
                                            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjkxMzAxNTMsInJvbGVJZCI6MSwidXNlcklkIjozLCJ1c2VybmFtZSI6ImFkbWluIn0.InmvPtiTAvOpmKXg_jqNkgqafYg3XjdPcUf_JvdR7sI"
                                        }
                                    }
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Error response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "type": "object",
                                    "example": {
                                        "status": 400,
                                        "message": "record not found",
                                        "data": null
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "/v1/signup": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "summary": "Signup",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "name" : "Dr Fulan",
                                    "license_number" : "123123",
                                    "username" : "fulan123",
                                    "password" : "fulan123",
                                    "speciality_id" : 1,
                                    "role_id" : 2
                                }
                            }
                        }
                    }
                },
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "type": "object",
                                    "example": {
                                        "status": 200,
                                        "message": "signup success",
                                        "data": {
                                            "id": 5,
                                            "created_at": "2022-11-22T21:20:48.043+07:00",
                                            "updated_at": "2022-11-22T21:20:48.043+07:00",
                                            "deleted_at": null,
                                            "name": "Dr Fulan",
                                            "username": "fulan123",
                                            "role_id": 2,
                                            "role": "",
                                            "password": "$2a$10$89qf0IzQ3OPhuqV/AlIdG.QpLjZ3B/fFNxEBMTedL5pSjbJgnNDYC",
                                            "license_number": "123123"
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "/v1/create": {
            "post": {
                "tags": [
                    "CRUD Patients"
                ],
                "summary": "Create Patient",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "Nik": "12341234",
                                    "Name": "John Doe john",
                                    "Gender": 2,
                                    "Address": "Bekasi",
                                    "MaritalStatus": 2,
                                    "ReligionID": 2,
                                    "StatusID": 2
                                }
                            }
                        }
                    }
                },
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "content": {
                            "application/json": {}
                        }
                    }
                }
            }
        }
    }
}`

var b, err = os.ReadFile("swagger.json") // just pass the file name

var doc = string(b) // convert content to a 'string'

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

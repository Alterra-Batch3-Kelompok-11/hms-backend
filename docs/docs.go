// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"github.com/swaggo/swag"
)

func init() {
	//var b, _ = os.ReadFile("./docs/swagger.json") // just pass the file name
	//
	//var doc = string(b) // convert content to a 'string'
	var doc string = `
{
    "openapi": "3.0.0",
    "info": {
        "title": "HMS-Backend",
        "version": "1.0.0"
    },
    "servers": [
        {
            "url": "http://ec2-18-142-246-127.ap-southeast-1.compute.amazonaws.com"
        },
        {
            "url": "http://localhost:8080"
        }
    ],
    "tags": [
        {
            "name": "Roles"
        },
        {
            "name": "Auth"
        },
        {
            "name": "CRUD Patients"
        }
    ],
    "paths": {
        "/v1/roles": {
            "get": {
                "tags": [
                    "Roles"
                ],
                "summary": "Get All Roles",
                "requestBody": {
                    "content": {
                        "application/json": {
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
                                        "message": "success get data",
                                        "data": [
                                            {
                                                "id": 2,
                                                "created_at": "2022-11-23T08:51:16Z",
                                                "updated_at": "2022-11-23T08:51:16Z",
                                                "deleted_at": null,
                                                "name": "Doctor"
                                            },
                                            {
                                                "id": 3,
                                                "created_at": "2022-11-23T08:51:16Z",
                                                "updated_at": "2022-11-23T08:51:16Z",
                                                "deleted_at": null,
                                                "name": "Nurse"
                                            }
                                        ]
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "/v1/roles/1": {
            "get": {
                "tags": [
                    "Roles"
                ],
                "summary": "Get Role By Id",
                "requestBody": {
                    "content": {
                        "application/json": {
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
                                        "message": "success get data",
                                        "data": {
                                            "id": 1,
                                            "created_at": "2022-11-23T08:51:16Z",
                                            "updated_at": "2022-11-23T08:51:16Z",
                                            "deleted_at": null,
                                            "name": "Admin"
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
        },
        "/v1/specialities": {
            "get": {
                "tags": ["Specialities"],
                "summary": "Get Specialities",
                "requestBody": {
                    "content": {
                        "application/json": {}
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
                                        "message": "success get data",
                                        "data": [
                                            {
                                                "id": 1,
                                                "created_at": "2022-11-22T21:20:27+07:00",
                                                "updated_at": "2022-11-22T21:20:27+07:00",
                                                "deleted_at": null,
                                                "name": "Umum"
                                            },
                                            {
                                                "id": 2,
                                                "created_at": "2022-11-24T23:20:52.049+07:00",
                                                "updated_at": "2022-11-24T23:20:52.049+07:00",
                                                "deleted_at": null,
                                                "name": "Spesialis penyakit dalam"
                                            }
                                        ]
                                    }
                                }
                            }
                        }
                    }
                }
            },
            "post": {
                "tags": ["Specialities"],
                "summary": "Create Speciality",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "name": "Spesialis penyakit dalam"
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
                                        "message": "success create data",
                                        "data": {
                                            "id": 2,
                                            "created_at": "2022-11-24T23:20:52.049+07:00",
                                            "updated_at": "2022-11-24T23:20:52.049+07:00",
                                            "deleted_at": null,
                                            "name": "Spesialis penyakit dalam"
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "/v1/specialities/1": {
            "get": {
                "tags": ["Specialities"],
                "summary": "Get Speciality By Id",
                "requestBody": {
                    "content": {
                        "application/json": {}
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
                                        "message": "success get data",
                                        "data": {
                                            "id": 1,
                                            "created_at": "2022-11-24T23:20:52.049+07:00",
                                            "updated_at": "2022-11-24T23:20:52.049+07:00",
                                            "deleted_at": null,
                                            "name": "Umum"
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
            },
            "put": {
                "tags": ["Specialities"],
                "summary": "Update Speciality",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "name": "Umum 2"
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
                                        "message": "success update data",
                                        "data": {
                                            "id": 1,
                                            "created_at": "2022-11-24T23:20:52.049+07:00",
                                            "updated_at": "2022-11-24T23:20:52.049+07:00",
                                            "deleted_at": null,
                                            "name": "Umum 2"
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
            },
            "delete": {
                "tags": ["Specialities"],
                "summary": "Delete Speciality",
                "requestBody": {
                    "content": {
                        "application/json": {
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
                                        "message": "success delete data",
                                        "data": null
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
        }
    }
}
`

	// SwaggerInfo holds exported Swagger Info so clients can modify it
	var SwaggerInfo = &swag.Spec{
		Version:          "",
		Host:             "",
		BasePath:         "",
		Schemes:          []string{},
		Title:            "",
		Description:      "",
		InfoInstanceName: "swagger",
		//SwaggerTemplate:  docTemplate,
		SwaggerTemplate: doc,
	}
	//fmt.Println("b", b)
	//fmt.Println("doc", doc)
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

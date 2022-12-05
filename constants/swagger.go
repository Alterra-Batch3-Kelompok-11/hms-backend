package constants

const SwaggerDocTemplate = `
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
            "name": "Auth"
        },
        {
            "name": "Patients"
        },
        {
            "name": "Roles"
        },
        {
            "name": "Specialities"
        },
        {
            "name": "Religions"
        },
        {
            "name": "Doctors"
        },
        {
            "name": "Doctor Schedules"
        },
        {
            "name": "Nurse"
        },
        {
            "name": "Outpatient Sessions"
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
        "/v1/roles/{id}": {
            "get": {
                "tags": [
                    "Roles"
                ],
                "summary": "Get Role By Id",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of role",
                        "example": 1
                    }
                ],
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
                                    "username": "1029384756",
                                    "password": "john123"
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
                                            "user_id": 8,
                                            "name": "John Doe",
                                            "username": "1029384756",
                                            "role_id": 2,
                                            "role": "",
                                            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzAyMTY4NjQsInJvbGVJZCI6MiwidXNlcklkIjo4LCJ1c2VybmFtZSI6IjEwMjkzODQ3NTYifQ.mEvjUHcX8BtP0LhGwwzU9ZwXGmwdfwDLdQ5RwIf-j90",
                                            "doctor": {
                                                "id": 7,
                                                "created_at": "2022-11-30T13:05:43.168+07:00",
                                                "updated_at": "2022-11-30T13:05:43.168+07:00",
                                                "deleted_at": null,
                                                "name": "John Doe",
                                                "speciality_id": 1,
                                                "license_number": "1029384756",
                                                "speciality_name": "",
                                                "doctor_schedules": null
                                            },
                                            "nurse": null
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
                "parameters": [
                    {
                        "name": "Authorization",
                        "in": "header",
                        "description": "Authorization header token from login",
                        "example": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjkzMTAwOTcsInJvbGVJZCI6MSwidXNlcklkIjozLCJ1c2VybmFtZSI6ImFkbWluIn0.xfRdOVwqer4s9bKAxOX7LDE90tfnM-01ji6ae6HcLj4"
                    }
                ],
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
        "/v1/specialities/{id}": {
            "get": {
                "tags": ["Specialities"],
                "summary": "Get Speciality By Id",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of role",
                        "example": 1
                    }
                ],
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
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of role",
                        "example": 1
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "description": "Authorization header token from login",
                        "example": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjkzMTAwOTcsInJvbGVJZCI6MSwidXNlcklkIjozLCJ1c2VybmFtZSI6ImFkbWluIn0.xfRdOVwqer4s9bKAxOX7LDE90tfnM-01ji6ae6HcLj4"
                    }
                ],
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
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of role",
                        "example": 1
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "description": "Authorization header token from login",
                        "example": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjkzMTAwOTcsInJvbGVJZCI6MSwidXNlcklkIjozLCJ1c2VybmFtZSI6ImFkbWluIn0.xfRdOVwqer4s9bKAxOX7LDE90tfnM-01ji6ae6HcLj4"
                    }
                ],
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
        },
        "/v1/religions": {
            "get": {
                "tags": ["Religions"],
                "summary": "Get Religions",
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
                                                "name": "Islam"
                                            },
                                            {
                                                "id": 2,
                                                "created_at": "2022-11-24T23:20:52.049+07:00",
                                                "updated_at": "2022-11-24T23:20:52.049+07:00",
                                                "deleted_at": null,
                                                "name": "Protestan"
                                            },
                                            {
                                                "id": 3,
                                                "created_at": "2022-11-24T23:20:52.049+07:00",
                                                "updated_at": "2022-11-24T23:20:52.049+07:00",
                                                "deleted_at": null,
                                                "name": "Katolik"
                                            },
                                            {
                                                "id": 4,
                                                "created_at": "2022-11-24T23:20:52.049+07:00",
                                                "updated_at": "2022-11-24T23:20:52.049+07:00",
                                                "deleted_at": null,
                                                "name": "Hindu"
                                            },
                                            {
                                                "id": 5,
                                                "created_at": "2022-11-24T23:20:52.049+07:00",
                                                "updated_at": "2022-11-24T23:20:52.049+07:00",
                                                "deleted_at": null,
                                                "name": "Buddha"
                                            },
                                            {
                                                "id": 6,
                                                "created_at": "2022-11-24T23:20:52.049+07:00",
                                                "updated_at": "2022-11-24T23:20:52.049+07:00",
                                                "deleted_at": null,
                                                "name": "Konghucu"
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
        "/v1/religions/{id}": {
            "get": {
                "tags": ["Religions"],
                "summary": "Get Religion By Id",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of religion",
                        "example": 1
                    }
                ],
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
                                            "name": "Islam"
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
        "/v1/doctors": {
            "get": {
                "tags": [
                    "Doctors"
                ],
                "summary": "Get All",
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
                                                "id": 3,
                                                "created_at": "2022-11-23T12:32:38.222+07:00",
                                                "updated_at": "2022-11-23T12:32:38.222+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Mulan",
                                                "speciality_id": 1,
                                                "license_number": "123123",
                                                "speciality_name": "Umum",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 1,
                                                        "created_at": "2022-11-30T18:19:24.22+07:00",
                                                        "updated_at": "2022-11-30T18:19:24.22+07:00",
                                                        "deleted_at": null,
                                                        "doctor_id": 3,
                                                        "day_int": 0,
                                                        "day_string": "Minggu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 2,
                                                        "created_at": "2022-11-30T18:19:31.138+07:00",
                                                        "updated_at": "2022-11-30T18:19:31.138+07:00",
                                                        "deleted_at": null,
                                                        "doctor_id": 3,
                                                        "day_int": 1,
                                                        "day_string": "Senin",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 3,
                                                        "created_at": "2022-11-30T18:19:34.567+07:00",
                                                        "updated_at": "2022-11-30T18:19:34.567+07:00",
                                                        "deleted_at": null,
                                                        "doctor_id": 3,
                                                        "day_int": 2,
                                                        "day_string": "Selasa",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 4,
                                                        "created_at": "2022-11-30T18:58:30.652+07:00",
                                                        "updated_at": "2022-11-30T18:58:30.652+07:00",
                                                        "deleted_at": null,
                                                        "doctor_id": 3,
                                                        "day_int": 4,
                                                        "day_string": "Kamis",
                                                        "start_time": "08:00",
                                                        "end_time": "11:00"
                                                    }
                                                ]
                                            },
                                            {
                                                "id": 4,
                                                "created_at": "2022-11-25T13:35:19.236+07:00",
                                                "updated_at": "2022-11-25T13:35:19.236+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Fulan",
                                                "speciality_id": 1,
                                                "license_number": "1231232",
                                                "speciality_name": "Umum",
                                                "doctor_schedules": null
                                            },
                                            {
                                                "id": 5,
                                                "created_at": "2022-11-25T13:35:43.268+07:00",
                                                "updated_at": "2022-11-25T13:35:43.268+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Fulan",
                                                "speciality_id": 1,
                                                "license_number": "",
                                                "speciality_name": "Umum",
                                                "doctor_schedules": null
                                            },
                                            {
                                                "id": 6,
                                                "created_at": "2022-11-28T08:54:44.924+07:00",
                                                "updated_at": "2022-11-28T08:54:44.924+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Mulannn",
                                                "speciality_id": 1,
                                                "license_number": "12312333",
                                                "speciality_name": "Umum",
                                                "doctor_schedules": null
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
                "tags": [
                    "Doctors"
                ],
                "summary": "Create",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "name": "John Doe",
                                    "license_number": "1234567890",
                                    "password": "john123",
                                    "speciality_id": 1
                                }
                            }
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "Authorization",
                        "in": "header",
                        "description": "Authorization header token from login",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjkzMTAwOTcsInJvbGVJZCI6MSwidXNlcklkIjozLCJ1c2VybmFtZSI6ImFkbWluIn0.xfRdOVwqer4s9bKAxOX7LDE90tfnM-01ji6ae6HcLj4"
                    }
                ],
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
                                            "id": 8,
                                            "created_at": "2022-12-02T00:05:36.226+07:00",
                                            "updated_at": "2022-12-02T00:05:36.226+07:00",
                                            "deleted_at": null,
                                            "name": "John Doe",
                                            "speciality_id": 1,
                                            "license_number": "1234567890",
                                            "speciality_name": "Umum",
                                            "doctor_schedules": []
                                        }
                                    }
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "error response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "type": "object",
                                    "example": {
                                        "status": 400,
                                        "message": "license number already exist",
                                        "data": null
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "/v1/doctors/{id}": {
            "get": {
                "tags": [
                    "Doctors"
                ],
                "summary": "Get By Id",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of doctor",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 1
                    }
                ],
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
                                            "id": 8,
                                            "created_at": "2022-12-02T00:05:36.226+07:00",
                                            "updated_at": "2022-12-02T00:05:36.226+07:00",
                                            "deleted_at": null,
                                            "name": "John Doe",
                                            "speciality_id": 1,
                                            "license_number": "1234567890",
                                            "speciality_name": "Umum",
                                            "doctor_schedules": []
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
                "tags": [
                    "Doctors"
                ],
                "summary": "Update",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "name": "Fulan",
                                    "license_number": "780198012345",
                                    "password": "fulan123",
                                    "speciality_id": 2
                                }
                            }
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of doctor",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 1
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "description": "Authorization header token from login",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjkzMTAwOTcsInJvbGVJZCI6MSwidXNlcklkIjozLCJ1c2VybmFtZSI6ImFkbWluIn0.xfRdOVwqer4s9bKAxOX7LDE90tfnM-01ji6ae6HcLj4"
                    }
                ],
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
                                            "id": 0,
                                            "created_at": "0001-01-01T00:00:00Z",
                                            "updated_at": "0001-01-01T00:00:00Z",
                                            "deleted_at": null,
                                            "name": "John Doe",
                                            "speciality_id": 1,
                                            "license_number": "1234567891",
                                            "speciality_name": "Umum",
                                            "doctor_schedules": null
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
                "tags": [
                    "Doctors"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of doctor",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 1
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "description": "Authorization header token from login",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjkzMTAwOTcsInJvbGVJZCI6MSwidXNlcklkIjozLCJ1c2VybmFtZSI6ImFkbWluIn0.xfRdOVwqer4s9bKAxOX7LDE90tfnM-01ji6ae6HcLj4"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "content": {
                            "application/json": {}
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
        "/v1/doctors/license_number/{license_number}": {
            "get": {
                "tags": [
                    "Doctors"
                ],
                "summary": "Get By license number",
                "parameters": [
                    {
                        "name": "license_number",
                        "in": "path",
                        "description": "license number of doctor",
                        "schema": {
                            "type": "string"
                        },
                        "example": "1234567891"
                    }
                ],
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
                                            "created_at": "0001-01-01T00:00:00Z",
                                            "updated_at": "0001-01-01T00:00:00Z",
                                            "deleted_at": null,
                                            "name": "John Doe",
                                            "speciality_id": 1,
                                            "license_number": "1234567891",
                                            "speciality_name": "Umum",
                                            "doctor_schedules": null
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
        "/v1/doctors/speciality/{speciality_id}": {
            "get": {
                "tags": [
                    "Doctors"
                ],
                "summary": "Get By Speciality Id",
                "parameters": [
                    {
                        "name": "speciality_id",
                        "in": "path",
                        "description": "ID of speciality",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 1
                    }
                ],
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
                                                "id": 3,
                                                "created_at": "2022-11-23T12:32:38.222+07:00",
                                                "updated_at": "2022-11-23T12:32:38.222+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Mulan",
                                                "speciality_id": 1,
                                                "license_number": "123123",
                                                "speciality_name": "Umum",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 1,
                                                        "created_at": "2022-11-30T18:19:24.22+07:00",
                                                        "updated_at": "2022-11-30T18:19:24.22+07:00",
                                                        "deleted_at": null,
                                                        "doctor_id": 3,
                                                        "day_int": 0,
                                                        "day_string": "Minggu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 2,
                                                        "created_at": "2022-11-30T18:19:31.138+07:00",
                                                        "updated_at": "2022-11-30T18:19:31.138+07:00",
                                                        "deleted_at": null,
                                                        "doctor_id": 3,
                                                        "day_int": 1,
                                                        "day_string": "Senin",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 3,
                                                        "created_at": "2022-11-30T18:19:34.567+07:00",
                                                        "updated_at": "2022-11-30T18:19:34.567+07:00",
                                                        "deleted_at": null,
                                                        "doctor_id": 3,
                                                        "day_int": 2,
                                                        "day_string": "Selasa",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 4,
                                                        "created_at": "2022-11-30T18:58:30.652+07:00",
                                                        "updated_at": "2022-11-30T18:58:30.652+07:00",
                                                        "deleted_at": null,
                                                        "doctor_id": 3,
                                                        "day_int": 4,
                                                        "day_string": "Kamis",
                                                        "start_time": "08:00",
                                                        "end_time": "11:00"
                                                    }
                                                ]
                                            },
                                            {
                                                "id": 4,
                                                "created_at": "2022-11-25T13:35:19.236+07:00",
                                                "updated_at": "2022-11-25T13:35:19.236+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Fulan",
                                                "speciality_id": 1,
                                                "license_number": "1231232",
                                                "speciality_name": "Umum",
                                                "doctor_schedules": null
                                            },
                                            {
                                                "id": 5,
                                                "created_at": "2022-11-25T13:35:43.268+07:00",
                                                "updated_at": "2022-11-25T13:35:43.268+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Fulan",
                                                "speciality_id": 1,
                                                "license_number": "",
                                                "speciality_name": "Umum",
                                                "doctor_schedules": null
                                            },
                                            {
                                                "id": 6,
                                                "created_at": "2022-11-28T08:54:44.924+07:00",
                                                "updated_at": "2022-11-28T08:54:44.924+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Mulannn",
                                                "speciality_id": 1,
                                                "license_number": "12312333",
                                                "speciality_name": "Umum",
                                                "doctor_schedules": null
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
        "/v1/doctors/today": {
            "get": {
                "tags": [
                    "Doctors"
                ],
                "summary": "Today",
                "description": "Get today doctors",
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
                                                "id": 3,
                                                "created_at": "2022-11-23T12:32:38.222+07:00",
                                                "updated_at": "2022-11-23T12:32:38.222+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Mulan",
                                                "speciality_id": 1,
                                                "license_number": "123123",
                                                "speciality_name": "Umum",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 1,
                                                        "created_at": "2022-11-30T18:19:24.22+07:00",
                                                        "updated_at": "2022-11-30T18:19:24.22+07:00",
                                                        "deleted_at": null,
                                                        "doctor_id": 3,
                                                        "day_int": 0,
                                                        "day_string": "Minggu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 2,
                                                        "created_at": "2022-11-30T18:19:31.138+07:00",
                                                        "updated_at": "2022-11-30T18:19:31.138+07:00",
                                                        "deleted_at": null,
                                                        "doctor_id": 3,
                                                        "day_int": 1,
                                                        "day_string": "Senin",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 3,
                                                        "created_at": "2022-11-30T18:19:34.567+07:00",
                                                        "updated_at": "2022-11-30T18:19:34.567+07:00",
                                                        "deleted_at": null,
                                                        "doctor_id": 3,
                                                        "day_int": 2,
                                                        "day_string": "Selasa",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 4,
                                                        "created_at": "2022-11-30T18:58:30.652+07:00",
                                                        "updated_at": "2022-11-30T18:58:30.652+07:00",
                                                        "deleted_at": null,
                                                        "doctor_id": 3,
                                                        "day_int": 4,
                                                        "day_string": "Kamis",
                                                        "start_time": "08:00",
                                                        "end_time": "11:00"
                                                    }
                                                ]
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
        "/v1/doctor_schedules/{id}": {
            "get": {
                "tags": [
                    "Doctor Schedules"
                ],
                "summary": "Get By Id",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of doctor schedule",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 3
                    }
                ],
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
                                            "id": 3,
                                            "created_at": "2022-11-30T18:19:34.567+07:00",
                                            "updated_at": "2022-11-30T18:19:34.567+07:00",
                                            "deleted_at": null,
                                            "doctor_id": 3,
                                            "day_int": 2,
                                            "day_string": "Selasa",
                                            "start_time": "08:00",
                                            "end_time": "15:00"
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
                "tags": [
                    "Doctor Schedules"
                ],
                "summary": "Update",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "name": "Fulan",
                                    "license_number": "780198012345",
                                    "password": "fulan123",
                                    "speciality_id": 2
                                }
                            }
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of doctor schedule",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 3
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {{TOKEN}}"
                    }
                ],
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
                                            "id": 3,
                                            "created_at": "2022-11-30T18:19:34.567+07:00",
                                            "updated_at": "2022-11-30T18:19:34.567+07:00",
                                            "deleted_at": null,
                                            "doctor_id": 3,
                                            "day_int": 2,
                                            "day_string": "Selasa",
                                            "start_time": "08:00",
                                            "end_time": "15:00"
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Doctor Schedules"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of doctor schedule",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 3
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {{TOKEN}}"
                    }
                ],
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
                    }
                }
            }
        },
        "/v1/doctor_schedules/doctor/{doctor_id}": {
            "get": {
                "tags": [
                    "Doctor Schedules"
                ],
                "summary": "Get By Doctor Id",
                "parameters": [
                    {
                        "name": "doctor_id",
                        "in": "path",
                        "description": "ID of doctor",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 1
                    }
                ],
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
                                                "created_at": "2022-11-30T18:19:24.22+07:00",
                                                "updated_at": "2022-11-30T18:19:24.22+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "day_int": 0,
                                                "day_string": "Minggu",
                                                "start_time": "08:00",
                                                "end_time": "15:00"
                                            },
                                            {
                                                "id": 2,
                                                "created_at": "2022-11-30T18:19:31.138+07:00",
                                                "updated_at": "2022-11-30T18:19:31.138+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "day_int": 1,
                                                "day_string": "Senin",
                                                "start_time": "08:00",
                                                "end_time": "15:00"
                                            },
                                            {
                                                "id": 3,
                                                "created_at": "2022-11-30T18:19:34.567+07:00",
                                                "updated_at": "2022-11-30T18:19:34.567+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "day_int": 2,
                                                "day_string": "Selasa",
                                                "start_time": "08:00",
                                                "end_time": "15:00"
                                            },
                                            {
                                                "id": 4,
                                                "created_at": "2022-11-30T18:58:30.652+07:00",
                                                "updated_at": "2022-11-30T18:58:30.652+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "day_int": 4,
                                                "day_string": "Kamis",
                                                "start_time": "08:00",
                                                "end_time": "11:00"
                                            }
                                        ]
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
        "/v1/doctor_schedules/doctor/license_number/{license_number}": {
            "get": {
                "tags": [
                    "Doctor Schedules"
                ],
                "summary": "Get By License Number",
                "parameters": [
                    {
                        "name": "license_number",
                        "in": "path",
                        "description": "Doctor's License Number",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 123123
                    }
                ],
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
                                                "created_at": "2022-11-30T18:19:24.22+07:00",
                                                "updated_at": "2022-11-30T18:19:24.22+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "day_int": 0,
                                                "day_string": "Minggu",
                                                "start_time": "08:00",
                                                "end_time": "15:00"
                                            },
                                            {
                                                "id": 2,
                                                "created_at": "2022-11-30T18:19:31.138+07:00",
                                                "updated_at": "2022-11-30T18:19:31.138+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "day_int": 1,
                                                "day_string": "Senin",
                                                "start_time": "08:00",
                                                "end_time": "15:00"
                                            },
                                            {
                                                "id": 3,
                                                "created_at": "2022-11-30T18:19:34.567+07:00",
                                                "updated_at": "2022-11-30T18:19:34.567+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "day_int": 2,
                                                "day_string": "Selasa",
                                                "start_time": "08:00",
                                                "end_time": "15:00"
                                            },
                                            {
                                                "id": 4,
                                                "created_at": "2022-11-30T18:58:30.652+07:00",
                                                "updated_at": "2022-11-30T18:58:30.652+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "day_int": 4,
                                                "day_string": "Kamis",
                                                "start_time": "08:00",
                                                "end_time": "11:00"
                                            }
                                        ]
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
        "/v1/doctor_schedules": {
            "post": {
                "tags": [
                    "Doctor Schedules"
                ],
                "summary": "Create",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "doctor_id": 3,
                                    "day_int": 2,
                                    "start_time": "08:00",
                                    "end_time": "15:00"
                                }
                            }
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {{TOKEN}}"
                    }
                ],
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
                                            "id": 3,
                                            "created_at": "2022-11-30T18:19:34.567+07:00",
                                            "updated_at": "2022-11-30T18:19:34.567+07:00",
                                            "deleted_at": null,
                                            "doctor_id": 3,
                                            "day_int": 2,
                                            "day_string": "Selasa",
                                            "start_time": "08:00",
                                            "end_time": "15:00"
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
        "/v1/outpatient_sessions": {
            "get": {
                "tags": [
                    "Outpatient Sessions"
                ],
                "summary": "Get All",
                "parameters": [
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {{TOKEN}}"
                    }
                ],
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
                                                "created_at": "2022-12-04T22:54:54.411+07:00",
                                                "updated_at": "2022-12-04T22:54:54.411+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "patient_id": 1,
                                                "schedule": "2022-12-04T10:00:00+07:00",
                                                "complaint": "batuk berdahak",
                                                "is_approved": 0,
                                                "is_finish": false,
                                                "finished_at": "0001-01-01T00:00:00Z",
                                                "schedule_date": "2022-12-04",
                                                "schedule_time": "10:00",
                                                "patient": {
                                                    "nik": "12341234",
                                                    "name": "John tor",
                                                    "birth_date": "0001-01-01T00:00:00Z",
                                                    "gender": 1,
                                                    "phone": "0812121212",
                                                    "address": "Bekasi",
                                                    "marital_status": true,
                                                    "religion_id": 1
                                                },
                                                "doctor": {
                                                    "id": 3,
                                                    "created_at": "2022-11-23T12:32:38.222+07:00",
                                                    "updated_at": "2022-11-23T12:32:38.222+07:00",
                                                    "deleted_at": null,
                                                    "name": "Dr Mulan",
                                                    "speciality_id": 1,
                                                    "license_number": "123123",
                                                    "speciality_name": "Umum",
                                                    "doctor_schedules": null
                                                }
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
                "tags": [
                    "Outpatient Sessions"
                ],
                "summary": "Create",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "patient_id": 1,
                                    "doctor_id": 3,
                                    "complaint": "batuk berdahak",
                                    "schedule_date": "2022-12-04",
                                    "schedule_time": "10:00:00"
                                }
                            }
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {{TOKEN}}"
                    }
                ],
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
                                            "doctor_id": 3,
                                            "patient_id": 1,
                                            "schedule": "2022-12-04T10:00:00+07:00",
                                            "complaint": "batuk berdahak",
                                            "is_approved": 0,
                                            "is_finish": false,
                                            "finished_at": "0001-01-01T00:00:00Z",
                                            "schedule_date": "2022-12-04",
                                            "schedule_time": "10:00",
                                            "patient": {
                                                "nik": "12341234",
                                                "name": "John tor",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "gender": 1,
                                                "phone": "0812121212",
                                                "address": "Bekasi",
                                                "marital_status": true,
                                                "religion_id": 1
                                            },
                                            "doctor": {
                                                "id": 3,
                                                "created_at": "2022-11-23T12:32:38.222+07:00",
                                                "updated_at": "2022-11-23T12:32:38.222+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Mulan",
                                                "speciality_id": 1,
                                                "license_number": "123123",
                                                "speciality_name": "Umum",
                                                "doctor_schedules": null
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "/v1/outpatient_sessions/{id}": {
            "get": {
                "tags": [
                    "Outpatient Sessions"
                ],
                "summary": "Get By Id",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of outpatient session",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 2
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {{TOKEN}}"
                    }
                ],
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
                                            "id": 2,
                                            "created_at": "2022-12-04T22:54:54.411+07:00",
                                            "updated_at": "2022-12-04T22:54:54.411+07:00",
                                            "deleted_at": null,
                                            "doctor_id": 3,
                                            "patient_id": 1,
                                            "schedule": "2022-12-04T10:00:00+07:00",
                                            "complaint": "batuk berdahak",
                                            "is_approved": 0,
                                            "is_finish": false,
                                            "finished_at": "0001-01-01T00:00:00Z",
                                            "schedule_date": "2022-12-04",
                                            "schedule_time": "10:00",
                                            "patient": {
                                                "nik": "12341234",
                                                "name": "John tor",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "gender": 1,
                                                "phone": "0812121212",
                                                "address": "Bekasi",
                                                "marital_status": true,
                                                "religion_id": 1
                                            },
                                            "doctor": {
                                                "id": 3,
                                                "created_at": "2022-11-23T12:32:38.222+07:00",
                                                "updated_at": "2022-11-23T12:32:38.222+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Mulan",
                                                "speciality_id": 1,
                                                "license_number": "123123",
                                                "speciality_name": "Umum",
                                                "doctor_schedules": null
                                            }
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
                "tags": [
                    "Outpatient Sessions"
                ],
                "summary": "Update",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "patient_id": 1,
                                    "doctor_id": 3,
                                    "complaint": "batuk berdahak",
                                    "schedule_date": "2022-12-04",
                                    "schedule_time": "11:00:00"
                                }
                            }
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of outpatient sessions",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 2
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {{TOKEN}}"
                    }
                ],
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
                                            "id": 2,
                                            "created_at": "2022-12-04T22:54:54.411+07:00",
                                            "updated_at": "2022-12-04T22:54:54.411+07:00",
                                            "deleted_at": null,
                                            "doctor_id": 3,
                                            "patient_id": 1,
                                            "schedule": "2022-12-04T10:00:00+07:00",
                                            "complaint": "batuk berdahak",
                                            "is_approved": 0,
                                            "is_finish": false,
                                            "finished_at": "0001-01-01T00:00:00Z",
                                            "schedule_date": "2022-12-04",
                                            "schedule_time": "10:00",
                                            "patient": {
                                                "nik": "12341234",
                                                "name": "John tor",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "gender": 1,
                                                "phone": "0812121212",
                                                "address": "Bekasi",
                                                "marital_status": true,
                                                "religion_id": 1
                                            },
                                            "doctor": {
                                                "id": 3,
                                                "created_at": "2022-11-23T12:32:38.222+07:00",
                                                "updated_at": "2022-11-23T12:32:38.222+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Mulan",
                                                "speciality_id": 1,
                                                "license_number": "123123",
                                                "speciality_name": "Umum",
                                                "doctor_schedules": null
                                            }
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
                "tags": [
                    "Outpatient Sessions"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of outpatient sessions",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 2
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {{TOKEN}}"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "content": {
                            "application/json": {
                                "schema": {
                                    "type": "object",
                                    "example": {
                                        "status" : 200,
                                        "message" : "success delete data",
                                        "data" : null
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
        "/v1/outpatient_sessions/patient/{patient_id}": {
            "get": {
                "tags": [
                    "Outpatient Sessions"
                ],
                "summary": "Get By Patient Id",
                "parameters": [
                    {
                        "name": "patient_id",
                        "in": "path",
                        "description": "ID of patient",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 1
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {{TOKEN}}"
                    }
                ],
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
                                                "created_at": "2022-12-04T22:54:54.411+07:00",
                                                "updated_at": "2022-12-04T22:54:54.411+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "patient_id": 1,
                                                "schedule": "2022-12-04T10:00:00+07:00",
                                                "complaint": "batuk berdahak",
                                                "is_approved": 0,
                                                "is_finish": false,
                                                "finished_at": "0001-01-01T00:00:00Z",
                                                "schedule_date": "2022-12-04",
                                                "schedule_time": "10:00",
                                                "patient": {
                                                    "nik": "12341234",
                                                    "name": "John tor",
                                                    "birth_date": "0001-01-01T00:00:00Z",
                                                    "gender": 1,
                                                    "phone": "0812121212",
                                                    "address": "Bekasi",
                                                    "marital_status": true,
                                                    "religion_id": 1
                                                },
                                                "doctor": {
                                                    "id": 3,
                                                    "created_at": "2022-11-23T12:32:38.222+07:00",
                                                    "updated_at": "2022-11-23T12:32:38.222+07:00",
                                                    "deleted_at": null,
                                                    "name": "Dr Mulan",
                                                    "speciality_id": 1,
                                                    "license_number": "123123",
                                                    "speciality_name": "Umum",
                                                    "doctor_schedules": null
                                                }
                                            }
                                        ]
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
        "/v1/outpatient_sessions/doctor/{doctor_id}": {
            "get": {
                "tags": [
                    "Outpatient Sessions"
                ],
                "summary": "Get By Doctor Id",
                "parameters": [
                    {
                        "name": "doctor_id",
                        "in": "path",
                        "description": "ID of doctor",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 3
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {{TOKEN}}"
                    }
                ],
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
                                                "created_at": "2022-12-04T22:54:54.411+07:00",
                                                "updated_at": "2022-12-04T22:54:54.411+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "patient_id": 1,
                                                "schedule": "2022-12-04T10:00:00+07:00",
                                                "complaint": "batuk berdahak",
                                                "is_approved": 0,
                                                "is_finish": false,
                                                "finished_at": "0001-01-01T00:00:00Z",
                                                "schedule_date": "2022-12-04",
                                                "schedule_time": "10:00",
                                                "patient": {
                                                    "nik": "12341234",
                                                    "name": "John tor",
                                                    "birth_date": "0001-01-01T00:00:00Z",
                                                    "gender": 1,
                                                    "phone": "0812121212",
                                                    "address": "Bekasi",
                                                    "marital_status": true,
                                                    "religion_id": 1
                                                },
                                                "doctor": {
                                                    "id": 3,
                                                    "created_at": "2022-11-23T12:32:38.222+07:00",
                                                    "updated_at": "2022-11-23T12:32:38.222+07:00",
                                                    "deleted_at": null,
                                                    "name": "Dr Mulan",
                                                    "speciality_id": 1,
                                                    "license_number": "123123",
                                                    "speciality_name": "Umum",
                                                    "doctor_schedules": null
                                                }
                                            }
                                        ]
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
        "/v1/outpatient_sessions/doctor/{doctor_id}/unprocesseds": {
            "get": {
                "tags": [
                    "Outpatient Sessions"
                ],
                "summary": "Get Unprocesseds By Doctor Id",
                "parameters": [
                    {
                        "name": "doctor_id",
                        "in": "path",
                        "description": "ID of doctor",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 3
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {{TOKEN}}"
                    }
                ],
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
                                                "created_at": "2022-12-04T22:54:54.411+07:00",
                                                "updated_at": "2022-12-04T22:54:54.411+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "patient_id": 1,
                                                "schedule": "2022-12-04T10:00:00+07:00",
                                                "complaint": "batuk berdahak",
                                                "is_approved": 0,
                                                "is_finish": false,
                                                "finished_at": "0001-01-01T00:00:00Z",
                                                "schedule_date": "2022-12-04",
                                                "schedule_time": "10:00",
                                                "patient": {
                                                    "nik": "12341234",
                                                    "name": "John tor",
                                                    "birth_date": "0001-01-01T00:00:00Z",
                                                    "gender": 1,
                                                    "phone": "0812121212",
                                                    "address": "Bekasi",
                                                    "marital_status": true,
                                                    "religion_id": 1
                                                },
                                                "doctor": {
                                                    "id": 3,
                                                    "created_at": "2022-11-23T12:32:38.222+07:00",
                                                    "updated_at": "2022-11-23T12:32:38.222+07:00",
                                                    "deleted_at": null,
                                                    "name": "Dr Mulan",
                                                    "speciality_id": 1,
                                                    "license_number": "123123",
                                                    "speciality_name": "Umum",
                                                    "doctor_schedules": null
                                                }
                                            }
                                        ]
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
        "/v1/outpatient_sessions/doctor/{doctor_id}/processeds": {
            "get": {
                "tags": [
                    "Outpatient Sessions"
                ],
                "summary": "Get Processeds By Doctor Id",
                "parameters": [
                    {
                        "name": "doctor_id",
                        "in": "path",
                        "description": "ID of doctor",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 3
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {{TOKEN}}"
                    }
                ],
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
                                                "created_at": "2022-12-04T22:54:54.411+07:00",
                                                "updated_at": "2022-12-04T22:54:54.411+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "patient_id": 1,
                                                "schedule": "2022-12-04T10:00:00+07:00",
                                                "complaint": "batuk berdahak",
                                                "is_approved": 1,
                                                "is_finish": false,
                                                "finished_at": "0001-01-01T00:00:00Z",
                                                "schedule_date": "2022-12-04",
                                                "schedule_time": "10:00",
                                                "patient": {
                                                    "nik": "12341234",
                                                    "name": "John tor",
                                                    "birth_date": "0001-01-01T00:00:00Z",
                                                    "gender": 1,
                                                    "phone": "0812121212",
                                                    "address": "Bekasi",
                                                    "marital_status": true,
                                                    "religion_id": 1
                                                },
                                                "doctor": {
                                                    "id": 3,
                                                    "created_at": "2022-11-23T12:32:38.222+07:00",
                                                    "updated_at": "2022-11-23T12:32:38.222+07:00",
                                                    "deleted_at": null,
                                                    "name": "Dr Mulan",
                                                    "speciality_id": 1,
                                                    "license_number": "123123",
                                                    "speciality_name": "Umum",
                                                    "doctor_schedules": null
                                                }
                                            }
                                        ]
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
        "/v1/outpatient_sessions/{id}/approval": {
            "put": {
                "tags": [
                    "Outpatient Sessions"
                ],
                "summary": "Approval",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "is_approved": 1
                                }
                            }
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of outpatient session",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 2
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {{TOKEN}}"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "content": {
                            "application/json": {}
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

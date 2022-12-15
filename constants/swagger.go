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
            "name": "Dashboard"
        },
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
            "name": "Nurses"
        },
        {
            "name": "Outpatient Sessions"
        },
        {
            "name": "Patient Conditions"
        },
        {
            "name": "Histories"
        },
        {
            "name": "Notifications"
        }
    ],
    "components": {
        "securitySchemes": {
            "bearerAuth": {
                "type": "apiKey",
                "scheme": "bearer",
                "bearerFormat": "JWT",
                "name": "Authorization",
                "in": "header"
            }
        }
    },

    "paths": {
        "/v1/dashboard/web": {
            "get": {
                "tags": [
                    "Dashboard"
                ],
                "summary": "Data Dashboard For Web",
                "security": [
                    {
                        "bearerAuth" : []
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
                                            "total_doctors": 5,
                                            "total_nurses": 0,
                                            "total_patients": 1,
                                            "today_doctors": null,
                                            "today_outpatient_sessions": null,
                                            "patients": [
                                                {
                                                    "patient": {
                                                        "nik": "12341234",
                                                        "name": "John tor",
                                                        "birth_date": "0001-01-01T00:00:00Z",
                                                        "gender": 1,
                                                        "age": 2021,
                                                        "phone": "0812121212",
                                                        "address": "Bekasi",
                                                        "marital_status": true,
                                                        "religion_name": "Islam"
                                                    },
                                                    "doctor": {
                                                        "name": "Dr Mulan",
                                                        "license_number": "123123",
                                                        "speciality_name": "Umum"
                                                    },
                                                    "schedule": "2022-12-07T10:00:00+07:00",
                                                    "complaint": "sakit mata",
                                                    "is_approved": 0,
                                                    "is_finish": false,
                                                    "finished_at": "0001-01-01T00:00:00Z",
                                                    "schedule_date": "2022-12-07",
                                                    "schedule_time": "10:00"
                                                },
                                                {
                                                    "patient": {
                                                        "nik": "12341234",
                                                        "name": "John tor",
                                                        "birth_date": "0001-01-01T00:00:00Z",
                                                        "gender": 1,
                                                        "age": 2021,
                                                        "phone": "0812121212",
                                                        "address": "Bekasi",
                                                        "marital_status": true,
                                                        "religion_name": "Islam"
                                                    },
                                                    "doctor": {
                                                        "name": "Dr Mulan",
                                                        "license_number": "123123",
                                                        "speciality_name": "Umum"
                                                    },
                                                    "schedule": "2022-12-04T10:00:00+07:00",
                                                    "complaint": "batuk berdahak",
                                                    "is_approved": 0,
                                                    "is_finish": true,
                                                    "finished_at": "2022-12-09T18:48:21.559+07:00",
                                                    "schedule_date": "2022-12-04",
                                                    "schedule_time": "10:00"
                                                }
                                            ]
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "/v1/dashboard/mobile/doctor/{doctor_id}": {
            "get": {
                "tags": [
                    "Dashboard"
                ],
                "summary": "Data Dashboard For Mobile",
                "parameters": [
                    {
                        "name": "doctor_id",
                        "in": "path",
                        "description": "ID of doctor",
                        "example": 7
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {TOKEN}"
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
                                            "total_queue_today": 1,
                                            "total_finished_today": 0,
                                            "patients_today": [
                                                {
                                                    "name": "John tor",
                                                    "schedule_date": "2022-12-11",
                                                    "schedule_date_indo": "11 Desember 2022",
                                                    "schedule_time": "10:00",
                                                    "schedule": "2022-12-11T10:00:00+07:00",
                                                    "complaint": "sakit gigi"
                                                }
                                            ]
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
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
                                            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzAyMTY4NjQsInJvbGVJZCI6MiwidXNlcklkIjo4LCJ1c2VybmFtZSI6IjEwMjkzODQ3NTYifQ.mEvjUHcX8BtP0LhGwwzU9ZwXGmwdfwDLdQ5RwIf-j90",
                                            "doctor_id": 7,
                                            "nurse_id": null
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
        "/v1/auth/refresh": {
            "get": {
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh Token",
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
                "security": [
                    {
                        "bearerAuth" : []
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
                                        "message": "login success",
                                        "data": {
                                            "user_id": 8,
                                            "name": "John Doe",
                                            "username": "1029384756",
                                            "role_id": 2,
                                            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzAyMTY4NjQsInJvbGVJZCI6MiwidXNlcklkIjo4LCJ1c2VybmFtZSI6IjEwMjkzODQ3NTYifQ.mEvjUHcX8BtP0LhGwwzU9ZwXGmwdfwDLdQ5RwIf-j90",
                                            "doctor_id": 7,
                                            "nurse_id": null
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
                                        "message": "user not found",
                                        "data": null
                                    }
                                }
                            }
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
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 7,
                                                        "doctor_id": 3,
                                                        "date": "2022-12-14T15:35:30.4974884+07:00",
                                                        "date_indo": "14 Desember 2022",
                                                        "day_int": 3,
                                                        "day_string": "Rabu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 1,
                                                        "doctor_id": 3,
                                                        "date": "2022-12-19T15:35:30.4974884+07:00",
                                                        "date_indo": "19 Desember 2022",
                                                        "day_int": 1,
                                                        "day_string": "Senin",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 4,
                                                        "doctor_id": 3,
                                                        "date": "2022-12-20T15:35:30.4974884+07:00",
                                                        "date_indo": "20 Desember 2022",
                                                        "day_int": 2,
                                                        "day_string": "Selasa",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
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
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
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
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 8,
                                                        "doctor_id": 5,
                                                        "date": "2022-12-14T15:35:30.975379+07:00",
                                                        "date_indo": "14 Desember 2022",
                                                        "day_int": 3,
                                                        "day_string": "Rabu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 2,
                                                        "doctor_id": 5,
                                                        "date": "2022-12-19T15:35:30.975379+07:00",
                                                        "date_indo": "19 Desember 2022",
                                                        "day_int": 1,
                                                        "day_string": "Senin",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 5,
                                                        "doctor_id": 5,
                                                        "date": "2022-12-20T15:35:30.975379+07:00",
                                                        "date_indo": "20 Desember 2022",
                                                        "day_int": 2,
                                                        "day_string": "Selasa",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    }
                                                ]
                                            },
                                            {
                                                "id": 6,
                                                "created_at": "2022-11-28T08:54:44.924+07:00",
                                                "updated_at": "2022-11-28T08:54:44.924+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Mulannn",
                                                "speciality_id": 1,
                                                "license_number": "12312333",
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 9,
                                                        "doctor_id": 6,
                                                        "date": "2022-12-14T15:35:31.0179182+07:00",
                                                        "date_indo": "14 Desember 2022",
                                                        "day_int": 3,
                                                        "day_string": "Rabu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    }
                                                ]
                                            },
                                            {
                                                "id": 7,
                                                "created_at": "2022-11-30T13:05:43.168+07:00",
                                                "updated_at": "2022-11-30T13:05:43.168+07:00",
                                                "deleted_at": null,
                                                "name": "John Doe",
                                                "speciality_id": 1,
                                                "license_number": "1029384756",
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 12,
                                                        "doctor_id": 7,
                                                        "date": "2022-12-14T15:35:31.0377342+07:00",
                                                        "date_indo": "14 Desember 2022",
                                                        "day_int": 3,
                                                        "day_string": "Rabu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 13,
                                                        "doctor_id": 7,
                                                        "date": "2022-12-15T15:35:31.0377342+07:00",
                                                        "date_indo": "15 Desember 2022",
                                                        "day_int": 4,
                                                        "day_string": "Kamis",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 14,
                                                        "doctor_id": 7,
                                                        "date": "2022-12-16T15:35:31.0377342+07:00",
                                                        "date_indo": "16 Desember 2022",
                                                        "day_int": 5,
                                                        "day_string": "Jumat",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 10,
                                                        "doctor_id": 7,
                                                        "date": "2022-12-19T15:35:31.0377342+07:00",
                                                        "date_indo": "19 Desember 2022",
                                                        "day_int": 1,
                                                        "day_string": "Senin",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 11,
                                                        "doctor_id": 7,
                                                        "date": "2022-12-20T15:35:31.0377342+07:00",
                                                        "date_indo": "20 Desember 2022",
                                                        "day_int": 2,
                                                        "day_string": "Selasa",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    }
                                                ]
                                            },
                                            {
                                                "id": 8,
                                                "created_at": "2022-12-05T12:14:41.072+07:00",
                                                "updated_at": "2022-12-05T12:14:41.072+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Hulan",
                                                "speciality_id": 1,
                                                "license_number": "12345",
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": null
                                            },
                                            {
                                                "id": 9,
                                                "created_at": "2022-12-09T15:24:54.961+07:00",
                                                "updated_at": "2022-12-09T15:24:54.961+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Moris",
                                                "speciality_id": 1,
                                                "license_number": "121",
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": null
                                            },
                                            {
                                                "id": 10,
                                                "created_at": "2022-12-14T15:33:40.988+07:00",
                                                "updated_at": "2022-12-14T15:33:40.988+07:00",
                                                "deleted_at": null,
                                                "name": "Dr. Bones",
                                                "speciality_id": 1,
                                                "license_number": "1234515243",
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "1998-01-02T00:00:00+07:00",
                                                "birth_date_string": "1998-1-02",
                                                "birth_date_string_indo": "02 Januari 1998",
                                                "phone": "081234567891",
                                                "marital_status": true,
                                                "email": "bones@mail.com",
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
                                    "name" : "Dr. Bones",
                                    "license_number" : "1234515243",
                                    "password" : "bones123",
                                    "speciality_id" : 1,
                                    "profile_pic": "",
                                    "birth_date": "1998-01-02",
                                    "phone": "081234567891",
                                    "marital_status" : true,
                                    "email" : "bones@mail.com"
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
                        "example": 10
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
                                            "id": 10,
                                            "created_at": "2022-12-14T15:33:40.988+07:00",
                                            "updated_at": "2022-12-14T15:33:40.988+07:00",
                                            "deleted_at": null,
                                            "name": "Dr. Bones",
                                            "speciality_id": 1,
                                            "license_number": "1234515243",
                                            "speciality_name": "Umum 2",
                                            "profile_pic": "",
                                            "birth_date": "1998-01-02T00:00:00+07:00",
                                            "birth_date_string": "1998-1-02",
                                            "birth_date_string_indo": "02 Januari 1998",
                                            "phone": "081234567891",
                                            "marital_status": true,
                                            "email": "bones@mail.com",
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
                                    "name" : "Dr. Bones",
                                    "license_number" : "1234515243",
                                    "password" : "bones123",
                                    "speciality_id" : 1,
                                    "profile_pic": "",
                                    "birth_date": "1998-01-02",
                                    "phone": "081234567891",
                                    "marital_status" : true,
                                    "email" : "bones@mail.com"
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
                        "example": "1234515243"
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
                                            "id": 10,
                                            "created_at": "2022-12-14T15:33:40.988+07:00",
                                            "updated_at": "2022-12-14T15:33:40.988+07:00",
                                            "deleted_at": null,
                                            "name": "Dr. Bones",
                                            "speciality_id": 1,
                                            "license_number": "1234515243",
                                            "speciality_name": "Umum 2",
                                            "profile_pic": "",
                                            "birth_date": "1998-01-02T00:00:00+07:00",
                                            "birth_date_string": "1998-1-02",
                                            "birth_date_string_indo": "02 Januari 1998",
                                            "phone": "081234567891",
                                            "marital_status": true,
                                            "email": "bones@mail.com",
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
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 7,
                                                        "doctor_id": 3,
                                                        "date": "2022-12-14T15:47:55.9345302+07:00",
                                                        "date_indo": "14 Desember 2022",
                                                        "day_int": 3,
                                                        "day_string": "Rabu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 1,
                                                        "doctor_id": 3,
                                                        "date": "2022-12-19T15:47:55.9345302+07:00",
                                                        "date_indo": "19 Desember 2022",
                                                        "day_int": 1,
                                                        "day_string": "Senin",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 4,
                                                        "doctor_id": 3,
                                                        "date": "2022-12-20T15:47:55.9345302+07:00",
                                                        "date_indo": "20 Desember 2022",
                                                        "day_int": 2,
                                                        "day_string": "Selasa",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
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
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
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
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 8,
                                                        "doctor_id": 5,
                                                        "date": "2022-12-14T15:47:56.4168133+07:00",
                                                        "date_indo": "14 Desember 2022",
                                                        "day_int": 3,
                                                        "day_string": "Rabu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 2,
                                                        "doctor_id": 5,
                                                        "date": "2022-12-19T15:47:56.4168133+07:00",
                                                        "date_indo": "19 Desember 2022",
                                                        "day_int": 1,
                                                        "day_string": "Senin",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 5,
                                                        "doctor_id": 5,
                                                        "date": "2022-12-20T15:47:56.4168133+07:00",
                                                        "date_indo": "20 Desember 2022",
                                                        "day_int": 2,
                                                        "day_string": "Selasa",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    }
                                                ]
                                            },
                                            {
                                                "id": 6,
                                                "created_at": "2022-11-28T08:54:44.924+07:00",
                                                "updated_at": "2022-11-28T08:54:44.924+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Mulannn",
                                                "speciality_id": 1,
                                                "license_number": "12312333",
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 9,
                                                        "doctor_id": 6,
                                                        "date": "2022-12-14T15:47:56.4338012+07:00",
                                                        "date_indo": "14 Desember 2022",
                                                        "day_int": 3,
                                                        "day_string": "Rabu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    }
                                                ]
                                            },
                                            {
                                                "id": 7,
                                                "created_at": "2022-11-30T13:05:43.168+07:00",
                                                "updated_at": "2022-11-30T13:05:43.168+07:00",
                                                "deleted_at": null,
                                                "name": "John Doe",
                                                "speciality_id": 1,
                                                "license_number": "1029384756",
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 12,
                                                        "doctor_id": 7,
                                                        "date": "2022-12-14T15:47:56.4595375+07:00",
                                                        "date_indo": "14 Desember 2022",
                                                        "day_int": 3,
                                                        "day_string": "Rabu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 13,
                                                        "doctor_id": 7,
                                                        "date": "2022-12-15T15:47:56.4595375+07:00",
                                                        "date_indo": "15 Desember 2022",
                                                        "day_int": 4,
                                                        "day_string": "Kamis",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 14,
                                                        "doctor_id": 7,
                                                        "date": "2022-12-16T15:47:56.4595375+07:00",
                                                        "date_indo": "16 Desember 2022",
                                                        "day_int": 5,
                                                        "day_string": "Jumat",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 10,
                                                        "doctor_id": 7,
                                                        "date": "2022-12-19T15:47:56.4595375+07:00",
                                                        "date_indo": "19 Desember 2022",
                                                        "day_int": 1,
                                                        "day_string": "Senin",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    },
                                                    {
                                                        "id": 11,
                                                        "doctor_id": 7,
                                                        "date": "2022-12-20T15:47:56.4595375+07:00",
                                                        "date_indo": "20 Desember 2022",
                                                        "day_int": 2,
                                                        "day_string": "Selasa",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    }
                                                ]
                                            },
                                            {
                                                "id": 8,
                                                "created_at": "2022-12-05T12:14:41.072+07:00",
                                                "updated_at": "2022-12-05T12:14:41.072+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Hulan",
                                                "speciality_id": 1,
                                                "license_number": "12345",
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": null
                                            },
                                            {
                                                "id": 9,
                                                "created_at": "2022-12-09T15:24:54.961+07:00",
                                                "updated_at": "2022-12-09T15:24:54.961+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Moris",
                                                "speciality_id": 1,
                                                "license_number": "121",
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": null
                                            },
                                            {
                                                "id": 10,
                                                "created_at": "2022-12-14T15:33:40.988+07:00",
                                                "updated_at": "2022-12-14T15:33:40.988+07:00",
                                                "deleted_at": null,
                                                "name": "Dr. Bones",
                                                "speciality_id": 1,
                                                "license_number": "1234515243",
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "1998-01-02T00:00:00+07:00",
                                                "birth_date_string": "1998-1-02",
                                                "birth_date_string_indo": "02 Januari 1998",
                                                "phone": "081234567891",
                                                "marital_status": true,
                                                "email": "bones@mail.com",
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
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 7,
                                                        "doctor_id": 3,
                                                        "date": "2022-12-14T15:48:15.913605+07:00",
                                                        "date_indo": "14 Desember 2022",
                                                        "day_int": 3,
                                                        "day_string": "Rabu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    }
                                                ]
                                            },
                                            {
                                                "id": 5,
                                                "created_at": "2022-11-25T13:35:43.268+07:00",
                                                "updated_at": "2022-11-25T13:35:43.268+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Fulan",
                                                "speciality_id": 1,
                                                "license_number": "",
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 8,
                                                        "doctor_id": 5,
                                                        "date": "2022-12-14T15:48:15.9305719+07:00",
                                                        "date_indo": "14 Desember 2022",
                                                        "day_int": 3,
                                                        "day_string": "Rabu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    }
                                                ]
                                            },
                                            {
                                                "id": 6,
                                                "created_at": "2022-11-28T08:54:44.924+07:00",
                                                "updated_at": "2022-11-28T08:54:44.924+07:00",
                                                "deleted_at": null,
                                                "name": "Dr Mulannn",
                                                "speciality_id": 1,
                                                "license_number": "12312333",
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 9,
                                                        "doctor_id": 6,
                                                        "date": "2022-12-14T15:48:15.9355137+07:00",
                                                        "date_indo": "14 Desember 2022",
                                                        "day_int": 3,
                                                        "day_string": "Rabu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
                                                    }
                                                ]
                                            },
                                            {
                                                "id": 7,
                                                "created_at": "2022-11-30T13:05:43.168+07:00",
                                                "updated_at": "2022-11-30T13:05:43.168+07:00",
                                                "deleted_at": null,
                                                "name": "John Doe",
                                                "speciality_id": 1,
                                                "license_number": "1029384756",
                                                "speciality_name": "Umum 2",
                                                "profile_pic": "",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "birth_date_string": "1-1-01",
                                                "birth_date_string_indo": "01 Januari 1",
                                                "phone": "",
                                                "marital_status": false,
                                                "email": "",
                                                "doctor_schedules": [
                                                    {
                                                        "id": 12,
                                                        "doctor_id": 7,
                                                        "date": "2022-12-14T15:48:15.9427151+07:00",
                                                        "date_indo": "14 Desember 2022",
                                                        "day_int": 3,
                                                        "day_string": "Rabu",
                                                        "start_time": "08:00",
                                                        "end_time": "15:00"
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
                        "example": "Bearer {TOKEN}"
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
                        "example": "Bearer {TOKEN}"
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
                        "example": "Bearer {TOKEN}"
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
                        "example": "Bearer {TOKEN}"
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
                                                "schedule_date_indo": "04 Desember 2022",
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
                                                    "profile_pic": "",
                                                    "doctor_schedules": null
                                                }
                                            },
                                            {
                                                "id": 3,
                                                "created_at": "2022-12-07T18:23:32.412+07:00",
                                                "updated_at": "2022-12-07T18:23:32.412+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "patient_id": 1,
                                                "schedule": "2022-12-07T10:00:00+07:00",
                                                "complaint": "sakit mata",
                                                "is_approved": 0,
                                                "is_finish": false,
                                                "finished_at": "0001-01-01T00:00:00Z",
                                                "schedule_date": "2022-12-07",
                                                "schedule_date_indo": "07 Desember 2022",
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
                                                    "profile_pic": "",
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
                        "example": "Bearer {TOKEN}"
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
                        "example": "Bearer {TOKEN}"
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
                                            "schedule_date_indo": "04 Desember 2022",
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
                                                "profile_pic": "",
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
                        "example": "Bearer {TOKEN}"
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
                        "example": "Bearer {TOKEN}"
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
                        "example": "Bearer {TOKEN}"
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
                                                "schedule_date_indo": "04 Desember 2022",
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
                                                    "profile_pic": "",
                                                    "doctor_schedules": null
                                                }
                                            },
                                            {
                                                "id": 3,
                                                "created_at": "2022-12-07T18:23:32.412+07:00",
                                                "updated_at": "2022-12-07T18:23:32.412+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "patient_id": 1,
                                                "schedule": "2022-12-07T10:00:00+07:00",
                                                "complaint": "sakit mata",
                                                "is_approved": 0,
                                                "is_finish": false,
                                                "finished_at": "0001-01-01T00:00:00Z",
                                                "schedule_date": "2022-12-07",
                                                "schedule_date_indo": "07 Desember 2022",
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
                                                    "profile_pic": "",
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
                        "example": "Bearer {TOKEN}"
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
                                                "schedule_date_indo": "04 Desember 2022",
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
                                                    "profile_pic": "",
                                                    "doctor_schedules": null
                                                }
                                            },
                                            {
                                                "id": 3,
                                                "created_at": "2022-12-07T18:23:32.412+07:00",
                                                "updated_at": "2022-12-07T18:23:32.412+07:00",
                                                "deleted_at": null,
                                                "doctor_id": 3,
                                                "patient_id": 1,
                                                "schedule": "2022-12-07T10:00:00+07:00",
                                                "complaint": "sakit mata",
                                                "is_approved": 0,
                                                "is_finish": false,
                                                "finished_at": "0001-01-01T00:00:00Z",
                                                "schedule_date": "2022-12-07",
                                                "schedule_date_indo": "07 Desember 2022",
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
                                                    "profile_pic": "",
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
                        "example": "Bearer {TOKEN}"
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
                        "example": "Bearer {TOKEN}"
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
        "/v1/outpatient_sessions/doctor/{doctor_id}/approveds": {
            "get": {
                "tags": [
                    "Outpatient Sessions"
                ],
                "summary": "Get Approved Outpatient Sessions By Doctor Id",
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
                        "example": "Bearer {TOKEN}"
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
        "/v1/outpatient_sessions/doctor/{doctor_id}/rejecteds": {
            "get": {
                "tags": [
                    "Outpatient Sessions"
                ],
                "summary": "Get Rejected Outpatient Sessions By Doctor Id",
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
                        "example": "Bearer {TOKEN}"
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
                                                "is_approved": 2,
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
                        "example": "Bearer {TOKEN}"
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
        "/v1/patients": {
            "get": {
                "tags": ["Patients"],
                "summary": "Get Patients",
                "requestBody": {
                    "content": {
                        "application/json" : {}
                    }
                    
                },
                "responses": {
                    "200": {
                        "description" : "Successful response",
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
                                                "created_at": "2022-12-05T11:12:31.321+07:00",
                                                "updated_at": "2022-12-05T11:12:31.321+07:00",
                                                "deleted_at": null,
                                                "nik": "12341234",
                                                "name": "John tor",
                                                "birth_date": "1992-01-01T00:00:00+07:00",
                                                "gender": 1,
                                                "address": "Bekasi",
                                                "phone": "0812121212",
                                                "marital_status": true,
                                                "religion_id": 1
                                            },
                                            {
                                                "id": 2,
                                                "created_at": "2022-12-07T13:06:59.487+07:00",
                                                "updated_at": "2022-12-07T13:06:59.487+07:00",
                                                "deleted_at": null,
                                                "nik": "1324354657",
                                                "name": "Abdil Tegar Arifin",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "gender": 1,
                                                "address": "Jember",
                                                "phone": "081234567890",
                                                "marital_status": true,
                                                "religion_id": 1
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
                "tags" : ["Patients"],
                "summary": "Create Patient",
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
                                    "nik": "1234509882",
                                    "name": "Ahmad Fulan",
                                    "gender": 1,
                                    "address": "Jakarta",
                                    "phone": "0812121213",
                                    "marital_status": false,
                                    "birth_date": "1997-01-02",
                                    "religion_id": 1
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
                                    "example" : {
                                        "status": 200,
                                        "message": "success create data",
                                        "data": {
                                            "id": 3,
                                            "created_at": "2022-12-15T12:23:20.437+07:00",
                                            "updated_at": "2022-12-15T12:23:20.437+07:00",
                                            "deleted_at": null,
                                            "nik": "1234509882",
                                            "name": "Ahmad Fulan",
                                            "birth_date": "2000-01-01T00:00:00Z",
                                            "gender": 1,
                                            "address": "Jakarta",
                                            "phone": "0812121213",
                                            "marital_status": false,
                                            "religion_id": 1
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        },
        "/v1/patient_conditions": {
            "post": {
                "tags": [
                    "Patient Conditions"
                ],
                "summary": "Insert Patient Condition",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "outpatient_session_id": 2,
                                    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque convallis, arcu id mattis lacinia, tellus nunc sagittis risus, ut egestas sem eros a neque. Ut efficitur non neque dictum vehicula.",
                                    "medicine": "paracetamol",
                                    "allergy": "debu"
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
                                            "treatment_id": 1,
                                            "outpatient_session_id": 2,
                                            "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque convallis, arcu id mattis lacinia, tellus nunc sagittis risus, ut egestas sem eros a neque. Ut efficitur non neque dictum vehicula.",
                                            "medicine": "paracetamol",
                                            "allergy": "debu",
                                            "is_finish": true,
                                            "finished_at": "2022-12-09T18:48:21.559971+07:00",
                                            "finished_at_indo": "09 Desember 2022"
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
                                        "message": "this outpatient session is not approved",
                                        "data": null
                                    }
                                }
                            }
                        }
                    }
                }
            }
         },
        "/v1/patient_conditions/{treatment_id}": {
            "get": {
                "tags": [
                    "Patient Conditions"
                ],
                "summary": "Get Patient Condition By Treatment Id",
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
                                        "message": "success get data",
                                        "data": {
                                            "treatment_id": 1,
                                            "patient": {
                                                "nik": "1324354657",
                                                "name": "Abdil Tegar Arifin",
                                                "birth_date": "0001-01-01T00:00:00Z",
                                                "gender": 1,
                                                "age": 2021,
                                                "phone": "081234567890",
                                                "address": "Jember",
                                                "marital_status": true,
                                                "religion_name": "11 Desember 2022"
                                            },
                                            "doctor": {
                                                "name": "John Doe",
                                                "license_number": "1029384756",
                                                "speciality_name": "Umum 2"
                                            },
                                            "schedule": "2022-12-07T03:00:00+07:00",
                                            "schedule_date": "2022-12-07",
                                            "schedule_date_indo": "07 Desember 2022",
                                            "schedule_time": "03:00",
                                            "complaint": "batuk berdahak",
                                            "is_approved": 1,
                                            "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque convallis, arcu id mattis lacinia, tellus nunc sagittis risus, ut egestas sem eros a neque. Ut efficitur non neque dictum vehicula.",
                                            "medicine": "paracetamol",
                                            "allergy": "debu",
                                            "is_finish": true,
                                            "finished_at": "2022-12-11T16:12:37.25+07:00",
                                            "finished_at_indo": ""
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
        "/v1/patient_conditions/patient/{patient_id}": {
            "get": {
                "tags": [
                    "Patient Conditions"
                ],
                "summary": "Get Patient Condition By Patient Id",
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
                                        "message": "success get data",
                                        "data": [
                                            {
                                                "treatment_id": 1,
                                                "patient": {
                                                    "nik": "1324354657",
                                                    "name": "Abdil Tegar Arifin",
                                                    "birth_date": "0001-01-01T00:00:00Z",
                                                    "gender": 1,
                                                    "age": 2021,
                                                    "phone": "081234567890",
                                                    "address": "Jember",
                                                    "marital_status": true,
                                                    "religion_name": "11 Desember 2022"
                                                },
                                                "doctor": {
                                                    "name": "John Doe",
                                                    "license_number": "1029384756",
                                                    "speciality_name": "Umum 2"
                                                },
                                                "schedule": "2022-12-07T03:00:00+07:00",
                                                "schedule_date": "2022-12-07",
                                                "schedule_date_indo": "07 Desember 2022",
                                                "schedule_time": "03:00",
                                                "complaint": "batuk berdahak",
                                                "is_approved": 1,
                                                "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque convallis, arcu id mattis lacinia, tellus nunc sagittis risus, ut egestas sem eros a neque. Ut efficitur non neque dictum vehicula.",
                                                "medicine": "paracetamol",
                                                "allergy": "debu",
                                                "is_finish": true,
                                                "finished_at": "2022-12-11T16:12:37.25+07:00",
                                                "finished_at_indo": ""
                                            }
                                        ]
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
        "/v1/histories/doctor/{doctor_id}/outpatient_sessions": {
            "get": {
                "tags": [
                    "Histories"
                ],
                "summary": "Get Outpatient Session Histories",
                "parameters": [
                    {
                        "name": "doctor_id",
                        "in": "path",
                        "description": "ID of doctor",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 7
                    },
                    {
                        "name": "Authorization",
                        "in": "header",
                        "schema": {
                            "type": "string"
                        },
                        "example": "Bearer {TOKEN}"
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
                                                "patient_name": "Abdil Tegar Arifin",
                                                "schedule": "2022-12-07T03:00:00+07:00",
                                                "schedule_date": "2022-12-07",
                                                "schedule_date_indo": "07 Desember 2022",
                                                "schedule_time": "03:00",
                                                "status": "Selesai"
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
        "/v1/histories/doctor/{doctor_id}/approvals": {
            "get": {
                "tags": [
                    "Histories"
                ],
                "summary": "Get Approval Histories",
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
                        "example": "Bearer {TOKEN}"
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
                                                "patient_name": "John tor",
                                                "schedule": "2022-12-06T03:00:00+07:00",
                                                "schedule_date": "2022-12-06",
                                                "schedule_date_indo": "06 Desember 2022",
                                                "schedule_time": "03:00",
                                                "status": "Kunjungan Ditolak"
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
        "/v1/patients/{id}" : {
            "get": {
                "tags": [
                    "Patients"
                ],
                "summary": "Get Patient by Id",
                "parameters": [
                {
                    "name" : "id",
                    "in" : "path",
                    "description" : "ID of patient",
                    "example": 2
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
                                            "id": 2,
                                            "created_at": "2022-12-07T13:06:59.487+07:00",
                                            "updated_at": "2022-12-07T13:06:59.487+07:00",
                                            "deleted_at": null,
                                            "nik": "1324354657",
                                            "name": "Abdil Tegar Arifin",
                                            "birth_date": "0001-01-01T00:00:00Z",
                                            "gender": 1,
                                            "address": "Jember",
                                            "phone": "081234567890",
                                            "marital_status": true,
                                            "religion_id": 1
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
                                        "message": "record not found",
                                        "data": null
                                    }
                                }
                            }
                        }
                    }
                }
            },
            "put" : {
                "tags" : ["Patients"],
                "summary" : "Update Patient",
                "parameters" : [
                    {
                        "name" : "id",
                        "in" : "path",
                        "description" : "ID of patient",
                        "example": 3
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
                                    "nik": "1234509882",
                                    "name": "Ahmad Fulan",
                                    "gender": 1,
                                    "address": "Jakarta",
                                    "phone": "0812121213",
                                    "marital_status": false,
                                    "birth_date": "1997-01-02",
                                    "religion_id": 1
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
                                            "id": 3,
                                            "created_at": "2022-12-15T12:23:20.437+07:00",
                                            "updated_at": "2022-12-15T12:23:20.437+07:00",
                                            "deleted_at": null,
                                            "nik": "1234509882",
                                            "name": "Ahmad Fulan",
                                            "birth_date": "2000-01-01T00:00:00Z",
                                            "gender": 1,
                                            "address": "Jakarta",
                                            "phone": "0812121213",
                                            "marital_status": false,
                                            "religion_id": 1
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
                "tags" : ["Patients"],
                "summary" : "Delete Patient",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of patient",
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
                                        "messages": "success delete data",
                                        "data": null
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
        "/v1/nurses": {
            "get": {
                "tags": [
                    "Nurses"
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
                                                "id": 1,
                                                "user_id": 15,
                                                "doctor_id": 7,
                                                "name": "Ns. Melati",
                                                "speciality_id": 1,
                                                "license_number": "1234567893",
                                                "speciality_name": "Umum 2",
                                                "birth_date": "1997-01-01T17:00:00+07:00",
                                                "birth_date_string": "1997-1-01",
                                                "birth_date_string_indo": "01 Januari 1997",
                                                "profile_pic": "",
                                                "phone": "081234567892",
                                                "marital_status": true,
                                                "email": "melati@mail.com",
                                                "nira": "1928374650",
                                                "sip": "",
                                                "created_at": "2022-12-14T09:50:10.05+07:00",
                                                "updated_at": "2022-12-14T09:50:10.05+07:00",
                                                "deleted_at": null
                                            },
                                            {
                                                "id": 2,
                                                "user_id": 16,
                                                "doctor_id": 7,
                                                "name": "Ns. Mawar",
                                                "speciality_id": 1,
                                                "license_number": "1234567894",
                                                "speciality_name": "Umum 2",
                                                "birth_date": "1997-01-02T00:00:00+07:00",
                                                "birth_date_string": "1997-1-02",
                                                "birth_date_string_indo": "02 Januari 1997",
                                                "profile_pic": "",
                                                "phone": "081234567892",
                                                "marital_status": true,
                                                "email": "mawar@mail.com",
                                                "nira": "1928374650",
                                                "sip": "",
                                                "created_at": "2022-12-14T19:01:32.822+07:00",
                                                "updated_at": "2022-12-14T19:01:32.822+07:00",
                                                "deleted_at": null
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
                    "Nurses"
                ],
                "summary": "Create",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "name" : "Ns. Mawar",
                                    "doctor_id" : 7,
                                    "license_number" : "1234567894",
                                    "password" : "mawar123",
                                    "speciality_id" : 1,
                                    "profile_pic": "",
                                    "birth_date": "1997-01-02",
                                    "phone": "081234567892",
                                    "marital_status" : true,
                                    "email" : "mawar@mail.com",
                                    "nira" : "1928374650",
                                    "sip" : ""
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
                                            "id": 2,
                                            "user_id": 16,
                                            "doctor_id": 7,
                                            "name": "",
                                            "speciality_id": 1,
                                            "license_number": "1234567894",
                                            "speciality_name": "Umum 2",
                                            "birth_date": "1997-01-02T00:00:00+07:00",
                                            "birth_date_string": "1997-1-02",
                                            "birth_date_string_indo": "02 Januari 1997",
                                            "profile_pic": "",
                                            "phone": "081234567892",
                                            "marital_status": true,
                                            "email": "mawar@mail.com",
                                            "nira": "1928374650",
                                            "sip": "",
                                            "created_at": "2022-12-14T19:01:32.822+07:00",
                                            "updated_at": "2022-12-14T19:01:32.822+07:00",
                                            "deleted_at": null
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
        "/v1/nurses/{id}": {
            "get": {
                "tags": [
                    "Nurses"
                ],
                "summary": "Get By Id",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of Nurse",
                        "schema": {
                            "type": "integer"
                        },
                        "example": 2
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
                                            "user_id": 15,
                                            "doctor_id": 7,
                                            "name": "Ns. Melati",
                                            "speciality_id": 1,
                                            "license_number": "1234567893",
                                            "speciality_name": "Umum 2",
                                            "birth_date": "1997-01-01T17:00:00+07:00",
                                            "birth_date_string": "1997-1-01",
                                            "birth_date_string_indo": "01 Januari 1997",
                                            "profile_pic": "",
                                            "phone": "081234567892",
                                            "marital_status": true,
                                            "email": "melati@mail.com",
                                            "nira": "1928374650",
                                            "sip": "",
                                            "created_at": "2022-12-14T09:50:10.05+07:00",
                                            "updated_at": "2022-12-14T09:50:10.05+07:00",
                                            "deleted_at": null
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
                    "Nurses"
                ],
                "summary": "Update",
                "requestBody": {
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "object",
                                "example": {
                                    "name" : "Ns. Melati",
                                    "license_number" : "1234567893",
                                    "password" : "melati123",
                                    "speciality_id" : 1,
                                    "profile_pic": "",
                                    "birth_date": "1997-01-02",
                                    "phone": "081234567892",
                                    "marital_status" : true,
                                    "email" : "melati@mail.com",
                                    "nira" : "1928374650",
                                    "sip" : ""
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
                                        "message": "success create data",
                                        "data": {
                                            "id": 2,
                                            "user_id": 13,
                                            "name": "",
                                            "speciality_id": 1,
                                            "license_number": "1234567893",
                                            "speciality_name": "Umum 2",
                                            "birth_date": "1997-01-02T00:00:00+07:00",
                                            "birth_date_string": "1997-1-02",
                                            "birth_date_string_indo": "02 Januari 1997",
                                            "marital_status": true,
                                            "email": "melati@mail.com",
                                            "nira": "1928374650",
                                            "sip": "",
                                            "created_at": "2022-12-14T15:54:27.317+07:00",
                                            "updated_at": "2022-12-14T15:54:27.317+07:00",
                                            "deleted_at": null
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
                    "Nurses"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of Nurse",
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
        "/v1/nurses/license_number/{license_number}": {
            "get": {
                "tags": [
                    "Nurses"
                ],
                "summary": "Get By License Number",
                "parameters": [
                    {
                        "name": "license_number",
                        "in": "path",
                        "description": "License Number of Nurse",
                        "schema": {
                            "type": "string"
                        },
                        "example": 1234567893
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
                                            "user_id": 15,
                                            "doctor_id": 7,
                                            "name": "Ns. Melati",
                                            "speciality_id": 1,
                                            "license_number": "1234567893",
                                            "speciality_name": "Umum 2",
                                            "birth_date": "1997-01-01T17:00:00+07:00",
                                            "birth_date_string": "1997-1-01",
                                            "birth_date_string_indo": "01 Januari 1997",
                                            "profile_pic": "",
                                            "phone": "081234567892",
                                            "marital_status": true,
                                            "email": "melati@mail.com",
                                            "nira": "1928374650",
                                            "sip": "",
                                            "created_at": "2022-12-14T09:50:10.05+07:00",
                                            "updated_at": "2022-12-14T09:50:10.05+07:00",
                                            "deleted_at": null
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
        "/v1/notifications": {
            "get": {
                "tags": [
                    "Notifications"
                ],
                "summary": "Get Notifications",
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
                                        "message": "success get data",
                                        "data": [
                                            {
                                                "outpatient_session_id": 13,
                                                "description": "Request Kunjungan",
                                                "schedule_date": "2022-12-15",
                                                "schedule_date_indo": "15 Desember 2022",
                                                "schedule_time": "05:00"
                                            },
                                            {
                                                "outpatient_session_id": 14,
                                                "description": "Request Kunjungan",
                                                "schedule_date": "2022-12-15",
                                                "schedule_date_indo": "15 Desember 2022",
                                                "schedule_time": "05:01"
                                            },
                                            {
                                                "outpatient_session_id": 15,
                                                "description": "Request Kunjungan",
                                                "schedule_date": "2022-12-15",
                                                "schedule_date_indo": "15 Desember 2022",
                                                "schedule_time": "05:01"
                                            },
                                            {
                                                "outpatient_session_id": 16,
                                                "description": "Request Kunjungan",
                                                "schedule_date": "2022-12-15",
                                                "schedule_date_indo": "15 Desember 2022",
                                                "schedule_time": "05:01"
                                            }
                                        ]
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

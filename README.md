<p align="center">
  <a href="" rel="noopener">
 <img width=300px height=auto src="./assets/logo.png" alt="Project logo"></a>
</p>

<h3 align="center">Konsultanku-GSC</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![Golang](https://img.shields.io/badge/Go-%2300ADD8.svg?style=flat&logo=go&logoColor=white)](https://golang.org/)
[![Firebase](https://img.shields.io/badge/Firebase-%23039BE5.svg?style=flat&logo=firebase)](https://firebase.google.com/)
[![Google Cloud Platform](https://img.shields.io/badge/Google_Cloud_Platform-%234285F4.svg?style=flat&logo=google-cloud&logoColor=white)](https://cloud.google.com/)
[![Docker](https://img.shields.io/badge/Docker-%230db7ed.svg?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)
[![Fiber](https://img.shields.io/badge/Fiber-%2320232a.svg?style=flat&logo=fiber&logoColor=%2361DAFB)](https://github.com/gofiber/fiber)
[![Gorm](https://img.shields.io/badge/Gorm-%2300ADD8.svg?style=flat&logo=gorm&logoColor=white)](https://gorm.io/)

</div>

## Architecture
<p align="center">
  <a href="" rel="noopener">
 <img width=400px height=auto src="./assets/architecture.jpg" alt="Project logo"></a>
</p>

## API Documentation
- [Auth](#auth)
- [Student](#student)
### Auth
This API field use for authentification which use firebase one of part google technology
- [Register](#register)
- [Login](#login)
- [Reset Password](#reset-password)
- [Logout](#logout)
#### Register
- Method : POST
- Endpoint : ```/api/auth/register```
- Header :
  - Content-Type: application/x-www-form-urlencoded
- Request
  - **email**: (string, required)
  - **password**: (string, required)
  - **phone_number**: (string, required)
  - **name**: (string, required)
  - **file**: (file, required)
  - **role**: (string, required) Only Student or MSME
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
```
#### Login
- Method : POST
- Endpoint : ```/api/auth/login```
- Header :
  - Content-Type: application/json
- Request
```json
{
  "email": "string",
  "password": "string"
}
```
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
```
#### Reset Password
- Method : POST
- Endpoint : ```/api/auth/reset-password```
- Header :
  - Content-Type: application/json
- Request
```json
{
  "email": "string"
}
```
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
```
#### Logout
- Method : POST
- Endpoint : ```/api/auth/logout```
- Header :
  - Content-Type: application/json
- Request
```
None
```
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
```
### Student
This API field use for student role
- [Add Profile](#add-profile)
- [Accept Offer](#accept-offer)
- [Add Comment](#add-comment)
- [Create Team](#create-team)
- [Join Team](#join-team)
- [Get Own Profile](#get-student-profile)
- [Get Problems](#get-problems)
- [Get Collaborations](#get-collaborations)
#### Add Profile
- Method : POST
- Endpoint : ```/api/student/profile```
- Header :
  - Content-Type: application/json
- Request
```json
{
  "major": "string",
  "university": "string",
  "class_of": "string"
}
```
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
```
#### Accept Offer
- Method : POST
- Endpoint : ```/api/student/collaboration/:msmeId```
- Header :
  - Content-Type: application/json
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
```
#### Add Comment
- Method : POST
- Endpoint : ```/api/student/comment/:problemId```
- Header :
  - Content-Type: application/json
- Request
```json
{
  "content": "string"
}
```
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
```
#### Accept Offer
- Method : POST
- Endpoint : ```/api/student/collaboration/:msmeId```
- Header :
  - Content-Type: application/json
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
```
#### Create Team
- Method : POST
- Endpoint : ```/api/student/team/create```
- Header :
  - Content-Type: application/json
- Request
```json
{
  "name": "string"
}
```
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
```
#### Join Team
- Method : POST
- Endpoint : ```/api/student/team/join/:teamId```
- Header :
  - Content-Type: application/json
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
```
#### Get Student Profile
- Method : GET
- Endpoint : ```/api/student/profile```
- Header :
  - Content-Type: application/json
- Response
```json
{
  "success": {
    "code": 200,
    "data": {
      "user" : {
        "uid": "string",
        "email": "string",
        "display_name": "string",
        "phone_number": "string",
        "photo_url": "string"
      }
    },
    "major": "string",
    "university": "string",
    "class_of": "string",
    "tags" : [
      {
        "id": "string",
        "name": "string"
      }
    ],
    "team": {
      "id": "string",
      "name": "string"
    },
    "collaboration" : [
      {
        "msme" : {
        "uid": "string",
        "email": "string",
        "display_name": "string",
        "phone_number": "string",
        "photo_url": "string"
        },
        "in_collaboration": "bool"
      }
    ]
  }
}
```
#### Get Problems
- Method : POST
- Endpoint : ```/api/student/problems```
- Header :
  - Content-Type: application/json
- Response
```json
{
  "success": {
    "code": 200,
    "data": [
      {
        "id": "string",
        "like": "int",
        "comment_count": "int",
        "created": "int",
        "title": "string",
        "content": "string",
        "msme": {
          "user" : {
          "uid": "string",
          "email": "string",
          "display_name": "string",
          "phone_number": "string",
          "photo_url": "string"
          },
          "name": "string",
          "since": "string"
        },
        "tags": [
          {
          "id": "string",
          "name": "string"
          }
        ],
      }
    ]
  }
}
```
#### Get Collaborations
- Method : GET
- Endpoint : ```/api/student/collaborations```
- Header :
  - Content-Type: application/json
- Response
```json
{
  "success": {
    "code": 200,
    "data": [
      {
        "msme": {
          "uid": "string",
          "email": "string",
          "display_name": "string",
          "phone_number": "string",
          "photo_url": "string"
        },
        "in_collaboration": "bool",
        "progress": "int",
        "description": "string",
        "finished": "bool",
        "feedback": "string",
        "rating": "float"
      }
    ]
  }
}
```
### Msme
This API field use for MSME role
- [Add Profile](#add-profile)
- [Add Problem](#add-problem)
- [Add Collaboration](#add-collaboration)
- [Give Progress](#give-progress)
- [End Collaboration](#end-collaboration)
- [Get Own Profile](#get-msme-profile)
- [Get Problems](#get-problems)
- [Get Collaborations](#get-collaborations)
#### Add Profile
- Method : POST
- Endpoint : ```/api/msme/profile```
- Header :
  - Content-Type: application/json
- Request
```json
{
  "name": "string",
  "since": "string",
  "type": "string"
}
```
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
```
#### Accept Problem
- Method : POST
- Endpoint : ```/api/msme/problem```
- Header :
  - Content-Type: application/json
- Request
```json
{
  "title": "string",
  "tag": [0,1,2,3],
  "type": "string"
}
```
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
``` 
#### Add Collaboration
- Method : POST
- Endpoint : ```/api/msme/collaboration/:studentId```
- Header :
  - Content-Type: application/json
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
``` 
#### Give Progress
- Method : PUT
- Endpoint : ```/api/msme/progress/:studentId```
- Header :
  - Content-Type: application/json
- Request
```json
{
  "progress": "int",
  "description": "string"
}
```
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
``` 
#### End Collaboration
- Method : PUT
- Endpoint : ```/api/msme/end-collaboration/:studentId```
- Header :
  - Content-Type: application/json
- Response
```json
{
  "success": {
    "code": 200,
    "data": "string"
  }
}
``` 
#### Get Msme Profile
- Method : GET
- Endpoint : ```/api/msme/profile```
- Header :
  - Content-Type: application/json
- Response
```json
{
  "success": {
    "code": 200,
    "data": {
      "user": {
          "uid": "string",
          "email": "string",
          "display_name": "string",
          "phone_number": "string",
          "photo_url": "string"
        },
      "name": "string",
      "since": "string",
      "type": "string",
      "problem": [
        {
        "id": "string",
        "like": "int",
        "comment_count": "int",
        "created": "int",
        "title": "string",
        "content": "string",
        "tag": [
            {
            "id": "int",
            "name": "string"
            }
          ]
        }
      ],
      "collaboration": [
        {
        "student": {
          "uid": "string",
          "email": "string",
          "display_name": "string",
          "phone_number": "string",
          "photo_url": "string"
        },
        "in_collaboration": "bool"
        }
      ]
    }
  }
}
``` 
#### Get Comments
- Method : GET
- Endpoint : ```/api/msme/comments```
- Header :
  - Content-Type: application/json
- Response
```json
{
  "success": {
    "code": 200,
    "data": [
      {
        "id": "string",
        "content": "string",
        "team": {
          "id": "string",
          "name": "string",
          "rating": "float"
        },
        "student": {
          "user" : {
            "uid": "string",
            "email": "string",
            "display_name": "string",
            "phone_number": "string",
            "photo_url": "string"
          },
          "major": "string",
          "university": "string",
          "class_of": "string"
        }
      }
    ]
  }
}
``` 
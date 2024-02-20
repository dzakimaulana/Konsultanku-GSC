<p align="center">
  <a href="" rel="noopener">
 <img width=300px height=auto src="./assets/logo.jpg" alt="Project logo"></a>
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

## Architecture🧱
<p align="center">
  <a href="" rel="noopener">
 <img width=400px height=auto src="./assets/architecture.jpg" alt="Project logo"></a>
</p>

## API Documentation🗄️
### Auth
This API field use for authentification which use firebase one of part google technology
#### Register
- Method : POST
- Endpoint : ```/api/auth/register```
- Header :
  - Content-Type: application/x-www-form-urlencoded
- Request
  - **id**: (string, required) The unique identifier for the item.
  - **name**: (string, required) The name of the item.
  - **price**: (number, required) The price of the item.
  - **quantity**: (integer, required) The quantity of the item.
- Response
```json
{
  "success": {
		"code": "int",
		"data": "string"
  }
}
```


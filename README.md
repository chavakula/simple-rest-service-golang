# Simple REST service in GOLANG

This example will show you how to quickly create a REST service in GOLANG.

## Installation

Use the package manager [dep](https://golang.github.io/dep/) to install dependencies for this project.

## Usage

```bash
cd src/restservice
go run main.go
```

server will be up at http://localhost:8000

REST CALLS

Endpoint : http://localhost:8000/api/sayhello

Method   : GET 

Output   :
```json
{
    "data": {
        "name": "Rajshekar",
        "phone": "+91-9922616111",
        "email": "ch.rajshekar@gmail.com"
    },
    "message": "success",
    "status": true
}
```

Endpoint : http://localhost:8000/api/sayhello

Method   : POST

Input    : 
```json
{"name":"Json","phone":"9922616111","email":"ch.rajshekar@gmail.com"}
```

Output   :
```json
{
    "data": {
        "name": "Json",
        "phone": "9922616111",
        "email": "ch.rajshekar@gmail.com"
    },
    "message": "success",
    "status": true
}
```
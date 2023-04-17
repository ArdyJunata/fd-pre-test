# Female Daily - Back End Developer (Pre-interview Test)

## Requirement

- Go **(go1.20.1 darwin/arm64)**
- Docker **(v20.10.17)**

## Usage

Run database

```bash
docker compose up -d
```

run app

```
go run cmd/api/main.go
```


## API Spec

## Fetch User

### GET /user/fetch

Request Query Params

```
/user/fetch?page=1
```

Response Body

```json
{
    "message": "fetch user success",
    "data": [
        {
            "id": 3,
            "email": "emma.wong@reqres.in",
            "first_name": "Emma",
            "last_name": "Wong",
            "avatar": "https://reqres.in/img/faces/3-image.jpg"
        },
        {
            "id": 4,
            "email": "eve.holt@reqres.in",
            "first_name": "Eve",
            "last_name": "Holt",
            "avatar": "https://reqres.in/img/faces/4-image.jpg"
        }
    ]
}
```
## Get One User

### GET /user/:id

Response Body

```json
{
    "message": "find one user success",
    "data": {
        "id": 2,
        "email": "janet.weaver@reqres.in",
        "first_name": "Janet",
        "last_name": "Weaver",
        "avatar": "https://reqres.in/img/faces/2-image.jpg",
        "created_at": "2023-04-16T17:59:13.79616Z",
        "updated_at": "2023-04-16T17:59:13.796161Z",
        "deleted_at": null
    }
}
```

## Get All User

### GET /user

Response Body

```json
{
    "message": "find all user success",
    "data": [
        {
            "id": 2,
            "email": "janet.weaver@reqres.in",
            "first_name": "Janet",
            "last_name": "Weaver",
            "avatar": "https://reqres.in/img/faces/2-image.jpg",
            "created_at": "2023-04-16T17:59:13.79616Z",
            "updated_at": "2023-04-16T17:59:13.796161Z",
            "deleted_at": null
        },
        {
            "id": 3,
            "email": "emma.wong@reqres.in",
            "first_name": "Emma",
            "last_name": "Wong",
            "avatar": "https://reqres.in/img/faces/3-image.jpg",
            "created_at": "2023-04-16T17:59:13.823727Z",
            "updated_at": "2023-04-16T17:59:13.823727Z",
            "deleted_at": null
        },
    ]
}
```

## Create User

### POST /user

Request Body

```json
{
    "email": "ardyjunata53@gmail.com",
    "first_name": "ardy",
    "last_name": "junata",
    "avatar": "https://prnt.sc/RA1Vic8mz1Ym"
}
```

Response Body

```json
{
    "message": "create user success"
}
```

## Update User

### PUT /user/:id

Request Body

```json
{
    "email": "ardyjunata53@gmail.com",
    "first_name": "ardy",
    "last_name": "junata",
    "avatar": "https://prnt.sc/RA1Vic8mz1Ym"
}
```

Response Body

```json
{
    "message": "update user success"
}
```

## Delete User

### DELETE /user/:id

Request Headers
```
Authorization: 3cdcnTiBsl
```

Response Body

```json
{
    "message": "delete user success"
}
```
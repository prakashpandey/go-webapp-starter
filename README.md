# go-webapp-starter

A web application starter project in golang

## Project Structure

```bash
├── Makefile
├── README.md
├── auth
│   └── auth.go
├── db
│   └── mongodb.go
├── index
│   └── index.go
├── main.go
├── routes.go
└── user
    ├── dao.go
    └── user.go
```

## How to run?

- `git clone https://github.com/prakashpandey/go-webapp-starter.git`

- `cd go-webapp-starter`

- `make build && make run`

- It starts at default address `http://localhost:8284`

### Apis

- `/`
  
- `/user`
  
- `/user/delete`
  
## Packages

### main

Package `main` is our main package. File `main.go` starts our server. File `route.go` defines all routes.

### auth

Contains code to authenticate client request.

### db

`db` package must implement all the dao interfaces.

### index

A sample package implementing hello world functionality.

### user

Package `user` implement user `CRUD` operations.

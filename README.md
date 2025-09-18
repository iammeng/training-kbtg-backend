# Training KBTG Backend

A REST API server built with Go and Fiber framework, featuring JWT authentication and SQLite database.

## Features

- 🚀 RESTful API server built with Go
- 🔐 JWT Authentication (Register/Login)
- 🗄️ SQLite database with GORM ORM
- 📝 Swagger API documentation
- 🛡️ Password hashing with bcrypt
- 🌐 CORS and Logger middleware
- 🔒 Protected routes with JWT middleware

## Quick Start

1. Install dependencies:
```bash
go mod tidy
```

2. Run the server:
```bash
go run main.go
```

3. Access Swagger documentation:
```
http://localhost:3000/swagger/
```

## API Endpoints

### General
- `GET /` - Returns a hello world message
- `GET /swagger/*` - Swagger API documentation

### Authentication
- `POST /auth/register` - Register a new user
- `POST /auth/login` - Login and get JWT token

### Protected Routes
- `GET /protected` - Example protected route (requires JWT token)

## Usage Examples

### Register a new user:
```bash
curl -X POST http://localhost:3000/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"123456"}'
```

### Login:
```bash
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"123456"}'
```

### Access protected route:
```bash
curl -X GET http://localhost:3000/protected \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

## Dependencies

- [Fiber v2](https://github.com/gofiber/fiber) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [SQLite](https://www.sqlite.org/) - Database
- [JWT](https://github.com/golang-jwt/jwt) - JSON Web Tokens
- [bcrypt](https://golang.org/x/crypto/bcrypt) - Password hashing
- [Swagger](https://github.com/swaggo/fiber-swagger) - API documentation

## Database

The application uses SQLite database (`app.db`) which is automatically created and migrated when the server starts.

## Environment

- Go 1.21+
- Port: 3000 (default)
- Database: SQLite (app.db)

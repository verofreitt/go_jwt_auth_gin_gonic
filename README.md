# Go JWT Auth API with Gin & MongoDB

A robust authentication REST API built with Go, Gin Gonic, JWT, and MongoDB. This project demonstrates secure user registration, login, and protected routes using industry best practices.

## Features
- User registration and login with hashed passwords
- JWT-based authentication and authorization
- MongoDB integration for persistent storage
- Middleware for route protection
- Environment variable configuration
- Modular folder structure for scalability
- Example `.env` file for easy setup

## Folder Structure
```
go_jwt_auth_gin_gonic/
├── controllers/         # Request handlers (user, auth)
├── database/            # Database connection logic
├── helpers/             # Helper functions (auth, token)
├── middleware/          # Authentication middleware
├── models/              # Data models (User, etc.)
├── routes/              # Route definitions (auth, user)
├── main.go              # Application entry point
├── go.mod, go.sum       # Go modules
├── .env-exemple         # Example environment variables
```

## Environment Variables
Copy `.env-exemple` to `.env` and fill in your values:
```
PORT=8080
MONGODB_URL=mongodb://localhost:27017/your-db
```

## Prerequisites
- Go 1.18 or higher
- MongoDB instance (local or remote)

## Installation & Usage
1. **Clone the repository:**
   ```bash
   git clone https://github.com/verofreitt/go_jwt_auth_gin_gonic.git
   cd go_jwt_auth_gin_gonic
   ```
2. **Install dependencies:**
   ```bash
   go mod tidy
   ```
3. **Set up your `.env` file:**
   - Copy `.env-exemple` to `.env`
   - Set your `PORT` and `MONGODB_URL` values
4. **Run the application:**
   ```bash
   go run main.go
   ```
5. **API will be available at:**
   - `http://localhost:8080` (or your chosen port)

## API Endpoints
### Auth
- `POST /auth/register` — Register a new user
  - Body: `{ "email": "user@example.com", "password": "yourpassword" }`
- `POST /auth/login` — Login and receive a JWT
  - Body: `{ "email": "user@example.com", "password": "yourpassword" }`
  - Response: `{ "token": "<JWT>" }`

### User (Protected)
- `GET /user/profile` — Get user profile (requires JWT in `Authorization: Bearer <token>` header)

## Example Requests
**Register:**
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"yourpassword"}'
```

**Login:**
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"yourpassword"}'
```

**Get Profile (Protected):**
```bash
curl -X GET http://localhost:8080/user/profile \
  -H "Authorization: Bearer <JWT>"
```

## Technologies Used
- [Go](https://golang.org/)
- [Gin Gonic](https://gin-gonic.com/)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)
- [JWT](https://jwt.io/)
- [godotenv](https://github.com/joho/godotenv)

## License
MIT

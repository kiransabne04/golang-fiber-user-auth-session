# Golang Fiber User Authentication and Session Management Template

## Overview
This project is a Golang template built using the Fiber web framework, providing a comprehensive authentication, authorization, and session management system. It is designed to serve as a starting point for building scalable web and mobile applications, featuring modularity, security, and ease of integration.

## Key Features
- **Session Management (Cookie-based):** 
  - Ideal for web applications with session storage in PostgreSQL.
- **JWT Authentication:**
  - Designed for mobile apps and frontend frameworks (React, Vue, Angular). Supports access and refresh token mechanisms.
- **Middleware:**
  - Exportable middlewares for session validation and JWT verification.
- **Token Refresh API:**
  - Manages refresh token expiry and renewal.
- **Reusable Modular Structure:**
  - Facilitates easy extension for new projects.
- **PostgreSQL Database:**
  - Used for storing user data, sessions, and tokens.
- **Configuration Management:**
  - Uses Viper to manage configuration via `config.json`.
- **Fiber Framework:**
  - Lightweight, fast, and scalable.

## Project Structure
```
┣ 📂cmd
┃ ┗ 📜main.go
┣ 📂config
┃ ┣ 📜app.go
┃ ┣ 📜config.go
┃ ┗ 📜config.json
┣ 📂database
┣ 📂internal
┃ ┣ 📂auth
┃ ┃ ┣ 📜handler.go
┃ ┃ ┣ 📜repository.go
┃ ┃ ┗ 📜service.go
┃ ┣ 📂middleware
┃ ┃ ┗ 📜session_middleware.go
┃ ┣ 📂session
┃ ┃ ┣ 📜repository.go
┃ ┃ ┗ 📜service.go
┃ ┣ 📂user
┃ ┃ ┣ 📜handler.go
┃ ┃ ┣ 📜repository.go
┃ ┃ ┗ 📜service.go
┃ ┣ 📂utils
┃ ┗ 📜services.go
┣ 📂middleware
┃ ┗ 📜session_pg_middleware.go
┣ 📂pkg
┃ ┣ 📜database.go
┃ ┣ 📜jwt.go
┃ ┣ 📜logger.go
┃ ┗ 📜response.go
┣ 📂routers
┃ ┣ 📜auth_routes.go
┃ ┣ 📜router.go
┃ ┗ 📜user_routes.go
┣ 📂utils
┣ 📜database.sql
┣ 📜go.mod
┣ 📜go.sum
┗ 📜readme.md
```

## Installation and Setup
### Prerequisites
- Golang 1.18+
- PostgreSQL database

### Steps
1. **Clone the Repository:**
   ```bash
   git clone https://github.com/kiransabne04/golang-fiber-user-auth-session.git
   cd golang-fiber-user-auth-session
   ```
2. **Install Dependencies:**
   ```bash
   go mod tidy
   ```
3. **Configure Application:**
   Update `config/config.json` with your database and application configuration.
   ```json
   {
     "db_url": "postgres://postgres:root@localhost:5432/reporting_db?sslmode=disable",
     "jwt_secret": "supersecretkey",
     "server_port": "8000",
     "token_ttl": 900,
     "refresh_ttl": 86400,
     "session_ttl_mins": 15
   }
   ```
4. **Run Database Migrations:**
   ```bash
   psql -U <username> -d <database> -f database.sql
   ```
5. **Start the Application:**
   ```bash
   go run cmd/main.go
   ```

## API Endpoints
### Authentication
- `POST /auth/login` - User login (returns JWT tokens)
- `POST /auth/register` - User registration
- `POST /auth/refresh` - Refresh JWT tokens

### User Management
- `GET /user/me` - Get current user information
- `PUT /user/update` - Update user details

### Session Management
- `POST /session` - Create a new session
- `GET /session` - Get active sessions
- `DELETE /session` - Logout and remove session

## Contribution
Contributions are welcome! Please open an issue for discussion. Many features can be added & few things can be improved for more better usage. I started this project for having a basic starter template, but will continue to enhance it in future. Any suggestion is welcome at kiransabne04@gmail.com


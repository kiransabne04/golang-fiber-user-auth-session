Key Features
Session Management (Cookie-based):
For web applications where sessions are stored in PostgreSQL.
JWT Authentication:
For mobile apps or frontend frameworks like React, Vue, Angular.
Includes access token & refresh token mechanisms.
Middleware:
Exportable middlewares for session validation and JWT verification.
Token Refresh API:
Handles JWT refresh token expiry and token renewal.
Reusable Structure:
Modular structure to serve as a starting codebase for other projects.
PostgreSQL Database:
Store users, sessions, and tokens.
Configuration via Environment Variables:
Use viper for configuration management.
Fiber Framework:
Lightweight, fast, and scalable.

.
├── README.md
├── cmd
│   ├── app.go            # Application setup and initialization
│   └── main.go           # Entry point
├── config
│   ├── config.go         # Configuration management
│   └── config.json       # Default configuration
├── database
│   ├── migrations.sql    # Database schema
├── internal
│   ├── auth              # Authentication module
│   │   ├── handler.go    # HTTP handlers for login, register, etc.
│   │   ├── middleware.go # Middleware for authentication
│   │   ├── service.go    # Business logic for authentication
│   │   └── repository.go # Database queries for authentication
│   ├── session           # Session management module
│   │   ├── handler.go    # HTTP handlers for session APIs
│   │   ├── middleware.go # Middleware for session validation
│   │   ├── service.go    # Business logic for session
│   │   └── repository.go # Database queries for sessions
│   ├── user              # User module
│   │   ├── handler.go    # HTTP handlers for user management
│   │   ├── middleware.go # Middleware for user validation
│   │   ├── service.go    # Business logic for users
│   │   └── repository.go # Database queries for users
├── pkg
│   ├── database.go       # Database connection setup
│   ├── jwt.go            # JWT utility functions
│   ├── logger.go         # Logging utilities
│   ├── response.go       # Response formatting utilities
├── go.mod
├── go.sum
└── .env                  # Environment variables

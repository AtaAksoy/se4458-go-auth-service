# SE4458 Go Auth Service

A professional, production-ready authentication service built with Go, featuring user registration, login, JWT token generation, and comprehensive API documentation with Swagger/OpenAPI.

## 🏗️ Architecture & Design

### Design Principles
- **Clean Architecture**: Separation of concerns with distinct layers (Handler → Service → Repository → Model)
- **RESTful API**: Standard HTTP methods and status codes
- **Security First**: Password hashing with bcrypt, JWT tokens for authentication
- **Database Agnostic**: GORM ORM for easy database switching
- **Comprehensive Documentation**: Auto-generated Swagger/OpenAPI documentation

### Technology Stack
- **Language**: Go 1.20+
- **Framework**: Gin (HTTP router)
- **ORM**: GORM with MySQL driver
- **Authentication**: JWT tokens
- **Password Hashing**: bcrypt
- **Documentation**: Swagger/OpenAPI
- **Environment**: .env configuration
- **Containerization**: Docker

### Project Structure
```
se4458-go-auth-service/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── handler/
│   │   └── auth.go          # HTTP handlers (Register, Login)
│   ├── service/
│   │   └── auth.go          # Business logic layer
│   ├── repository/
│   │   └── user.go          # Data access layer
│   └── model/
│       ├── user.go          # Data models and DTOs
│       └── response.go      # Response structures
├── docs/                    # Auto-generated Swagger docs
├── .env                     # Environment variables
├── .gitignore              # Git ignore rules
├── Dockerfile              # Docker configuration
├── go.mod                  # Go module file
└── README.md               # This file
```

## 📊 Data Models (ER Diagram)

### User Entity
```
┌─────────────────┐
│      User       │
├─────────────────┤
│ id (PK)         │ uint, auto-increment
│ name            │ varchar(255), required
│ email           │ varchar(255), unique, required
│ password        │ varchar(255), hashed, required
└─────────────────┘
```

### API Request/Response Models

#### Register Request
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword123"
}
```

#### Login Request
```json
{
  "email": "john@example.com",
  "password": "securepassword123"
}
```

#### Auth Response (Success)
```json
{
  "status": "success",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

#### Error Response
```json
{
  "status": "error",
  "error": "email already exists"
}
```

## 🚀 Features

### Core Functionality
- ✅ User registration with validation
- ✅ User login with credential verification
- ✅ Password hashing using bcrypt
- ✅ JWT token generation and validation
- ✅ Structured JSON responses
- ✅ Comprehensive error handling

### Security Features
- ✅ Password hashing with bcrypt (cost: 10)
- ✅ JWT tokens with 72-hour expiration
- ✅ Input validation and sanitization
- ✅ Environment-based configuration
- ✅ CORS support for frontend integration

### Developer Experience
- ✅ Auto-generated Swagger/OpenAPI documentation
- ✅ Hot-reload development setup
- ✅ Docker containerization
- ✅ Comprehensive logging
- ✅ Environment-based configuration

## 🛠️ Installation & Setup

### Prerequisites
- Go 1.20 or higher
- MySQL 5.7+ or MySQL 8.0+
- Docker (optional, for containerized deployment)

### 1. Clone the Repository
```bash
git clone <repository-url>
cd se4458-go-auth-service
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Environment Configuration
Create a `.env` file in the project root:
```env
# Database Configuration
DB_USER=root
DB_PASS=your_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=auth_service

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key_here

# Server Configuration
PORT=8080
```

### 4. Database Setup
```sql
CREATE DATABASE auth_service;
```
> Note: Tables are automatically created by GORM AutoMigrate

### 5. Swagger Documentation
```bash
# Install Swagger CLI (first time only)
go install github.com/swaggo/swag/cmd/swag@latest

# Generate documentation
swag init -g cmd/main.go
```

### 6. Run the Application
```bash
go run cmd/main.go
```

The server will start at `http://localhost:8080`

## 🐳 Docker Deployment

### Build and Run with Docker
```bash
# Build the image
docker build -t auth-service .

# Run the container
docker run -p 8080:8080 --env-file .env auth-service
```

### Docker Compose (Recommended)
Create a `docker-compose.yml` file:
```yaml
version: '3.8'
services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=mysql
      - DB_PORT=3306
      - DB_USER=root
      - DB_PASS=password
      - DB_NAME=auth_service
    depends_on:
      - mysql
    networks:
      - auth-network

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: auth_service
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql
    networks:
      - auth-network

volumes:
  mysql_data:

networks:
  auth-network:
    driver: bridge
```

Run with:
```bash
docker-compose up -d
```

## 📚 API Documentation

### Swagger UI
Access the interactive API documentation at:
```
http://localhost:8080/swagger/index.html
```

### Available Endpoints

#### POST /register
Register a new user account.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword123"
}
```

**Response (201 Created):**
```json
{
  "status": "success",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

#### POST /login
Authenticate an existing user.

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "securepassword123"
}
```

**Response (200 OK):**
```json
{
  "status": "success",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

## 🔧 Development

### Project Structure Explanation

#### Handler Layer (`internal/handler/`)
- **Purpose**: HTTP request/response handling
- **Responsibilities**: 
  - Parse incoming requests
  - Call service layer
  - Format responses
  - Handle HTTP status codes

#### Service Layer (`internal/service/`)
- **Purpose**: Business logic implementation
- **Responsibilities**:
  - Password hashing/verification
  - JWT token generation
  - Business rule validation
  - Orchestrating repository calls

#### Repository Layer (`internal/repository/`)
- **Purpose**: Data access abstraction
- **Responsibilities**:
  - Database operations
  - Query optimization
  - Data persistence

#### Model Layer (`internal/model/`)
- **Purpose**: Data structures and DTOs
- **Responsibilities**:
  - Define data models
  - Request/Response DTOs
  - Validation tags

### Adding New Features

1. **New Endpoint**: Add handler in `internal/handler/`
2. **Business Logic**: Implement in `internal/service/`
3. **Data Access**: Add methods in `internal/repository/`
4. **Models**: Define in `internal/model/`
5. **Documentation**: Add Swagger comments

## 🔒 Security Considerations

### Implemented Security Measures
- **Password Hashing**: bcrypt with cost factor 10
- **JWT Tokens**: 72-hour expiration with secure signing
- **Input Validation**: Required field validation and email format checking
- **Environment Variables**: Sensitive data stored in .env files
- **CORS Configuration**: Specific origin allowlisting

### Recommended Production Security
- Use HTTPS in production
- Implement rate limiting
- Add request logging and monitoring
- Use strong, unique JWT secrets
- Regular security audits
- Database connection pooling
- Implement refresh tokens

## 🧪 Testing

### Manual Testing
1. Start the application
2. Access Swagger UI at `http://localhost:8080/swagger/index.html`
3. Test register endpoint with valid data
4. Test login endpoint with registered credentials
5. Verify JWT token generation

### Automated Testing (Future Enhancement)
- Unit tests for service layer
- Integration tests for API endpoints
- Database migration tests
- Security vulnerability tests

## 📈 Performance Considerations

### Current Optimizations
- GORM connection pooling
- Efficient password hashing
- Structured JSON responses
- Minimal memory footprint

## 🙏 Acknowledgments

- [Gin Framework](https://github.com/gin-gonic/gin) for HTTP routing
- [GORM](https://gorm.io/) for database ORM
- [Swagger](https://swagger.io/) for API documentation
- [JWT-Go](https://github.com/golang-jwt/jwt) for JWT implementation
- [bcrypt](https://golang.org/x/crypto/bcrypt) for password hashing

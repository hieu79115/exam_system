# Exam System

An examination management system built with Go, using Gin framework and GORM ORM.

## Technologies

- **Language**: Go 1.25.1
- **Web Framework**: Gin
- **ORM**: GORM
- **Database**: MySQL

## Project Structure

```
exam_system/
├── api/                    # API specifications
├── cmd/
│   └── server/            # Application entry point
│       └── main.go
├── config/                # Database configuration and settings
│   └── database.go
├── internal/              # Internal source code
│   ├── app/              # Application initialization
│   │   └── app.go
│   ├── dto/              # Data Transfer Objects
│   │   └── exam_response.go
│   ├── handlers/         # HTTP handlers
│   │   └── exam_handler.go
│   ├── models/           # Database models
│   │   └── examination.go
│   ├── repository/       # Data access layer
│   │   └── exam_repo.go
│   ├── routes/           # Route definitions
│   │   └── routes.go
│   └── service/          # Business logic
│       └── exam_service.go
├── pkg/
│   └── utils/            # Utility functions
├── .env                  # Environment variables
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Requirements

- Go 1.25.1 or higher
- MySQL

## Installation

1. Clone repository:
```bash
git clone <repository-url>
cd exam_system
```

2. Install dependencies:
```bash
go mod download
```

3. Create `.env` file and configure database:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=exam_system
```

4. Run migrations (if any)

## Running the Application

```bash
go run cmd/server/main.go
```

Server will run at `http://localhost:8080`

### API Documentation

Once the server is running, you can access the Swagger UI documentation at:

```
http://localhost:8080/swagger/index.html
```

The API documentation provides:
- Interactive API testing interface
- Complete endpoint descriptions
- Request/response schemas
- Example requests

### Regenerate Swagger Documentation

After modifying API endpoints or adding new handlers, regenerate the Swagger docs:

```bash
swag init -g cmd/server/main.go
```

## Build

```bash
go build -o server cmd/server/main.go
```

## API Endpoints

The API is organized into three main resource groups:

### Examinations
- `GET /api/v1/examinations` - List all examinations
- `GET /api/v1/examinations/:id` - Get examination details
- `POST /api/v1/examinations` - Create new examination
- `PUT /api/v1/examinations/:id` - Update examination
- `DELETE /api/v1/examinations/:id` - Delete examination

### Questions
- `GET /api/v1/questions?examId={id}` - List questions by exam
- `GET /api/v1/questions/:id` - Get question details
- `POST /api/v1/questions` - Create new question
- `PUT /api/v1/questions/:id` - Update question
- `DELETE /api/v1/questions/:id` - Delete question

### Reading Passages
- `GET /api/v1/passages` - List all reading passages
- `GET /api/v1/passages/:id` - Get passage details
- `POST /api/v1/passages` - Create new passage
- `PUT /api/v1/passages/:id` - Update passage
- `DELETE /api/v1/passages/:id` - Delete passage

For detailed API documentation with request/response schemas, visit the Swagger UI at `/swagger/index.html`

## Architecture

The project uses Layered Architecture:

- **Handler Layer**: Handles HTTP requests/responses
- **Service Layer**: Contains business logic
- **Repository Layer**: Database access
- **Model Layer**: Defines data structures

## License

All rights reserved

## Author

- **Name**: Nguyễn Phúc Thịnh
- **Email**: npt911@gmail.com
- **GitHub**: [@hieu79115](https://github.com/hieu79115)

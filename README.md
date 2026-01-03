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

## Build

```bash
go build -o server cmd/server/main.go
```

## API Endpoints

(Add information about your API endpoints here)

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

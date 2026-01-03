# Exam System

Hệ thống quản lý thi cử được xây dựng bằng Go, sử dụng Gin framework và GORM ORM.

## Công nghệ sử dụng

- **Language**: Go 1.25.1
- **Web Framework**: Gin
- **ORM**: GORM
- **Database**: MySQL

## Cấu trúc dự án

```
exam_system/
├── api/                    # API specifications
├── cmd/
│   └── server/            # Entry point của ứng dụng
│       └── main.go
├── config/                # Cấu hình database và các thiết lập khác
│   └── database.go
├── internal/              # Mã nguồn nội bộ
│   ├── app/              # Khởi tạo ứng dụng
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

## Yêu cầu

- Go 1.25.1 hoặc cao hơn
- MySQL

## Cài đặt

1. Clone repository:
```bash
git clone <repository-url>
cd exam_system
```

2. Cài đặt dependencies:
```bash
go mod download
```

3. Tạo file `.env` và cấu hình database:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=exam_system
```

4. Chạy migration (nếu có)

## Chạy ứng dụng

```bash
go run cmd/server/main.go
```

Server sẽ chạy tại `http://localhost:8080`

## Build

```bash
go build -o server cmd/server/main.go
```

## API Endpoints

(Thêm thông tin về các API endpoints của bạn ở đây)

## Kiến trúc

Dự án sử dụng kiến trúc phân lớp (Layered Architecture):

- **Handler Layer**: Xử lý HTTP requests/responses
- **Service Layer**: Chứa business logic
- **Repository Layer**: Truy cập database
- **Model Layer**: Định nghĩa cấu trúc dữ liệu

## License

All rights reserved

## Tác giả

- **Name**: Nguyễn Phúc Thịnh
- **Email**: npt911@gmail.com
- **GitHub**: [@hieu79115](https://github.com/hieu79115)

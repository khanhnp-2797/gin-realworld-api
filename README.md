# Gin RealWorld API

API backend được xây dựng với Go, Gin Framework và PostgreSQL theo chuẩn RealWorld API.

## Công nghệ sử dụng

- **Go** 1.21+
- **Gin** - HTTP Web Framework
- **GORM** - ORM cho Go
- **PostgreSQL** - Cơ sở dữ liệu
- **JWT** - Authentication
- **godotenv** - Quản lý biến môi trường

## Cấu trúc thư mục

```
gin-realworld-api/
├── controllers/                 # HTTP handlers
│   ├── auth_controller.go      # Register, Login
│   └── user_controller.go      # User management
├── services/                    # Business logic
│   └── user_service.go
├── repositories/                # Data access layer
│   └── user_repository.go
├── models/                      # Database models (GORM)
│   └── user.go
├── middlewares/                 # HTTP middlewares
│   └── auth.go                 # JWT authentication & CORS
├── dto/                         # Data Transfer Objects
│   └── user_dto.go             # Request/Response structs
├── config/                      # Configuration
│   ├── config.go               # Environment variables
│   ├── database.go             # Database connection
│   └── migration.go            # Auto migrations
├── utils/                       # Helper functions
│   ├── auth.go                 # JWT & password hashing
│   └── errors.go               # Error responses
├── routes/                      # Route definitions
│   └── routes.go
├── main.go                      # Entry point
├── .env                         # Environment variables
├── .env.example
├── go.mod
├── go.sum
└── README.md
```

## Kiến trúc

```
┌─────────────────────────────────────────────────┐
│              HTTP Request                       │
└─────────────────┬───────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────┐
│  Controllers (HTTP Handlers)                    │
│  - AuthController: Register, Login              │
│  - UserController: GetCurrentUser, UpdateUser   │
└─────────────────┬───────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────┐
│  Services (Business Logic)                      │
│  - UserService                                  │
└─────────────────┬───────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────┐
│  Repositories (Data Access)                     │
│  - UserRepository (Interface-based)             │
└─────────────────┬───────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────┐
│           PostgreSQL Database                   │
└─────────────────────────────────────────────────┘
```

## Cài đặt

### 1. Clone repository

```bash
git clone <repository-url>
cd gin-realworld-api
```

### 2. Cài đặt dependencies

```bash
go mod download
```

### 3. Cấu hình database

Tạo database PostgreSQL:

```bash
createdb realworld_db
```

### 4. Cấu hình biến môi trường

File `.engin_realworld
```

### 4. Cấu hình biến môi trường

Copy file `.env.example` và chỉnh sửa thông tin:

```bash
cp .env.example .env
```

Nội dung `.env`:

```env
ENV=development
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=gin_realworld
DB_SSLMODE=disable

JWT_SECRET=your-secret-key-change-this-in-production
JWT_EXPIRATION_HOURS=72

go run main.go
```

### Build và chạy

```bash
# Build binary
go build -o bin/app .
go build -o bin/app cmd/app/main.go

# Run binary
./bAPI Endpoints

Tuân thủ [RealWorld API Spec](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints)

### Authentication (Public)

- `POST /api/users` - Đăng ký user mới
- `POST /api/users/login` - Đăng nhập

### User (Protected - requires `Authorization: Token <jwt>`)

- `GET /api/user` - Lấy thông tin user hiện tại
- `PUT /api/user` - Cập nhật thông tin usn tại
- `PUT /api/user` - Cập nhật thông tin user

### Health Check

- `GET /health` - Kiểm tra trạng thái server

## Request/Response Format

**Tất cả requests và responses đều wrap data trong object `user`:**

### Đăng ký - POST /api/users
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "username": "jake",
      "email": "jake@jake.jake",
      "password": "jakejake"
    }
  }'
```

**Response (201 Created):**
```json
{
  "user": {
    "username": "jake",
    "email": "jake@jake.jake",
    "bio": "",
    "image": "",
    "token": "jwt.token.here"
  }
}
```

### Đăng nhập - POST /api/users/login
```bash
curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "email": "jake@jake.jake",
      "password": "jakejake"
    }
  }'
```

### Lấy thông tin user - GET /api/user
```bash
curl -X GET http://localhost:8080/api/user \
  -H "Authorization: Token YOUR_JWT_TOKEN"
```

### Cập nhật user - PUT /api/user
```bash
curl -X PUT http://localhost:8080/api/user \
  -H "Content-Type: application/json" \
  -H "Authorization: Token YOUR_JWT_TOKEN" \
  -d '{
    "user": {
      "bio": "I like to code",
      "image": "https://example.com/avatar.jpg"
    }
  }'
```

### Error Response Format
```json
{
  "errors": {
    "body": ["error message"]
  }
}
```

**Lưu ý:**
- Sử dụng `Token` thay vì `Bearer` trong Authorization header
- Status code 422 (Unprocessable Entity) cho validation errors

## Ví dụ sử dụng API

### Đăng ký user

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "username": "jake",
      "email": "jake@jake.jake",
      "password": "jakejake"
    }
  }'
```
uild binary
go build -o bin/app cmd/app/main.go

# Format code
go fmt ./...

# Run tests
go test ./...

# Check for errors
go vet ./...
```

## Tips cho người mới học Go

1. **Exported vs Unexported**:
   - `UserService` (uppercase) = public/exported
   - `userService` (lowercase) = private/unexported

2. **Error handling**: Luôn check error ngay sau function call

3. **Nil checking**: Check nil trước khi dereference pointer

4. **Go is not OOP**: Không có classes, inheritance. Dùng composition.

5. **Defer**: Dùng để cleanup (như finally trong JS)
   ```go
   defer config.CloseDatabase()  // Sẽ chạy khi function return
   ```

## Deployment

### Docker (nếu cần)

```bash
# Build image
docker build -t gin-realworld-api .

# Run container
docker run -p 8080:8080 --env-file .env gin-realworld-api
```

## License

MIT
main.go

# Build binary
go build -o bin/app .

# Format code
go fmt ./...

# Run tests
go test ./...

# Check for errors
go vet ./...
```

## Resources

- [Go Documentation](https://go.dev/doc/)
- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [RealWorld API Spec](https://realworld-docs.netlify.app/docs/specs/backend-specs/endpoints)

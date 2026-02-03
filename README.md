# Gin RealWorld API

A RealWorld API implementation using Go (Golang) with Gin framework, following the [RealWorld API Spec](https://realworld-docs.netlify.app/specifications/backend/api-response-format/).

## 🚀 Tech Stack

- **Framework**: [Gin](https://gin-gonic.com/) - High-performance HTTP web framework
- **Database**: PostgreSQL with [GORM](https://gorm.io/) ORM
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Configuration**: godotenv for environment variables
- **Validation**: Gin validator (built-in)

## 📁 Project Structure

```
gin-realworld-api/
├── config/           # Configuration and database setup
├── controllers/      # HTTP request handlers
├── dto/             # Data Transfer Objects (request/response)
├── middlewares/     # Custom middleware (auth, CORS)
├── models/          # Database models
├── repositories/    # Data access layer
├── routes/          # Route definitions
├── services/        # Business logic layer
├── utils/           # Utility functions (JWT, slug, errors)
├── docs/            # Documentation
├── .env.example     # Environment variables template
├── go.mod           # Go module dependencies
└── main.go          # Application entry point
```

## 🔧 Installation & Setup

### Prerequisites
- Go 1.24 or higher
- PostgreSQL database

### Steps

1. **Clone the repository**
```bash
git clone https://github.com/khanhnp-2797/gin-realworld-api.git
cd gin-realworld-api
```

2. **Install dependencies**
```bash
go mod download
```

3. **Configure environment variables**
```bash
cp .env.example .env
```

Edit `.env` with your configuration:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=realworld_db
DB_SSLMODE=disable

JWT_SECRET=your-secret-key-here
JWT_EXPIRATION_HOURS=72

SERVER_PORT=8080
```

4. **Run the application**
```bash
go run main.go
```

The server will start at `http://localhost:8080`

## 📚 API Endpoints

### Authentication
- `POST /api/users` - Register a new user
- `POST /api/users/login` - Login (returns JWT token with Bearer prefix)

### User Management
- `GET /api/user` - Get current user (requires auth)
- `PUT /api/user` - Update current user (requires auth)

### Articles
- `GET /api/articles` - List articles (with pagination, filters)
  - Query params: `limit`, `offset`, `tag`, `author`, `favorited`
- `GET /api/articles/feed` - Get user's feed (requires auth)
- `GET /api/articles/:slug` - Get article by slug
- `POST /api/articles` - Create article (requires auth)
- `PUT /api/articles/:slug` - Update article (requires auth)
- `DELETE /api/articles/:slug` - Delete article (requires auth)

### Favorites
- `POST /api/articles/:slug/favorite` - Favorite article (requires auth)
- `DELETE /api/articles/:slug/favorite` - Unfavorite article (requires auth)

### Comments
- `GET /api/articles/:slug/comments` - Get comments for article
- `POST /api/articles/:slug/comments` - Add comment (requires auth)
- `DELETE /api/articles/:slug/comments/:id` - Delete comment (requires auth)

### Tags
- `GET /api/tags` - Get all tags

## 🔐 Authentication

The API uses JWT for authentication. After login, the token is returned with `Bearer` prefix:

```json
{
  "user": {
    "email": "user@example.com",
    "token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "username": "username",
    "bio": "I work at State Farm",
    "image": "https://api.realworld.io/images/demo-avatar.png"
  }
}
```

Include the token in the Authorization header for protected endpoints:
```
Authorization: Bearer <your-token>
```

## 📊 Database Schema

### Main Tables
- **users** - User accounts
- **articles** - Blog articles with slug-based URLs
- **comments** - Comments on articles
- **tags** - Article tags
- **favorites** - User favorites (many-to-many: users ↔ articles)
- **follows** - User following relationships
- **article_tags** - Article-tag associations (many-to-many)

See [database_schema.md](docs/database_schema.md) for detailed schema.

## ✨ Features

### Implemented
- ✅ User registration and authentication (JWT)
- ✅ Token format: `Bearer <token>` (only login returns token)
- ✅ User profile management
- ✅ CRUD operations for articles
- ✅ Article pagination with filters (tag, author, favorited)
- ✅ Automatic slug generation from title
- ✅ Slug update when title changes
- ✅ Comments system
- ✅ Favorites system
- ✅ Tags management
- ✅ User feed (articles from followed users)
- ✅ Comprehensive validation
- ✅ Structured error handling
- ✅ CORS support

### Security Features
- Password hashing with bcrypt
- JWT token validation middleware
- Authorization checks for protected resources
- Input validation on all endpoints

## 🧪 API Testing

### Example Requests

**Register**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "username": "testuser",
      "email": "test@example.com",
      "password": "password123"
    }
  }'
```

**Login**
```bash
curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "email": "test@example.com",
      "password": "password123"
    }
  }'
```

**Create Article**
```bash
curl -X POST http://localhost:8080/api/articles \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{
    "article": {
      "title": "How to train your dragon",
      "description": "Ever wonder how?",
      "body": "You have to believe",
      "tagList": ["dragons", "training"]
    }
  }'
```

## 🛠️ Error Handling

The API returns structured error responses:

```json
{
  "errors": {
    "body": [
      "email already exists"
    ]
  }
}
```

### HTTP Status Codes
- `200` - Success
- `201` - Created
- `204` - No Content (successful deletion)
- `401` - Unauthorized (missing/invalid token)
- `403` - Forbidden (insufficient permissions)
- `404` - Not Found
- `409` - Conflict (duplicate resource)
- `422` - Unprocessable Entity (validation errors)
- `500` - Internal Server Error

## 📝 Validation Rules

### User
- Username: 3-50 characters
- Email: Valid email format
- Password: Minimum 8 characters

### Article
- Title: 1-200 characters (required)
- Description: 1-500 characters (required)
- Body: Minimum 1 character (required)

### Comment
- Body: 1-2000 characters (required)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License.

## 👨‍💻 Author

**Nguyen Phi Khanh**
- GitHub: [@khanhnp-2797](https://github.com/khanhnp-2797)

## 🙏 Acknowledgments

- [RealWorld API Spec](https://realworld-docs.netlify.app/)
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)

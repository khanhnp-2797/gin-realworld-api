# API Documentation

## Base URL
```
http://localhost:8080/api
```

## Authentication
Most endpoints require authentication. Include the JWT token in the Authorization header:
```
Authorization: Bearer <your-token>
```

## Endpoints Summary

| Method | Endpoint | Auth Required | Description |
|--------|----------|---------------|-------------|
| POST | `/users` | No | Register new user |
| POST | `/users/login` | No | Login user |
| GET | `/user` | Yes | Get current user |
| PUT | `/user` | Yes | Update user |
| GET | `/articles` | No | List articles |
| GET | `/articles/feed` | Yes | Get user feed |
| GET | `/articles/:slug` | No | Get article |
| POST | `/articles` | Yes | Create article |
| PUT | `/articles/:slug` | Yes | Update article |
| DELETE | `/articles/:slug` | Yes | Delete article |
| POST | `/articles/:slug/favorite` | Yes | Favorite article |
| DELETE | `/articles/:slug/favorite` | Yes | Unfavorite article |
| GET | `/articles/:slug/comments` | No | Get comments |
| POST | `/articles/:slug/comments` | Yes | Add comment |
| DELETE | `/articles/:slug/comments/:id` | Yes | Delete comment |
| GET | `/tags` | No | Get all tags |

## Detailed Endpoints

### Users

#### Register
```http
POST /api/users
Content-Type: application/json

{
  "user": {
    "username": "johndoe",
    "email": "john@example.com",
    "password": "password123"
  }
}
```

Response (201):
```json
{
  "user": {
    "username": "johndoe",
    "email": "john@example.com",
    "bio": "",
    "image": ""
  }
}
```

#### Login
```http
POST /api/users/login
Content-Type: application/json

{
  "user": {
    "email": "john@example.com",
    "password": "password123"
  }
}
```

Response (200):
```json
{
  "user": {
    "username": "johndoe",
    "email": "john@example.com",
    "bio": "",
    "image": "",
    "token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Articles

#### List Articles
```http
GET /api/articles?limit=20&offset=0&tag=dragons&author=johndoe&favorited=janedoe
```

Query Parameters:
- `limit` (default: 20) - Number of articles
- `offset` (default: 0) - Offset for pagination
- `tag` - Filter by tag
- `author` - Filter by author username
- `favorited` - Filter by favorited username

#### Create Article
```http
POST /api/articles
Authorization: Bearer <token>
Content-Type: application/json

{
  "article": {
    "title": "How to train your dragon",
    "description": "Ever wonder how?",
    "body": "You have to believe",
    "tagList": ["dragons", "training"]
  }
}
```

### Comments

#### Get Comments
```http
GET /api/articles/how-to-train-your-dragon/comments
```

#### Add Comment
```http
POST /api/articles/how-to-train-your-dragon/comments
Authorization: Bearer <token>
Content-Type: application/json

{
  "comment": {
    "body": "Great article!"
  }
}
```

### Tags

#### Get All Tags
```http
GET /api/tags
```

Response:
```json
{
  "tags": ["dragons", "training", "golang", "web"]
}
```

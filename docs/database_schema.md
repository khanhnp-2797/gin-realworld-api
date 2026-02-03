# Database Schema

## Tables

### users
- id (PK, auto increment)
- username (unique, not null)
- email (unique, not null)
- password (hashed, not null)
- bio (text, nullable)
- image (varchar, nullable)
- created_at
- updated_at
- deleted_at (soft delete)

### articles
- id (PK, auto increment)
- slug (unique, not null)
- title (not null)
- description (not null)
- body (text, not null)
- author_id (FK -> users.id, not null)
- created_at
- updated_at
- deleted_at (soft delete)

### tags
- id (PK, auto increment)
- name (unique, not null)
- created_at

### article_tags (many-to-many)
- article_id (FK -> articles.id)
- tag_id (FK -> tags.id)

### comments
- id (PK, auto increment)
- body (text, not null)
- article_id (FK -> articles.id, not null)
- author_id (FK -> users.id, not null)
- created_at
- updated_at
- deleted_at (soft delete)

### favorites (many-to-many)
- id (PK, auto increment)
- user_id (FK -> users.id, not null)
- article_id (FK -> articles.id, not null)
- created_at
- Unique constraint: (user_id, article_id)

### follows (many-to-many)
- id (PK, auto increment)
- follower_id (FK -> users.id, not null)
- following_id (FK -> users.id, not null)
- created_at
- Unique constraint: (follower_id, following_id)

## Indexes
- users: username, email
- articles: slug, author_id
- comments: article_id, author_id
- favorites: user_id, article_id
- follows: follower_id, following_id

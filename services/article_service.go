package services

import (
	"errors"
	"strings"

	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"github.com/khanhnp-2797/gin-realworld-api/repositories"
	"github.com/khanhnp-2797/gin-realworld-api/utils"
	"gorm.io/gorm"
)

type ArticleService interface {
	CreateArticle(userID uint, req *dto.CreateArticleRequest) (*dto.ArticleResponse, error)
	GetArticle(slug string, userID *uint) (*dto.ArticleResponse, error)
	UpdateArticle(slug string, userID uint, req *dto.UpdateArticleRequest) (*dto.ArticleResponse, error)
	DeleteArticle(slug string, userID uint) error
	GetFeed(userID uint, limit, offset int) (*dto.ArticlesResponse, error)
}

type articleService struct {
	articleRepo repositories.ArticleRepository
	userRepo    repositories.UserRepository
}

func NewArticleService(articleRepo repositories.ArticleRepository, userRepo repositories.UserRepository) ArticleService {
	return &articleService{
		articleRepo: articleRepo,
		userRepo:    userRepo,
	}
}

func (s *articleService) CreateArticle(userID uint, req *dto.CreateArticleRequest) (*dto.ArticleResponse, error) {
	// Tạo slug từ title
	articleSlug := utils.MakeSlug(req.Article.Title)

	// Tạo article
	article := &models.Article{
		Slug:        articleSlug,
		Title:       req.Article.Title,
		Description: req.Article.Description,
		Body:        req.Article.Body,
		AuthorID:    userID,
	}

	// Xử lý tags
	var tags []models.Tag
	for _, tagName := range req.Article.TagList {
		tag, err := s.articleRepo.FindOrCreateTag(strings.TrimSpace(tagName))
		if err != nil {
			return nil, err
		}
		tags = append(tags, *tag)
	}
	article.Tags = tags

	if err := s.articleRepo.Create(article); err != nil {
		return nil, err
	}

	// Load lại article với author
	article, err := s.articleRepo.FindBySlug(articleSlug)
	if err != nil {
		return nil, err
	}

	return s.buildArticleResponse(article, userID), nil
}

func (s *articleService) GetArticle(slug string, userID *uint) (*dto.ArticleResponse, error) {
	article, err := s.articleRepo.FindBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("article not found")
		}
		return nil, err
	}

	var currentUserID uint
	if userID != nil {
		currentUserID = *userID
	}

	return s.buildArticleResponse(article, currentUserID), nil
}

func (s *articleService) UpdateArticle(slug string, userID uint, req *dto.UpdateArticleRequest) (*dto.ArticleResponse, error) {
	article, err := s.articleRepo.FindBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("article not found")
		}
		return nil, err
	}

	// Kiểm tra quyền sở hữu
	if article.AuthorID != userID {
		return nil, errors.New("unauthorized to update this article")
	}

	// Update các fields
	if req.Article.Title != "" {
		article.Title = req.Article.Title
		article.Slug = utils.MakeSlug(req.Article.Title)
	}
	if req.Article.Description != "" {
		article.Description = req.Article.Description
	}
	if req.Article.Body != "" {
		article.Body = req.Article.Body
	}

	if err := s.articleRepo.Update(article); err != nil {
		return nil, err
	}

	// Load lại article
	article, err = s.articleRepo.FindBySlug(article.Slug)
	if err != nil {
		return nil, err
	}

	return s.buildArticleResponse(article, userID), nil
}

func (s *articleService) DeleteArticle(slug string, userID uint) error {
	article, err := s.articleRepo.FindBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("article not found")
		}
		return err
	}

	// Kiểm tra quyền sở hữu
	if article.AuthorID != userID {
		return errors.New("unauthorized to delete this article")
	}

	return s.articleRepo.Delete(article)
}

func (s *articleService) GetFeed(userID uint, limit, offset int) (*dto.ArticlesResponse, error) {
	if limit == 0 {
		limit = 20
	}

	articles, err := s.articleRepo.GetFeed(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	var articleList []dto.ArticleData
	for _, article := range articles {
		articleList = append(articleList, *s.buildArticleData(&article, userID))
	}

	return &dto.ArticlesResponse{
		Articles:      articleList,
		ArticlesCount: len(articleList),
	}, nil
}

func (s *articleService) buildArticleResponse(article *models.Article, userID uint) *dto.ArticleResponse {
	return &dto.ArticleResponse{
		Article: *s.buildArticleData(article, userID),
	}
}

func (s *articleService) buildArticleData(article *models.Article, userID uint) *dto.ArticleData {
	tagList := make([]string, len(article.Tags))
	for i, tag := range article.Tags {
		tagList[i] = tag.Name
	}

	return &dto.ArticleData{
		Slug:           article.Slug,
		Title:          article.Title,
		Description:    article.Description,
		Body:           article.Body,
		TagList:        tagList,
		CreatedAt:      article.CreatedAt,
		UpdatedAt:      article.UpdatedAt,
		Favorited:      false, // TODO: implement favorite functionality
		FavoritesCount: 0,     // TODO: implement favorite count
		Author: dto.ArticleAuthor{
			Username:  article.Author.Username,
			Bio:       article.Author.Bio,
			Image:     article.Author.Image,
			Following: false, // TODO: implement following check
		},
	}
}

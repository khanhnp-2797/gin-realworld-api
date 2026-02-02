package services

import (
	"errors"

	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"github.com/khanhnp-2797/gin-realworld-api/repositories"
	"gorm.io/gorm"
)

type CommentService interface {
	AddComment(slug string, userID uint, req *dto.AddCommentRequest) (*dto.CommentResponse, error)
	GetComments(slug string, userID *uint) (*dto.CommentsResponse, error)
	DeleteComment(slug string, commentID uint, userID uint) error
}

type commentService struct {
	commentRepo repositories.CommentRepository
	articleRepo repositories.ArticleRepository
}

func NewCommentService(commentRepo repositories.CommentRepository, articleRepo repositories.ArticleRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
		articleRepo: articleRepo,
	}
}

func (s *commentService) AddComment(slug string, userID uint, req *dto.AddCommentRequest) (*dto.CommentResponse, error) {
	// Tìm article
	article, err := s.articleRepo.FindBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("article not found")
		}
		return nil, err
	}

	// Tạo comment
	comment := &models.Comment{
		Body:      req.Comment.Body,
		ArticleID: article.ID,
		AuthorID:  userID,
	}

	if err := s.commentRepo.Create(comment); err != nil {
		return nil, err
	}

	// Load lại comment với author
	comment, err = s.commentRepo.FindByID(comment.ID)
	if err != nil {
		return nil, err
	}

	return s.buildCommentResponse(comment, userID), nil
}

func (s *commentService) GetComments(slug string, userID *uint) (*dto.CommentsResponse, error) {
	// Kiểm tra article có tồn tại
	_, err := s.articleRepo.FindBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("article not found")
		}
		return nil, err
	}

	// Lấy comments
	comments, err := s.commentRepo.FindByArticleSlug(slug)
	if err != nil {
		return nil, err
	}

	var currentUserID uint
	if userID != nil {
		currentUserID = *userID
	}

	var commentList []dto.CommentData
	for _, comment := range comments {
		commentList = append(commentList, *s.buildCommentData(&comment, currentUserID))
	}

	return &dto.CommentsResponse{
		Comments: commentList,
	}, nil
}

func (s *commentService) DeleteComment(slug string, commentID uint, userID uint) error {
	// Kiểm tra article có tồn tại
	_, err := s.articleRepo.FindBySlug(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("article not found")
		}
		return err
	}

	// Tìm comment
	comment, err := s.commentRepo.FindByID(commentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("comment not found")
		}
		return err
	}

	// Kiểm tra quyền sở hữu
	if comment.AuthorID != userID {
		return errors.New("unauthorized to delete this comment")
	}

	return s.commentRepo.Delete(comment)
}

func (s *commentService) buildCommentResponse(comment *models.Comment, userID uint) *dto.CommentResponse {
	return &dto.CommentResponse{
		Comment: *s.buildCommentData(comment, userID),
	}
}

func (s *commentService) buildCommentData(comment *models.Comment, userID uint) *dto.CommentData {
	return &dto.CommentData{
		ID:        comment.ID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Body:      comment.Body,
		Author: dto.CommentAuthor{
			Username:  comment.Author.Username,
			Bio:       comment.Author.Bio,
			Image:     comment.Author.Image,
			Following: false, // TODO: implement following check
		},
	}
}

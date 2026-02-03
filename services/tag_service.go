package services

import (
	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/repositories"
)

type TagService interface {
	GetTags() (*dto.TagsResponse, error)
}

type tagService struct {
	tagRepo repositories.TagRepository
}

func NewTagService(tagRepo repositories.TagRepository) TagService {
	return &tagService{tagRepo: tagRepo}
}

func (s *tagService) GetTags() (*dto.TagsResponse, error) {
	tags, err := s.tagRepo.GetAllTags()
	if err != nil {
		return nil, err
	}

	return &dto.TagsResponse{
		Tags: tags,
	}, nil
}

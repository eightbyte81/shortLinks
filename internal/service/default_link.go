package service

import (
	"shortLinks/internal/model"
	"shortLinks/internal/repository"
)

type DefaultLinkService struct {
	repo                     repository.DefaultLink
	shortLinkService         *ShortLinkService
	defaultShortLinksService *DefaultShortLinksService
}

func NewDefaultLinkService(repo repository.DefaultLink) *DefaultLinkService {
	return &DefaultLinkService{repo: repo}
}

func (s *DefaultLinkService) SetDefaultLink(link model.Link) (int, error) {
	defaultLinkId, err := s.repo.SetDefaultLink(link)
	if err != nil {
		return -1, err
	}

	generatedShortLink := s.shortLinkService.GenerateShortLink(link.LinkData)
	shortLinkId, err := s.shortLinkService.SetShortLink(model.Link{LinkData: generatedShortLink})
	if err != nil {
		return -1, err
	}

	err = s.defaultShortLinksService.SetDefaultShortLinks(defaultLinkId, shortLinkId)
	if err != nil {
		return -1, err
	}

	return defaultLinkId, nil
}

func (s *DefaultLinkService) GetDefaultLinkById(id int) (model.Link, error) {
	return s.repo.GetDefaultLinkById(id)
}

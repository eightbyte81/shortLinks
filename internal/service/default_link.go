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

	generatedShortLink, err := s.shortLinkService.GenerateShortLink(link)
	if err != nil {
		return -1, err
	}
	shortLinkId, err := s.shortLinkService.SetShortLink(generatedShortLink)
	if err != nil {
		return -1, err
	}

	err = s.defaultShortLinksService.SetDefaultShortLinks(model.Link{Id: defaultLinkId}, model.Link{Id: shortLinkId})
	if err != nil {
		return -1, err
	}

	return defaultLinkId, nil
}

func (s *DefaultLinkService) GetDefaultLinkId(link model.Link) (int, error) {
	return s.repo.GetDefaultLinkId(link)
}

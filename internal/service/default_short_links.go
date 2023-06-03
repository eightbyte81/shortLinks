package service

import (
	"shortLinks/internal/model"
	"shortLinks/internal/repository"
)

type DefaultShortLinksService struct {
	repo repository.DefaultShortLinks
}

func NewDefaultShortLinksService(repo repository.DefaultShortLinks) *DefaultShortLinksService {
	return &DefaultShortLinksService{repo: repo}
}

func (s *DefaultShortLinksService) SetDefaultShortLinks(defaultLinkId int, shortLinkId int) error {
	return s.repo.SetDefaultShortLinks(defaultLinkId, shortLinkId)
}

func (s *DefaultShortLinksService) GetDefaultShortLinksByShortLinkId(shortLinkId int) (model.DefaultShortLinks, error) {
	return s.repo.GetDefaultShortLinksByShortLinkId(shortLinkId)
}

func (s *DefaultShortLinksService) GetDefaultShortLinksByDefaultLinkId(defaultLinkId int) (model.DefaultShortLinks, error) {
	return s.repo.GetDefaultShortLinksByDefaultLinkId(defaultLinkId)
}

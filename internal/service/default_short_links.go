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

func (s *DefaultShortLinksService) SetDefaultShortLinks(defaultLink model.Link, shortLink model.Link) error {
	return s.repo.SetDefaultShortLinks(defaultLink, shortLink)
}

func (s *DefaultShortLinksService) GetShortLinkByDefaultLink(defaultLink model.Link) (string, error) {
	return s.repo.GetShortLinkByDefaultLink(defaultLink)
}

func (s *DefaultShortLinksService) GetDefaultLinkByShortLink(shortLink model.Link) (string, error) {
	return s.repo.GetDefaultLinkByShortLink(shortLink)
}

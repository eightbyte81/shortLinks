package service

import (
	"shortLinks/internal/model"
	"shortLinks/internal/repository"
)

type LinksDbService struct {
	repo         repository.LinksPostgres
	linksService LinksService
}

func NewLinksDbService(repo repository.LinksPostgres) *LinksDbService {
	return &LinksDbService{repo: repo}
}

func (s *LinksDbService) GetShortLinkByDefaultLink(defaultLink model.Link) (model.Link, error) {
	shortLink := s.linksService.GenerateShortLink(defaultLink)

	return s.repo.CreateAndGetShortLink(defaultLink, shortLink)
}

func (s *LinksDbService) GetDefaultLinkByShortLinkData(shortLinkData string) (model.Link, error) {
	return s.repo.GetDefaultLinkByShortLinkData(shortLinkData)
}

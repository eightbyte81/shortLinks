package service

import (
	"shortLinks/internal/model"
	"shortLinks/internal/repository"
)

type LinksCacheService struct {
	repo         repository.LinksCache
	linksService LinksService
}

func NewLinksCacheService(repo repository.LinksCache) *LinksCacheService {
	return &LinksCacheService{repo: repo}
}

func (s *LinksCacheService) SetLinksInCache(defaultLink model.Link) (string, error) {
	shortLink := s.linksService.GenerateShortLink(defaultLink)

	s.repo.SetLinksInCache(shortLink.LinkData, defaultLink.LinkData)

	return shortLink.LinkData, nil
}

func (s *LinksCacheService) GetDefaultLinkFromCacheByShortLink(shortLink string) (string, error) {
	return s.repo.GetDefaultLinkFromCacheByShortLink(shortLink)
}

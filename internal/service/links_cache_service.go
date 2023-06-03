package service

import "shortLinks/internal/repository"

type LinksCacheService struct {
	repo             repository.LinksCache
	shortLinkService *ShortLinkService
}

func NewLinksCacheService(repo repository.LinksCache) *LinksCacheService {
	return &LinksCacheService{repo: repo}
}

func (s *LinksCacheService) SetLinksInCache(defaultLink string) (string, error) {
	shortLink := s.shortLinkService.GenerateShortLink(defaultLink)

	s.repo.SetLinksInCache(shortLink, defaultLink)

	return shortLink, nil
}

func (s *LinksCacheService) GetDefaultLinkFromCacheByShortLink(shortLink string) (string, error) {
	return s.repo.GetDefaultLinkFromCacheByShortLink(shortLink)
}

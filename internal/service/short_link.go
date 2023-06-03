package service

import (
	"shortLinks/internal/model"
	"shortLinks/internal/repository"
	"shortLinks/internal/shorturl"
)

type ShortLinkService struct {
	repo     repository.ShortLink
	shortUrl *shorturl.ShortUrl
}

func NewShortLinkService(repo repository.ShortLink, shortUrl *shorturl.ShortUrl) *ShortLinkService {
	return &ShortLinkService{repo: repo, shortUrl: shortUrl}
}

func (s *ShortLinkService) GenerateShortLink(defaultLink string) string {
	return s.shortUrl.Generate(defaultLink)
}

func (s *ShortLinkService) SetShortLink(link model.Link) (int, error) {
	return s.repo.SetShortLink(link)
}

func (s *ShortLinkService) GetShortLinkById(id int) (model.Link, error) {
	return s.repo.GetShortLinkById(id)
}

func (s *ShortLinkService) GetShortLinkByLinkData(linkData string) (model.Link, error) {
	return s.repo.GetShortLinkByLinkData(linkData)
}

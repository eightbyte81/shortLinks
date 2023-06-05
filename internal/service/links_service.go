package service

import (
	"shortLinks/internal/model"
	"shortLinks/internal/shorturl"
)

type LinksService struct {
	shortUrl *shorturl.ShortUrl
}

func NewLinksService(shortUrl *shorturl.ShortUrl) *LinksService {
	return &LinksService{shortUrl: shortUrl}
}

func (s *LinksService) GenerateShortLink(defaultLink model.Link) model.Link {
	return model.Link{LinkData: s.shortUrl.Generate(defaultLink.LinkData)}
}

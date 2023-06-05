package service

import (
	"shortLinks/internal/model"
	"shortLinks/internal/repository"
	"shortLinks/internal/shorturl"
)

type Links interface {
	GenerateShortLink(defaultLink model.Link) model.Link
}

type LinksDb interface {
	GetShortLinkByDefaultLink(defaultLink model.Link) (model.Link, error)
	GetDefaultLinkByShortLinkData(shortLinkData string) (model.Link, error)
}

type LinksCache interface {
	SetLinksInCache(defaultLink model.Link) (string, error)
	GetDefaultLinkFromCacheByShortLink(shortLink string) (string, error)
}

type Service struct {
	Links
	LinksDb
	LinksCache
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Links:      NewLinksService(shorturl.NewShortUrl()),
		LinksDb:    NewLinksDbService(repos.LinksPostgres),
		LinksCache: NewLinksCacheService(repos.LinksCache),
	}
}

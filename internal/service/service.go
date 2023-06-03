package service

import (
	"shortLinks/internal/model"
	"shortLinks/internal/repository"
	"shortLinks/internal/shorturl"
)

type DefaultLink interface {
	SetDefaultLink(link model.Link) (int, error)
	GetDefaultLinkById(id int) (model.Link, error)
}

type ShortLink interface {
	GenerateShortLink(defaultLink string) string
	SetShortLink(link model.Link) (int, error)
	GetShortLinkById(id int) (model.Link, error)
	GetShortLinkByLinkData(linkData string) (model.Link, error)
}

type DefaultShortLinks interface {
	SetDefaultShortLinks(defaultLinkId int, shortLinkId int) error
	GetDefaultShortLinksByShortLinkId(shortLinkId int) (model.DefaultShortLinks, error)
	GetDefaultShortLinksByDefaultLinkId(defaultLinkId int) (model.DefaultShortLinks, error)
}

type LinksCache interface {
	SetLinksInCache(defaultLink string) (string, error)
	GetDefaultLinkFromCacheByShortLink(shortLink string) (string, error)
}

type Service struct {
	DefaultLink
	ShortLink
	DefaultShortLinks
	LinksCache
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		DefaultLink:       NewDefaultLinkService(repos),
		ShortLink:         NewShortLinkService(repos, shorturl.NewShortUrl()),
		DefaultShortLinks: NewDefaultShortLinksService(repos),
		LinksCache:        NewLinksCacheService(repos),
	}
}

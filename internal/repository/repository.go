package repository

import (
	"github.com/jmoiron/sqlx"
	"shortLinks/internal/model"
	"shortLinks/internal/repository/cache"
	"shortLinks/internal/repository/postgres"
)

type DefaultLink interface {
	SetDefaultLink(link model.Link) (int, error)
	GetDefaultLinkById(id int) (model.Link, error)
}

type ShortLink interface {
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
	SetLinksInCache(shortLink string, defaultLink string)
	GetDefaultLinkFromCacheByShortLink(shortLink string) (string, error)
}

type Repository struct {
	DefaultLink
	ShortLink
	DefaultShortLinks
	LinksCache
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		DefaultLink:       postgres.NewDefaultLinkPostgres(db),
		ShortLink:         postgres.NewShortLinkPostgres(db),
		DefaultShortLinks: postgres.NewDefaultShortLinksPostgres(db),
		LinksCache:        cache.NewLinksCache(cache.NewCache()),
	}
}

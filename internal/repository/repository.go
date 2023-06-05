package repository

import (
	"github.com/jinzhu/gorm"
	"shortLinks/internal/model"
	"shortLinks/internal/repository/cache"
	"shortLinks/internal/repository/postgres"
)

type LinksPostgres interface {
	CreateAndGetShortLink(defaultLink model.Link, shortLink model.Link) (model.Link, error)
	GetDefaultLinkByShortLinkData(shortLinkData string) (model.Link, error)
}

type LinksCache interface {
	SetLinksInCache(shortLink string, defaultLink string)
	GetDefaultLinkFromCacheByShortLink(shortLink string) (string, error)
}

type Repository struct {
	LinksPostgres
	LinksCache
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		LinksPostgres: postgres.NewLinksPostgres(db),
		LinksCache:    cache.NewLinksCache(cache.NewCache()),
	}
}

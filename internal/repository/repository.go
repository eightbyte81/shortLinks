package repository

import (
	"github.com/jmoiron/sqlx"
	"shortLinks/internal/model"
	"shortLinks/internal/repository/postgres"
)

type DefaultLink interface {
	SetDefaultLink(link model.Link) (int, error)
	GetDefaultLinkId(link model.Link) (int, error)
}

type ShortLink interface {
	SetShortLink(link model.Link) (int, error)
	GetShortLinkId(link model.Link) (int, error)
}

type DefaultShortLinks interface {
	SetDefaultShortLinks(defaultLink model.Link, shortLink model.Link) error
	GetShortLinkByDefaultLink(defaultLink model.Link) (string, error)
	GetDefaultLinkByShortLink(shortLink model.Link) (string, error)
}

type Repository struct {
	DefaultLink
	ShortLink
	DefaultShortLinks
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		DefaultLink:       postgres.NewDefaultLinkPostgres(db),
		ShortLink:         postgres.NewShortLinkPostgres(db),
		DefaultShortLinks: postgres.NewDefaultShortLinksPostgres(db),
	}
}

package service

import (
	"shortLinks/internal/model"
	"shortLinks/internal/repository"
)

type DefaultLink interface {
	SetDefaultLink(link model.Link) (int, error)
	GetDefaultLinkId(link model.Link) (int, error)
}

type ShortLink interface {
	GenerateShortLink(defaultLink model.Link) (model.Link, error)
	SetShortLink(link model.Link) (int, error)
	GetShortLinkId(link model.Link) (int, error)
}

type DefaultShortLinks interface {
	SetDefaultShortLinks(defaultLink model.Link, shortLink model.Link) error
	GetShortLinkByDefaultLink(defaultLink model.Link) (string, error)
	GetDefaultLinkByShortLink(shortLink model.Link) (string, error)
}

type Service struct {
	DefaultLink
	ShortLink
	DefaultShortLinks
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		DefaultLink:       NewDefaultLinkService(repos),
		ShortLink:         NewShortLinkService(repos),
		DefaultShortLinks: NewDefaultShortLinksService(repos),
	}
}

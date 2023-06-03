package service

import (
	"shortLinks/internal/model"
	"shortLinks/internal/repository"
)

type ShortLinkService struct {
	repo repository.ShortLink
}

func NewShortLinkService(repo repository.ShortLink) *ShortLinkService {
	return &ShortLinkService{repo: repo}
}

func (s *ShortLinkService) GenerateShortLink(defaultLink model.Link) (model.Link, error) {
	return model.Link{LinkData: defaultLink.LinkData + "/check"}, nil
}

func (s *ShortLinkService) SetShortLink(link model.Link) (int, error) {
	return s.repo.SetShortLink(link)
}

func (s *ShortLinkService) GetShortLinkId(link model.Link) (int, error) {
	return s.repo.GetShortLinkId(link)
}

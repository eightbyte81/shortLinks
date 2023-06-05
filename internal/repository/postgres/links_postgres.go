package postgres

import (
	"github.com/jinzhu/gorm"
	"shortLinks/internal/model"
)

type LinksPostgres struct {
	db *gorm.DB
}

func NewLinksPostgres(db *gorm.DB) *LinksPostgres {
	return &LinksPostgres{db: db}
}

func (r *LinksPostgres) CreateAndGetShortLink(defaultLink model.Link, shortLink model.Link) (model.Link, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.
			Table(defaultLinkTable).
			Create(&defaultLink).Error; err != nil {
			return err
		}
		if err := r.db.
			Table(shortLinkTable).
			Create(&shortLink).Error; err != nil {
			return err
		}
		linkChain := model.LinkChain{
			DefaultLinkId: defaultLink.Id,
			ShortLinkId:   shortLink.Id,
		}
		if err := r.db.
			Table(linkChainTable).
			Create(&linkChain).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return model.Link{}, err
	}

	return shortLink, nil
}

func (r *LinksPostgres) GetDefaultLinkByShortLinkData(shortLinkData string) (model.Link, error) {
	var defaultLink model.Link
	var shortLink model.Link
	var linkChain model.LinkChain

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.
			Table(shortLinkTable).
			Model(&model.Link{}).
			Find(&shortLink).
			Where("link = ?", shortLinkData).Error; err != nil {
			return err
		}
		if err := r.db.
			Table(linkChainTable).
			Model(&model.LinkChain{}).
			Find(&linkChain).
			Where("short_link_id = ?", shortLink.Id).Error; err != nil {
			return err
		}
		if err := r.db.
			Table(defaultLinkTable).
			Model(&model.Link{}).
			Find(&defaultLink).
			Where("id = ?", linkChain.DefaultLinkId).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return model.Link{}, err
	}

	return defaultLink, nil
}

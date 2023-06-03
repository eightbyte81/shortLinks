package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"shortLinks/internal/model"
)

type DefaultShortLinksPostgres struct {
	db              *sqlx.DB
	defaultLinkRepo *DefaultLinkPostgres
	shortLinkRepo   *ShortLinkPostgres
}

func NewDefaultShortLinksPostgres(db *sqlx.DB) *DefaultShortLinksPostgres {
	return &DefaultShortLinksPostgres{db: db}
}

func (r *DefaultShortLinksPostgres) SetDefaultShortLinks(defaultLink model.Link, shortLink model.Link) error {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (default_link_id, short_link_id) values ($1, $2) RETURNING id", defaultShortLinksTable)
	row := r.db.QueryRow(query, defaultLink.Id, shortLink.Id)

	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}

func (r *DefaultShortLinksPostgres) GetShortLinkByDefaultLink(defaultLink model.Link) (string, error) {
	var shortLink model.Link
	defaultLinkId, err := r.defaultLinkRepo.GetDefaultLinkId(defaultLink)
	if err != nil {
		return "", err
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE default_link_id=$1", defaultShortLinksTable)
	err = r.db.Get(&shortLink, query, defaultLinkId)

	return shortLink.LinkData, err
}

func (r *DefaultShortLinksPostgres) GetDefaultLinkByShortLink(shortLink model.Link) (string, error) {
	var defaultLink model.Link
	shortLinkId, err := r.shortLinkRepo.GetShortLinkId(shortLink)
	if err != nil {
		return "", err
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE short_link_id=$1", defaultShortLinksTable)
	err = r.db.Get(&shortLink, query, shortLinkId)

	return defaultLink.LinkData, err
}

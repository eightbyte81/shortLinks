package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"shortLinks/internal/model"
)

type ShortLinkPostgres struct {
	db *sqlx.DB
}

func NewShortLinkPostgres(db *sqlx.DB) *ShortLinkPostgres {
	return &ShortLinkPostgres{db: db}
}

func (r *ShortLinkPostgres) SetShortLink(link model.Link) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (link) values ($1) RETURNING id", shortLinkTable)
	row := r.db.QueryRow(query, link.LinkData)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *ShortLinkPostgres) GetShortLinkById(id int) (model.Link, error) {
	var shortLink model.Link

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", shortLinkTable)
	err := r.db.Get(&shortLink, query, id)

	return shortLink, err
}

func (r *ShortLinkPostgres) GetShortLinkByLinkData(linkData string) (model.Link, error) {
	var shortLink model.Link

	query := fmt.Sprintf("SELECT * FROM %s WHERE link=$1", shortLinkTable)
	err := r.db.Get(&shortLink, query, linkData)

	return shortLink, err
}

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

func (r *ShortLinkPostgres) GetShortLinkId(link model.Link) (int, error) {
	var shortLink model.Link
	query := fmt.Sprintf("SELECT * FROM %s WHERE link=$1", shortLinkTable)
	err := r.db.Get(&shortLink, query, link.LinkData)

	return shortLink.Id, err
}

package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"shortLinks/internal/model"
)

type DefaultLinkPostgres struct {
	db *sqlx.DB
}

func NewDefaultLinkPostgres(db *sqlx.DB) *DefaultLinkPostgres {
	return &DefaultLinkPostgres{db: db}
}

func (r *DefaultLinkPostgres) SetDefaultLink(link model.Link) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (link) values ($1) RETURNING id", defaultLinkTable)
	row := r.db.QueryRow(query, link.LinkData)

	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *DefaultLinkPostgres) GetDefaultLinkById(id int) (model.Link, error) {
	var defaultLink model.Link
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", defaultLinkTable)
	err := r.db.Get(&defaultLink, query, id)

	return defaultLink, err
}

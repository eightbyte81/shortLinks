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

func (r *DefaultShortLinksPostgres) SetDefaultShortLinks(defaultLinkId int, shortLinkId int) error {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (default_link_id, short_link_id) values ($1, $2) RETURNING id", defaultShortLinksTable)
	row := r.db.QueryRow(query, defaultLinkId, shortLinkId)

	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}

func (r *DefaultShortLinksPostgres) GetDefaultShortLinksByShortLinkId(shortLinkId int) (model.DefaultShortLinks, error) {
	var defaultShortLinks model.DefaultShortLinks

	query := fmt.Sprintf("SELECT * FROM %s WHERE short_link_id=$1", defaultShortLinksTable)
	row := r.db.QueryRow(query, shortLinkId)

	if err := row.Scan(&defaultShortLinks); err != nil {
		return model.DefaultShortLinks{}, err
	}

	return defaultShortLinks, nil
}

func (r *DefaultShortLinksPostgres) GetDefaultShortLinksByDefaultLinkId(defaultLinkId int) (model.DefaultShortLinks, error) {
	var defaultShortLinks model.DefaultShortLinks

	query := fmt.Sprintf("SELECT * FROM %s WHERE default_link_id=$1", defaultShortLinksTable)
	row := r.db.QueryRow(query, defaultLinkId)

	if err := row.Scan(&defaultShortLinks); err != nil {
		return model.DefaultShortLinks{}, err
	}

	return defaultShortLinks, nil
}

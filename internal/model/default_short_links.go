package model

type DefaultShortLinks struct {
	Id            int `db:"id"`
	DefaultLinkId int `db:"default_link_id"`
	ShortLinkId   int `db:"short_link_id"`
}

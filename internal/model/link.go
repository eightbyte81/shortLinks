package model

type Link struct {
	Id       int    `db:"id"`
	LinkData string `db:"link"`
}

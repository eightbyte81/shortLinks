package model

type Link struct {
	Id       int    `gorm:"primaryKey;column:id"`
	LinkData string `gorm:"column:link"`
}

package model

type LinkChain struct {
	Id            int `gorm:"primaryKey;column:id"`
	DefaultLinkId int `gorm:"column:default_link_id"`
	ShortLinkId   int `gorm:"column:short_link_id"`
}

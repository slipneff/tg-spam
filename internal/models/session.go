package models

type Session struct {
	Id   string `gorm:"primaryKey"`
	Path string `gorm:"column:path"`
}

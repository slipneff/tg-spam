package models

type Channel struct {
	Name          string `gorm:"primaryKey"`
	LastMessageID string `gorm:"column:last_message_id"`
}

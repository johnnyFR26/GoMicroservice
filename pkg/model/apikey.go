package model

type APIKey struct {
	Key     string `gorm:"primaryKey;column:key"`
	Owner   string `gorm:"column:owner"`
	Enabled bool   `gorm:"column:enabled"`
}

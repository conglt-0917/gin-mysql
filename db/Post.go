package db

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        uint   `gorm:"primary_key;auto_increment" json:"id"`
	Content   string `json:"content"`
	Owner     string `json:"owner"`
	ActivedAt sql.NullTime
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

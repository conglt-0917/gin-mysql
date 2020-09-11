package db

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string `gorm:"size:255;not null;unique" json:"username"`
	PassWord  string `gorm:"size:255;not null;unique" json:"password"`
	ActivedAt sql.NullTime
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

type LoginStruct struct {
	Username string `form:"username" json:"username" binding:"required" validate:"required"`
	Password string `form:"password" json:"password" binding:"required" validate:"required"`
}

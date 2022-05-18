package entity

import (
	guuid "github.com/google/uuid"
)

type Workout struct {
	ID        guuid.UUID `gorm:"primaryKey" json:"-"`
	Name      string     `json:"name"`
	CreatedAt int64      `gorm:"autoCreateTime" json:"-" `
	UpdatedAt int64      `gorm:"autoUpdateTime:milli" json:"-"`
}

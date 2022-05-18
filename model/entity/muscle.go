package entity

import (
	guuid "github.com/google/uuid"
)

type Muscle struct {
	ID guuid.UUID `gorm:"primaryKey" json:"-"`

	// Data variables
	Name string `json:"name"`

	CreatedAt int64 `gorm:"autoCreateTime" json:"-" `
	UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"-"`
}

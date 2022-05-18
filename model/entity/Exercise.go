package entity

import (
	guuid "github.com/google/uuid"
)

type Exercise struct {
	ID guuid.UUID `gorm:"primaryKey" json:"-"`

	// Data variables
	Name         string `json:"name"`
	Instructions string `json:"Instructions"`

	CreatedAt int64 `gorm:"autoCreateTime" json:"-" `
	UpdatedAt int64 `gorm:"autoUpdateTime:milli" json:"-"`
}

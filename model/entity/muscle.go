package entity

import "github.com/google/uuid"

type Muscle struct {
	Base

	// Keys
	Exercises  []Exercise `json:"muscles" gorm:"many2many:exercise_muscle"`
	BodyPartId uuid.UUID

	// Data variables
	Name     string `json:"name"`
	BodyPart BodyPart
}

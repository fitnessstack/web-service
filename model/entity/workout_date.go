package entity

import "github.com/google/uuid"

type WorkoutDate struct {
	Base

	// Keys
	WorkoutId          uuid.UUID
	CompletedExercises []CompletedExercise

	// Data variables
	Date      string `json:"date"'`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
}

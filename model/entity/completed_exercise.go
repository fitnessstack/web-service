package entity

import "github.com/google/uuid"

type CompletedExercise struct {
	Base

	// Keys
	ExerciseId    uuid.UUID
	WorkoutDateId uuid.UUID

	// Data variables
	Note string `json:"note"'`
	Reps string `json:"reps"`
}

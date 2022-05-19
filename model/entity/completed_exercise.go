package entity

import "github.com/google/uuid"

type CompletedExercise struct {
	Base

	// Keys
	ExerciseId    uuid.UUID
	WorkoutDateId uuid.UUID
	Reps          []Rep

	// Data variables
	Note        string `json:"note"'`
	Exercise    Exercise
	WorkoutDate WorkoutDate
}

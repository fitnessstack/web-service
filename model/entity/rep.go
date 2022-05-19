package entity

import "github.com/google/uuid"

type Rep struct {
	Base

	// Keys
	CompletedExerciseId uuid.UUID

	// Data variables
	Set               int `json:"name"`
	WeightAmount      int `json:"weight_amount"`
	Rep               int `json:"rep"`
	CompletedExercise CompletedExercise
}

package entity

type Exercise struct {
	Base

	// Keys
	Muscles            []Muscle `json:"muscles" gorm:"many2many:exercise_muscle"`
	CompletedExercises []CompletedExercise

	// Data variables
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
}

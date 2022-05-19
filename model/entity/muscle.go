package entity

type Muscle struct {
	Base

	// Keys
	Exercises []Exercise `json:"muscles" gorm:"many2many:exercise_muscle"`

	// Data variables
	Name string `json:"name"`
}

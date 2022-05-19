package entity

type Workout struct {
	Base

	// Keys
	WorkoutDates []WorkoutDate

	// Data variables
	Name string `json:"name"`
}

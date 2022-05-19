package entity

type BodyPart struct {
	Base

	// Keys
	Muscles []Muscle

	// Data variables
	Name string `json:"name"'`
}

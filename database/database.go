package database

import (
	"fmt"
	"github.com/fitnessstack/web-service/model/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	// DBConn is a database connection singleton
	DBConn *gorm.DB
)

func ConnectDatabase() {
	var err error // define error here to prevent overshadowing the global DB

	//Just sqlite for now
	env := os.Getenv("database.name")
	DBConn, err = gorm.Open(sqlite.Open(env+".db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Migrate the schema
	err = DBConn.AutoMigrate(
		&entity.Workout{},
		&entity.Exercise{},
		&entity.Muscle{},
		&entity.CompletedExercise{},
		&entity.WorkoutDate{},
		&entity.Rep{},
		&entity.BodyPart{},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting seeder...")
	// Seed data
	var workout entity.Workout
	if DBConn.First(&workout).Error == gorm.ErrRecordNotFound {
		fmt.Println("No workout record, creating now...")
		//DBConn.Create(&entity.Workout{Name: "Back & Biceps"})
		fmt.Println("Seeding done")
	} else {
		fmt.Println("workout record found")
	}

}

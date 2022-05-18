package database

import (
	"github.com/fitnessstack/web-service/model"
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
	env := os.Getenv("database.name" + ".db")
	DBConn, err = gorm.Open(sqlite.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = DBConn.AutoMigrate(&model.Workout{}, &model.Muscle{})
	if err != nil {
		log.Fatal(err)
	}

}

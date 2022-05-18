package seed

import (
	"github.com/fitnessstack/web-service/model/entity"
	"gorm.io/gorm"
	"log"
)

var workouts = []entity.Workout{
	{
		Name: "Pull up",
	},
	{
		Name: "Dip",
	},
}

func Load(db *gorm.DB) {

	err := db.Migrator().DropTable(
		&entity.Workout{},
		&entity.Muscle{},
	).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(
		&entity.Workout{},
		&entity.Muscle{},
	).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	//TODO Seed Data
}

package repository

import (
	"github.com/fitnessstack/web-service/model/data"
	"github.com/fitnessstack/web-service/model/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkoutRepository struct {
	db *gorm.DB
}

func (r *WorkoutRepository) FindAll() data.RepositoryResult {
	var workouts entity.Workout

	err := r.db.Find(&workouts).Error

	if err != nil {
		return data.RepositoryResult{ErrorMessage: err}
	}

	return data.RepositoryResult{Result: &workouts}
}
func (r *WorkoutRepository) FindById(id uuid.UUID) data.RepositoryResult {
	var workout entity.Workout

	err := r.db.Where(&entity.Workout{ID: id}).Take(&workout).Error

	if err != nil {
		return data.RepositoryResult{ErrorMessage: err}
	}

	return data.RepositoryResult{Result: &workout}
}

func (r *WorkoutRepository) Save(workout *entity.Workout) data.RepositoryResult {
	err := r.db.Save(workout).Error

	if err != nil {
		return data.RepositoryResult{ErrorMessage: err}
	}

	return data.RepositoryResult{Result: &workout}
}

func (r *WorkoutRepository) DeleteById(id uuid.UUID) data.RepositoryResult {
	err := r.db.Delete(&entity.Workout{ID: id}).Error

	if err != nil {
		return data.RepositoryResult{ErrorMessage: err}
	}

	return data.RepositoryResult{Result: nil}
}

func (r *WorkoutRepository) DeleteByIds(ids *[]uuid.UUID) data.RepositoryResult {
	err := r.db.Where("ID IN (?)", *ids).Delete(entity.Workout{}).Error

	if err != nil {
		return data.RepositoryResult{ErrorMessage: err}
	}

	return data.RepositoryResult{Result: nil}
}

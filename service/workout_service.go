package service

import (
	"github.com/fitnessstack/web-service/database/repository"
	"github.com/fitnessstack/web-service/model/entity"
	"github.com/fitnessstack/web-service/model/response"
	"github.com/google/uuid"
	"log"
)

func Create(workout *entity.Workout, repository repository.WorkoutRepository) response.GeneralResponse {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	workout.ID = uuidResult

	operationResult := repository.Save(workout)

	// Return error
	if operationResult.ErrorMessage != nil {
		return response.GeneralResponse{
			Success: false,
			Message: operationResult.ErrorMessage.Error(),
		}
	}

	var data = operationResult.Result.(*entity.Workout)

	// Return Ok
	return response.GeneralResponse{
		Success: true,
		Data:    data,
	}
}

func FindAll(workoutRepository repository.WorkoutRepository) response.GeneralResponse {
	operationResult := workoutRepository.FindAll()

	// Return Error
	if operationResult.ErrorMessage != nil {
		return response.GeneralResponse{
			Success: false,
			Message: operationResult.ErrorMessage.Error(),
		}
	}

	var data = operationResult.Result.(*entity.Workout)

	// Return Ok
	return response.GeneralResponse{
		Success: true,
		Data:    data,
	}
}

func FindById(id uuid.UUID, repository repository.WorkoutRepository) response.GeneralResponse {
	operationResult := repository.FindById(id)

	// Return Error
	if operationResult.ErrorMessage != nil {
		return response.GeneralResponse{
			Success: false,
			Message: operationResult.ErrorMessage.Error(),
		}
	}

	var data = operationResult.Result.(*entity.Workout)

	// Return Ok
	return response.GeneralResponse{
		Success: true,
		Data:    data,
	}
}

func UpdateById(id uuid.UUID, workout *entity.Workout, repository repository.WorkoutRepository) response.GeneralResponse {
	existingResultResponse := FindById(id, repository)

	// Return Error not found
	if !existingResultResponse.Success {
		return existingResultResponse
	}

	existingResult := existingResultResponse.Data.(*entity.Workout)

	existingResult.Name = workout.Name

	operationResult := repository.Save(existingResult)

	// Return Error
	if operationResult.ErrorMessage != nil {
		return response.GeneralResponse{
			Success: false,
			Message: operationResult.ErrorMessage.Error(),
		}
	}

	// Return Ok
	return response.GeneralResponse{
		Success: true,
		Data:    operationResult.Result,
	}
}

func DeleteById(id uuid.UUID, repository repository.WorkoutRepository) response.GeneralResponse {
	operationResult := repository.DeleteById(id)

	// Return Error
	if operationResult.ErrorMessage != nil {
		return response.GeneralResponse{
			Success: false,
			Message: operationResult.ErrorMessage.Error(),
		}
	}

	// Return Ok
	return response.GeneralResponse{
		Success: true,
	}
}

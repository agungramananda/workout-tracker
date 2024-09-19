package domain

import "gorm.io/gorm"

type ExerciseCategoryMuscleGroup struct {
	ID 					uint           `json:"id"`
	Name        string        `json:"name"`
	Category    string      `json:"category"`
	MuscleGroup []string `json:"muscle_group"`
}

type ExerciseRepository interface {
	GetExercises() ([]ExerciseCategoryMuscleGroup, error)
	GetExerciseByID(id int) (ExerciseCategoryMuscleGroup, error)
}

type GormExerciseRepository struct {
	DB *gorm.DB
}


func NewGormExerciseRepository(db *gorm.DB) *GormExerciseRepository {
	return &GormExerciseRepository{
		DB: db,
	}
}

func (r *GormExerciseRepository) GetExercises() ([]ExerciseCategoryMuscleGroup, error) {
	var exercises []Exercise
	if err := r.DB.Preload("MuscleGroup").Find(&exercises).Error; err != nil {
			return nil, err
	}

	var result []ExerciseCategoryMuscleGroup
	for _, exercise := range exercises {
			var muscleGroups []string
			for _, mg := range exercise.MuscleGroup {
					muscleGroups = append(muscleGroups, mg.MuscleGroup)
			}
			result = append(result, ExerciseCategoryMuscleGroup{
					ID : 				 exercise.ID,
					Name:        exercise.Name,
					MuscleGroup: muscleGroups,
			})
	}
	return result, nil
}

func (r *GormExerciseRepository) GetExerciseByID(id int) (ExerciseCategoryMuscleGroup, error) {
	var exercise Exercise
	if err := r.DB.Preload("MuscleGroup").First(&exercise, id).Error; err != nil {
			return ExerciseCategoryMuscleGroup{}, err
	}

	var muscleGroups []string
	for _, mg := range exercise.MuscleGroup {
			muscleGroups = append(muscleGroups, mg.MuscleGroup)
	}

	result := ExerciseCategoryMuscleGroup{
			ID : 				 exercise.ID,
			Name:        exercise.Name,
			MuscleGroup: muscleGroups,
	}
	return result, nil
}
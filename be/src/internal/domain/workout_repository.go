package domain

import "gorm.io/gorm"

type WorkoutRepository interface {
	CreateWorkout(workout *Workout) error
	GetWorkouts(userId int) ([]Workout, error)
	GetWorkoutByID(id int,userId int) (Workout, error)
	UpdateWorkout(workout *Workout) error
	DeleteWorkout(id int) error
}

type GormWorkoutRepository struct {
	DB *gorm.DB
}

func NewGormWorkoutRepository(db *gorm.DB) *GormWorkoutRepository {
	return &GormWorkoutRepository{
		DB: db,
	}
}

func (r *GormWorkoutRepository) CreateWorkout(workout *Workout) error {
	var user User
	err := r.DB.Where("id = ?", workout.UserID).First(&user).Error
	if err != nil {
		return err
	}
	return r.DB.Create(workout).Error
}

func (r *GormWorkoutRepository) GetWorkouts(userId int) ([]Workout, error) {
	var workouts []Workout
	var user User
	err := r.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return workouts, err
	}
	err = r.DB.Preload("User").Preload("ExercisesPlan.Exercise").Preload("Comments").Where("user_id = ?", userId).Find(&workouts).Error
	return workouts, err
}

func (r *GormWorkoutRepository) GetWorkoutByID(id int, userId int) (Workout, error) {
	var workout Workout
	var user User
	err := r.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return workout, err
	}
	err = r.DB.Preload("User").Preload("ExercisesPlan.Exercise").Preload("Comments").Where("id = ? AND user_id = ?", id, userId).First(&workout).Error
	return workout, err
}

func (r *GormWorkoutRepository) UpdateWorkout(workout *Workout) error {
	var user User
	err := r.DB.Where("id = ?", workout.UserID).First(&user).Error
	if err != nil {
		return err
	}
	return r.DB.Save(workout).Error
}

func (r *GormWorkoutRepository) DeleteWorkout(id int) error {
	var workout Workout
	err := r.DB.Where("id = ?", id).First(&workout).Error
	if err != nil {
		return err
	}
	return r.DB.Delete(&Workout{}, id).Error
}
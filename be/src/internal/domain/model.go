package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	gorm.Model
}

type Workout struct {
	UserID      int        `json:"user_id"`
	User        User       `gorm:"constraint:OnDelete:CASCADE;"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Date        time.Time  `json:"date" gorm:"type:date"`
	Time        time.Time  `json:"time"`
	RestBetweenExercises int `json:"rest_between_exercises"` // in seconds
	IsCompleted bool       `json:"is_completed"`
	ExercisesPlan   []ExercisePlan `gorm:"constraint:OnDelete:CASCADE;" json:"exercises_plan"`
	Comments    []Comment  `gorm:"constraint:OnDelete:CASCADE;" json:"comments"`
	gorm.Model
}

type Exercise struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	MuscleGroup []*MuscleGroup `gorm:"many2many:exercise_muscle_groups;" json:"muscle_group"`
	gorm.Model
}

type MuscleGroup struct {
	MuscleGroup string `json:"muscle_group"`
	gorm.Model
}

type Comment struct {
	WorkoutID int    `json:"workout_id"`
	Workout   Workout `gorm:"constraint:OnDelete:CASCADE;"`
	UserID    int    `json:"user_id"`
	User      User   `gorm:"constraint:OnDelete:CASCADE;"`
	Comment   string `json:"comment"`
	gorm.Model
}

type ExercisePlan struct {
	WorkoutID   int       `json:"workout_id"`
	Workout     Workout   `gorm:"constraint:OnDelete:CASCADE;"`
	ExerciseID  int       `json:"exercise_id"`
	Exercise    Exercise  `gorm:"constraint:OnDelete:CASCADE;"`
	Sets        int      `json:"sets"`
	Reps        int      `json:"reps"`
	Weight      *int      `json:"weight"` // in kg
	RestBetweenSets    int      `json:"rest_time"` // in seconds
	Order       int      `json:"order"`
	IsCompleted bool     `json:"is_completed"`
	gorm.Model
}
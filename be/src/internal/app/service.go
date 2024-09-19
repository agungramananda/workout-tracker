package app

import (
	"errors"
	"workout-tracker/m/v0.0.0/src/internal/domain"
	"workout-tracker/m/v0.0.0/src/pkg/utils"
)

type ExerciseService struct {
	ExerciseRepository domain.ExerciseRepository
}

type UserService struct {
	UserRepository domain.UserRepository
}

type WorkoutService struct {
	WorkoutRepository domain.WorkoutRepository
}

func NewExerciseService(er domain.ExerciseRepository) *ExerciseService {
	return &ExerciseService{
		ExerciseRepository: er,
	}
}

func NewUserService(ur domain.UserRepository) *UserService {
	return &UserService{
		UserRepository: ur,
	}
}

func NewWorkoutService(wr domain.WorkoutRepository) *WorkoutService {
	return &WorkoutService{
		WorkoutRepository: wr,
	}
}

func (us *UserService) Login(username, password string) (string, error) {
	user, err := us.UserRepository.VerifyUser(username, password)
	if err != nil {
		return "",  err
	}
	accessToken, err := utils.CreateAccessToken(user.ID)
	if err != nil {
		return "",err
	}
	return accessToken,  nil
}

func (us *UserService) Register(user *domain.User) ( string, error) {
	hashedPassword, err := utils.HashPasword(user.Password)
	if err != nil {
		return "", err
	}

	newUser := domain.User{
		Username: user.Username,
		Password: hashedPassword,
		Fullname: user.Fullname,
		Email:    user.Email,
	}

	result, err := us.UserRepository.CreateUser(&newUser)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_users_username\" (SQLSTATE 23505)" {
			return "", errors.New("username already exists")
		}
		return "", err
	}
	accessToken, err := utils.CreateAccessToken(result.ID)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
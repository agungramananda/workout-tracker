package domain

import (
	"errors"

	"workout-tracker/m/v0.0.0/src/pkg/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]User, error)
	GetUserByID(id int) (Profile, error)
	CreateUser(user *User) (User, error)
	UpdateUser(user Profile) (Profile, error)
	VerifyUser(username, password string) (User, error)
}

type Profile struct {
	ID 		 int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		DB: db,
	}
}

func (r *GormUserRepository) GetUsers() ([]User, error){
	var users []User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepository) GetUserByID(id int) (Profile, error) {
	var user User
	if err := r.DB.First(&user, id).Error; err != nil {
		return Profile{}, err
	}
	result := Profile{
		ID:       int(user.ID),
		Username: user.Username,
		Fullname: user.Fullname,
		Email:    user.Email,
	}
	return result, nil
}

func (r *GormUserRepository) CreateUser(user *User) (User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return User{}, err
	}
	return *user, nil
}

func (r *GormUserRepository) UpdateUser(user Profile) (Profile, error) {
	if err := r.DB.Save(&user).Error; err != nil {
		return Profile{}, err
	}
	profile := Profile{
		Username: user.Username,
		Fullname: user.Fullname,
		Email:    user.Email,
	}
	return profile, nil
}

func (r *GormUserRepository) VerifyUser(username, password string) (User, error) {
	var user User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
			return User{}, err
	}
	err := utils.CheckPassword(user.Password, password)
	if err != nil {
			return User{}, errors.New("invalid password")
	}
	return user, nil
}
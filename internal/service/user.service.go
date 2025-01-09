package service

import (
	"src/internal/models"
	"src/internal/repo"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// Users trả về danh sách []models.User trực tiếp từ repo
func (us *UserService) Users() []models.User {
	return us.userRepo.GetUsers()
}

func (us *UserService) GetUserById(id int) models.User {
	return us.userRepo.GetUserById(id)
}

package repo

import "src/internal/models"

// UserRepo là repository để quản lý dữ liệu User
type UserRepo struct{}

// NewUserRepo tạo một instance mới của UserRepo
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

var userList = []models.User{
	{ID: 1, Username: "user1", Email: "user1@example.com"},
	{ID: 2, Username: "user2", Email: "user2@example.com"},
	{ID: 3, Username: "user3", Email: "user3@example.com"},
}

// GetUsers trả về danh sách các user mock
func (ur *UserRepo) GetUsers() []models.User {
	return userList
}

func (ur *UserRepo) GetUserById(id int) models.User {
	for _, user := range userList {
		if user.ID == id {
			return user
		}
	}
	return models.User{}
}

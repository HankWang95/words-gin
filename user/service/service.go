package service

import (
	"github.com/HankWang95/words-gin/user"
)

const (
	USER_STATUS_OK = 1000 * (iota + 1)
	USER_STATUS_INVALID
)

type UserService struct {
	repo UserRepository
}

type UserRepository interface {
	VerifyUserWithNamePassword(name, password string) (result *user.User, err error)
	CreateUser(name, password string) (result *user.User, err error)
}

func NewUserService(repo UserRepository) user.UserServeice {
	return &UserService{repo: repo}
}

func (this *UserService) Login(name, password string) (status int) {
	return

}
func (this *UserService) Logout() {
	return

}

func (this *UserService) SignUp(name, password string) (status int, err error) {
	return

}

package mysql

import (
	"github.com/HankWang95/words-gin/user"
	"github.com/HankWang95/words-gin/user/service"
	"github.com/smartwalle/dbs"
)

const (
	K_DB_USER = "user"
)

type userRepository struct {
	db dbs.DB
}



func NewUserRepository(db dbs.DB) service.UserRepository {
	return &userRepository{db: db}
}

// todo Database interaction

func (this *userRepository) VerifyUserWithNamePassword(name string, password string) (result *user.User, err error) {
	return
}

func (this *userRepository) CreateUser(name string, password string) (result *user.User, err error) {
	return
}

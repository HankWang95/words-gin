package user

type User struct {
	Id        int64  `json:"id"          sql:"id"`
	Name      string `json:"name"        sql:"name"`
	Password  string `json:"password"    sql:"password"`
	LastLogin string `json:"last_login"  sql:"last_login"`
}

type UserServeice interface {
	Login(name, password string) (status int)
	Logout()
	SignUp(name, password string) (status int, err error)
}

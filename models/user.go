package models

type UserLoginForm struct {
	UsernameOrEmail string
	Password        string
}

func (loginForm *UserLoginForm) login() {

}

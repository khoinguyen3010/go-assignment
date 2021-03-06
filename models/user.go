package models

import "time"

type User struct {
	ID             string    `json:"id"`
	Firstname      string    `json:"first_name" form:"first_name"`
	Lastname       string    `json:"last_name" form:"last_name"`
	Alias          string    `json:"alias" form:"alias"`
	Email          string    `json:"email" form:"email"`
	Username       string    `json:"username" form:"username"`
	Password       string    `json:"password" form:"password"`
	HashedPassword string    `json:"hashed_password" form:"hashed_password"`
	DateofBirth    time.Time `json:"date_of_birth" form:"date_of_birth"`
	BaseObject
}

type UserSignUpForm struct {
	Firstname   string    `json:"first_name" form:"first_name"`
	Lastname    string    `json:"last_name" form:"last_name"`
	Email       string    `json:"email" form:"email"`
	Username    string    `json:"username" form:"username"`
	Password    string    `json:"password" form:"password"`
	Alias       string    `json:"alias" form:"alias"`
	DateofBirth time.Time `json:"date_of_birth" form:"date_of_birth"`
}

type UserLoginForm struct {
	UsernameOrEmail string `json:"credentials" form:"credentials"`
	Password        string `json:"password" form:"password"`
}

type UserSignupResponse struct {
	Firstname      string `json:"first_name"`
	Lastname       string `json:"last_name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	BaseObject
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

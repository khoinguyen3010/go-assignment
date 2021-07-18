package authentication

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/khoinguyen3010/go-assignment/models"
	"github.com/khoinguyen3010/go-assignment/utils"
	"github.com/labstack/echo/v4"
)

func Login(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	secret := os.Getenv("JWT_SECRET")

	// TODO: query from database for user instead of hard code
	if username != "Khoi Nguyen" || password != "Minh@nh0806" {
		return echo.ErrUnauthorized
	}

	claims := &models.JwtCustomClaims{
		Name:  "Khoi Nguyen",
		UUID:  "9E98C454-C7AC-4330-B2EF-983765E00547",
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, &models.UserLoginResponse{Token: t})
}

func Restricted(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)
	name := claims.Name
	return ctx.JSON(http.StatusOK, &models.BaseSuccessResponse{Message: "Restricted success! Welcome " + name})
}

func Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &models.BaseSuccessResponse{Message: "Hello, world!"})
}

func SignUp(ctx echo.Context) error {
	// Handling inputs: Bind is used to easily bind data from request to struct (form or json)
	signUpForm := &models.UserSignUpForm{}
	if err := ctx.Bind(signUpForm); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	hashed_password, err := utils.HashPassword(signUpForm.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	// Business logic
	created_time := time.Now()
	current_user := "Ngao" //Todo: change this to get current user from context

	// Returning API's response
	response := &models.UserSignupResponse{
		Firstname:      signUpForm.Firstname,
		Lastname:       signUpForm.Lastname,
		Email:          signUpForm.Email,
		HashedPassword: hashed_password,
		BaseObject: models.BaseObject{
			CreatedAt: created_time,
			UpdatedAt: created_time,
			CreatedBy: current_user,
			UpdatedBy: current_user,
		},
	}
	return ctx.JSON(http.StatusCreated, response)
}

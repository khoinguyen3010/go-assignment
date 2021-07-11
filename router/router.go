package router

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/khoinguyen3010/go-assignment/authentication"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	app := echo.New()
	rapp := app.Group("/restricted")

	// Application's Configurations
	logConfig := middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} | method: ${method} | endpoint: ${uri} | status: ${status} | agent: ${user_agent}\n",
	}
	jwtConfig := middleware.JWTConfig{
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		Claims:        &authentication.JwtCustomClaims{},
		AuthScheme:    "Bearer",
		SigningMethod: middleware.AlgorithmHS256,
		TokenLookup:   "header:" + echo.HeaderAuthorization,
	}
	corsConfig := middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodDelete, http.MethodPut, http.MethodPost},
	}

	// Normal Routes Middlewares
	app.Use(middleware.Recover())
	app.Use(middleware.LoggerWithConfig(logConfig))
	app.Use(middleware.CORSWithConfig(corsConfig))

	// Restricted Routes Middlewares
	rapp.Use(middleware.JWTWithConfig(jwtConfig))

	// Routes
	app.GET("/", hello)
	app.POST("/login", login)
	rapp.GET("", restricted)

	return app
}

func hello(ctx echo.Context) error {
	var res struct {
		Message string `json:"message"`
	}
	res.Message = "Hello, world!"
	return ctx.JSON(http.StatusOK, res)
}

func restricted(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*authentication.JwtCustomClaims)
	name := claims.Name
	var res struct {
		Message string `json:"message"`
	}
	res.Message = "Restricted success! Welcome " + name
	return ctx.JSON(http.StatusOK, res)
}

func login(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	secret := os.Getenv("JWT_SECRET")

	if username != "Khoi Nguyen" || password != "Minh@nh0806" {
		return echo.ErrUnauthorized
	}

	claims := &authentication.JwtCustomClaims{
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
		fmt.Println(err)
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

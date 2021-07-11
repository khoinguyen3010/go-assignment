package router

import (
	"net/http"
	"os"

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
	app.GET("/", authentication.Index)
	app.POST("/login", authentication.Login)
	rapp.GET("", authentication.Restricted)

	return app
}

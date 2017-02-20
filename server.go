package main

import (
	"net/http"
	"strings"

	"github.com/gouthamve/dredd-api/routes"
	"github.com/gouthamve/dredd-api/routes/mware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func main() {
	// setup
	viper.SetEnvPrefix("DREDD")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// user routes
	e.POST("/users", routes.SaveUser)
	e.GET("/user", routes.GetUser, mware.Auth)

	// challenge routes
	e.POST("/challenges", routes.SaveChallenge)
	e.GET("/challenges/:id", routes.GetChallenge)

	// submission routes
	e.POST("/challenges/:id/submit", routes.SaveSubmission, mware.Auth)
	e.Logger.Fatal(e.Start(":1323"))
}

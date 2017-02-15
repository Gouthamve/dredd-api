package main

import (
	"net/http"

	"github.com/gouthamve/gopherhack/routes"
	"github.com/gouthamve/gopherhack/routes/mware"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// user routes
	e.POST("/users", routes.SaveUser)
	e.GET("/user", routes.GetUser, mware.Auth)

	// challenge routes
	e.POST("/challenges", routes.SaveChallenge)
	e.GET("/challenges", routes.GetChallenge)
	e.Logger.Fatal(e.Start(":1323"))
}

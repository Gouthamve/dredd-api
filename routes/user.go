package routes

import (
	"net/http"

	"github.com/gouthamve/gopherhack/db"
	"github.com/labstack/echo"
)

// SaveUser saves the user
// POST /users
func SaveUser(c echo.Context) error {
	u := db.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}

	if db.Conn.NewRecord(u) {
		if err := db.Conn.Create(&u).Error; err != nil {
			// TODO: Better error reporting
			return c.JSON(http.StatusConflict, err.Error())
		}

		return c.JSON(http.StatusOK, u)
	}

	return echo.NewHTTPError(http.StatusConflict)
}

// GetUser gets the user
func GetUser(c echo.Context) error {
	return nil
}

// UpdateUser updates the user
func UpdateUser(c echo.Context) error {
	return nil
}

// DeleteUser deletes the user
func DeleteUser(c echo.Context) error {
	return nil
}

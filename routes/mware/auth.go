package mware

import (
	"net/http"

	"github.com/gouthamve/gopherhack/db"
	"github.com/labstack/echo"
)

// Auth is the authentication middleware
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get(echo.HeaderAuthorization)
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "No token provided")
		}

		// TODO: Use actual JWT
		userID, err := getUser(token)
		if err != nil {
			return err
		}

		count := 0
		db.Conn.Model(&db.User{}).Where("id = ?", userID).Count(&count)
		if count != 1 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("user", userID)
		return next(c)
	}
}

func getUser(token string) (string, error) {
	return token, nil
}

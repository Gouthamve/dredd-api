package routes

import (
	"net/http"

	"github.com/gouthamve/gopherhack/db"
	"github.com/labstack/echo"
)

// SaveChallenge saves the challenge
func SaveChallenge(c echo.Context) error {
	ch := db.Challenge{}
	if err := c.Bind(&ch); err != nil {
		return err
	}

	if db.Conn.NewRecord(ch) {
		if err := db.Conn.Create(&ch).Error; err != nil {
			return c.JSON(http.StatusConflict, err.Error())
		}

		if err := db.Conn.Save(&ch).Error; err != nil {
			return c.JSON(http.StatusConflict, err.Error())
		}

		return c.JSON(http.StatusOK, ch)
	}

	return echo.NewHTTPError(http.StatusConflict)
}

// GetChallenge returns the challenge
func GetChallenge(c echo.Context) error {
	ch := db.Challenge{
		Testcases: make([]db.Testcase, 0),
	}
	if err := db.Conn.First(&ch).Error; err != nil {
		return err
	}

	db.Conn.Model(&ch).Related(&ch.Testcases).Related(&ch.Limits)

	return c.JSON(200, ch)
}

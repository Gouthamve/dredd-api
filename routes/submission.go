package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gouthamve/gopherhack/db"
	"github.com/gouthamve/gopherhack/lib/fileserver"
	"github.com/juju/errors"
	"github.com/labstack/echo"
)

// SaveSubmission saves the submission
func SaveSubmission(c echo.Context) error {
	fh, err := c.FormFile("file")
	if err != nil {
		err := errors.Annotate(err, "submission: file is not sent")
		log.Println(errors.ErrorStack(err))
		return err
	}

	file, err := fh.Open()
	if err != nil {
		err := errors.Annotate(err, "submission: something is wrong with the file")
		log.Println(errors.ErrorStack(err))
		return err
	}
	defer file.Close()

	userID := c.Get("user").(string)
	userUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		err := errors.Annotate(err, "couldn't parse userID")
		log.Println(errors.ErrorStack(err))
		return err
	}

	obj, err := fileserver.SaveFile(file, userID)
	if err != nil {
		err := errors.Annotate(err, "couldn't save file")
		log.Println(errors.ErrorStack(err))
		return err
	}

	submission := db.Submission{
		FileURL: obj,
		Status:  "wut?",

		UserID: uint(userUint),
	}

	if err := db.Conn.Create(&submission).Error; err != nil {
		err := errors.Annotate(err, "couldn't save submission")
		log.Println(errors.ErrorStack(err))
		return err
	}

	return c.JSON(http.StatusOK, submission)
}

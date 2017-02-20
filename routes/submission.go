package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gouthamve/dredd"
	"github.com/gouthamve/dredd-api/db"
	"github.com/gouthamve/dredd-api/lib/fileserver"
	"github.com/gouthamve/dredd-api/lib/functions"
	"github.com/gouthamve/dredd/judge"
	"github.com/jinzhu/gorm"
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

	challengeID := c.Param("id")
	challengeIDUint, err := strconv.ParseUint(challengeID, 10, 64)
	if err != nil {
		err := errors.Annotate(err, "couldn't parse challengeID")
		log.Println(errors.ErrorStack(err))
		return err
	}

	challenge := db.Challenge{
		Testcases: make([]db.Testcase, 0),
	}
	if err := db.Conn.First(&challenge, challengeIDUint).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.NoContent(http.StatusNotFound)
		}

		err := errors.Annotate(err, "couldn't find the challenge")
		log.Println(errors.ErrorStack(err))
		return err
	}

	// TODO: Decide if error needs to be checked
	db.Conn.Model(&challenge).Related(&challenge.Limits).
		Related(&challenge.Testcases)

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

		UserID:      uint(userUint),
		ChallengeID: uint(challengeIDUint),
	}

	if err := db.Conn.Create(&submission).Error; err != nil {
		err := errors.Annotate(err, "couldn't save submission")
		log.Println(errors.ErrorStack(err))
		return err
	}

	// TODO: Sync waits on the execution. That is one goroutine per req.
	// Fix it by using pub-sub and Minio events
	runArgs := judge.RunnerArgs{
		Problem:  probFromChal(challenge),
		Filename: submission.FileURL,
	}

	results, err := functions.ExecuteSubmission(runArgs)
	if err != nil {
		err := errors.Annotate(err, "couldn't save submission")
		log.Println(errors.ErrorStack(err))
		return err
	}

	return c.JSON(http.StatusOK, results)
}

func probFromChal(c db.Challenge) dredd.Problem {
	ts := make([]dredd.Testcase, len(c.Testcases))
	for i, t := range c.Testcases {
		ts[i] = dredd.Testcase{
			Inp:      t.Inp,
			Expected: t.Expected,
		}
	}

	return dredd.Problem{
		Lang: "go",
		Limits: dredd.Limits{
			Memory: c.Limits.Memory,
			Time:   c.Limits.Time,
		},
		Testcases: ts,
	}
}

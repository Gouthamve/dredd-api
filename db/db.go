package db

import "github.com/jinzhu/gorm"

import _ "github.com/jinzhu/gorm/dialects/postgres" // pg driver

// Conn is the database connection
var Conn *gorm.DB

func init() {
	var err error

	Conn, err = gorm.Open("postgres", "host=localhost user=goutham dbname=dredd sslmode=disable password=getten1*")
	if err != nil {
		panic(err)
	}

	Conn.CreateTable(
		&Challenge{},
		&Submission{},
		&User{},
		&Limits{},
		&Testcase{},
	)
}

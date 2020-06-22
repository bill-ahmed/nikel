package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/nikel-api/nikel/common"
	"io/ioutil"
	"os"
)

func main() {
	os.Remove("database/database.db")
	db, _ := gorm.Open("sqlite3", "database/database.db")

	bytes, _ := ioutil.ReadFile(common.COURSEPATH)
	var courses []common.Course
	json.Unmarshal(bytes, &courses)
	db.CreateTable(&common.Course{})
	for _, course := range courses {
		db.Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.FOODPATH)
	var food []common.Food
	json.Unmarshal(bytes, &food)
	db.CreateTable(&common.Food{})
	for _, course := range food {
		db.Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.EXAMSPATH)
	var exams []common.Exam
	json.Unmarshal(bytes, &exams)
	db.CreateTable(&common.Exam{})
	for _, course := range exams {
		db.Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.EVALSPATH)
	var evals []common.Eval
	json.Unmarshal(bytes, &evals)
	db.CreateTable(&common.Eval{})
	for _, course := range evals {
		db.Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.SERVICESPATH)
	var services []common.Service
	json.Unmarshal(bytes, &services)
	db.CreateTable(&common.Service{})
	for _, course := range services {
		db.Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.TEXTBOOKPATH)
	var textbooks []common.Textbook
	json.Unmarshal(bytes, &textbooks)
	db.CreateTable(&common.Textbook{})
	for _, course := range textbooks {
		db.Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.BUILDINGSPATH)
	var buildings []common.Building
	json.Unmarshal(bytes, &buildings)
	db.CreateTable(&common.Building{})
	for _, course := range buildings {
		db.Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.PARKINGPATH)
	var parking []common.Parking
	json.Unmarshal(bytes, &parking)
	db.CreateTable(&common.Parking{})
	for _, course := range parking {
		db.Create(&course)
	}
}

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
	db.Table("courses").CreateTable(&common.Course{})
	for _, course := range courses {
		db.Table("courses").Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.FOODPATH)
	var food []common.Food
	json.Unmarshal(bytes, &food)
	db.Table("food").CreateTable(&common.Food{})
	for _, course := range food {
		db.Table("food").Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.EXAMSPATH)
	var exams []common.Exam
	json.Unmarshal(bytes, &exams)
	db.Table("exams").CreateTable(&common.Exam{})
	for _, course := range exams {
		db.Table("exams").Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.EVALSPATH)
	var evals []common.Eval
	json.Unmarshal(bytes, &evals)
	db.Table("evals").CreateTable(&common.Eval{})
	for _, course := range evals {
		db.Table("evals").Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.SERVICESPATH)
	var services []common.Service
	json.Unmarshal(bytes, &services)
	db.Table("services").CreateTable(&common.Service{})
	for _, course := range services {
		db.Table("services").Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.TEXTBOOKPATH)
	var textbooks []common.Textbook
	json.Unmarshal(bytes, &textbooks)
	db.Table("textbooks").CreateTable(&common.Textbook{})
	for _, course := range textbooks {
		db.Table("textbooks").Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.BUILDINGSPATH)
	var buildings []common.Building
	json.Unmarshal(bytes, &buildings)
	db.Table("buildings").CreateTable(&common.Building{})
	for _, course := range buildings {
		db.Table("buildings").Create(&course)
	}

	bytes, _ = ioutil.ReadFile(common.PARKINGPATH)
	var parking []common.Parking
	json.Unmarshal(bytes, &parking)
	db.Table("parking").CreateTable(&common.Parking{})
	for _, course := range parking {
		db.Table("parking").Create(&course)
	}
}

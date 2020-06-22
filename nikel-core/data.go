package main

import (
	"github.com/a8m/rql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/nikel-api/nikel/common"
	"github.com/thedevsaddam/gojsonq/v2"
	"os"
	"path/filepath"
)

// Database stores Nikel's data
type Database struct {
	Courses   *gojsonq.JSONQ
	Textbooks *gojsonq.JSONQ
	Buildings *gojsonq.JSONQ
	Food      *gojsonq.JSONQ
	Parking   *gojsonq.JSONQ
	Services  *gojsonq.JSONQ
	Exams     *gojsonq.JSONQ
	Evals     *gojsonq.JSONQ
}

type Parsers struct {
	Courses   *rql.Parser
	Textbooks *rql.Parser
	Buildings *rql.Parser
	Food      *rql.Parser
	Parking   *rql.Parser
	Services  *rql.Parser
	Exams     *rql.Parser
	Evals     *rql.Parser
}

var database = &Database{}
var parsers = &Parsers{}
var db *gorm.DB

// loadVals loads JSON data to database
func loadVals() {
	pathPrefix := ""
	wd, _ := os.Getwd()
	if filepath.Base(wd) == "nikel-core" {
		pathPrefix = "../"
	}

	db, _ = gorm.Open("sqlite3", pathPrefix+"database/database.db")

	database.Courses = gojsonq.New().File(pathPrefix + common.COURSEPATH).Reset()
	database.Textbooks = gojsonq.New().File(pathPrefix + common.TEXTBOOKPATH).Reset()
	database.Buildings = gojsonq.New().File(pathPrefix + common.BUILDINGSPATH).Reset()
	database.Food = gojsonq.New().File(pathPrefix + common.FOODPATH).Reset()
	database.Parking = gojsonq.New().File(pathPrefix + common.PARKINGPATH).Reset()
	database.Services = gojsonq.New().File(pathPrefix + common.SERVICESPATH).Reset()
	database.Exams = gojsonq.New().File(pathPrefix + common.EXAMSPATH).Reset()
	database.Evals = gojsonq.New().File(pathPrefix + common.EVALSPATH).Reset()

	parsers.Courses = rql.MustNewParser(rql.Config{
		Model:        common.Course{},
		DefaultLimit: common.DEFAULTLIMIT,
		FieldSep:     ".",
	})
	parsers.Textbooks = rql.MustNewParser(rql.Config{
		Model:        common.Textbook{},
		DefaultLimit: common.DEFAULTLIMIT,
		FieldSep:     ".",
	})
	parsers.Buildings = rql.MustNewParser(rql.Config{
		Model:        common.Building{},
		DefaultLimit: common.DEFAULTLIMIT,
		FieldSep:     ".",
	})
	parsers.Food = rql.MustNewParser(rql.Config{
		Model:        common.Food{},
		DefaultLimit: common.DEFAULTLIMIT,
		FieldSep:     ".",
	})
	parsers.Parking = rql.MustNewParser(rql.Config{
		Model:        common.Parking{},
		DefaultLimit: common.DEFAULTLIMIT,
		FieldSep:     ".",
	})
	parsers.Services = rql.MustNewParser(rql.Config{
		Model:        common.Service{},
		DefaultLimit: common.DEFAULTLIMIT,
		FieldSep:     ".",
	})
	parsers.Exams = rql.MustNewParser(rql.Config{
		Model:        common.Exam{},
		DefaultLimit: common.DEFAULTLIMIT,
		FieldSep:     ".",
	})
	parsers.Evals = rql.MustNewParser(rql.Config{
		Model:        common.Eval{},
		DefaultLimit: common.DEFAULTLIMIT,
		FieldSep:     ".",
	})
}

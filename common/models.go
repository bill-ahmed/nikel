package common

import (
	"database/sql"
)

// Course represents a course item
type Course struct {
	ID                         sql.NullString `json:"id" gorm:"PRIMARY_KEY" rql:"filter"`
	Code                       sql.NullString `json:"code" rql:"filter"`
	Name                       sql.NullString `json:"name" rql:"filter"`
	Description                sql.NullString `json:"description" rql:"filter"`
	Division                   sql.NullString `json:"division" rql:"filter"`
	Department                 sql.NullString `json:"department" rql:"filter"`
	Prerequisites              sql.NullString `json:"prerequisites" rql:"filter"`
	Corequisites               sql.NullString `json:"corequisites" rql:"filter"`
	Exclusions                 sql.NullString `json:"exclusions" rql:"filter"`
	RecommendedPreparation     sql.NullString `json:"recommended_preparation" rql:"filter"`
	Level                      sql.NullString `json:"level" rql:"filter"`
	Campus                     sql.NullString `json:"campus" rql:"filter"`
	Term                       sql.NullString `json:"term" rql:"filter"`
	ArtsAndScienceBreadth      sql.NullString `json:"arts_and_science_breadth" rql:"filter"`
	ArtsAndScienceDistribution sql.NullString `json:"arts_and_science_distribution" rql:"filter"`
	UtmDistribution            sql.NullString `json:"utm_distribution" rql:"filter"`
	UtscBreadth                sql.NullString `json:"utsc_breadth" rql:"filter"`
	ApscElectives              sql.NullString `json:"apsc_electives" rql:"filter"`
	MeetingSections            []struct {
		Code        sql.NullString   `json:"code"`
		Instructors []sql.NullString `json:"instructors"`
		Times       []struct {
			Day      sql.NullString `json:"day"`
			Start    sql.NullInt64  `json:"start"`
			End      sql.NullInt64  `json:"end"`
			Duration sql.NullInt64  `json:"duration"`
			Location sql.NullString `json:"location"`
		} `json:"times"`
		Size           sql.NullInt64  `json:"size"`
		Enrollment     sql.NullInt64  `json:"enrollment"`
		WaitlistOption sql.NullBool   `json:"waitlist_option"`
		Delivery       sql.NullString `json:"delivery"`
	} `json:"meeting_sections"`
	LastUpdated sql.NullString `json:"last_updated"`
}

// Textbook represents a textbook item
type Textbook struct {
	ID      sql.NullString  `json:"id" rql:"filter"`
	Isbn    sql.NullString  `json:"isbn" rql:"filter"`
	Title   sql.NullString  `json:"title" rql:"filter"`
	Edition sql.NullInt64   `json:"edition" rql:"filter"`
	Author  sql.NullString  `json:"author" rql:"filter"`
	Image   sql.NullString  `json:"image" rql:"filter"`
	Price   sql.NullFloat64 `json:"price" rql:"filter"`
	URL     sql.NullString  `json:"url" rql:"filter"`
	Courses []struct {
		ID              sql.NullString `json:"id"`
		Code            sql.NullString `json:"code"`
		Requirement     sql.NullString `json:"requirement"`
		MeetingSections []struct {
			Code        sql.NullString   `json:"code"`
			Instructors []sql.NullString `json:"instructors"`
		} `json:"meeting_sections"`
	} `json:"courses"`
	LastUpdated sql.NullString `json:"last_updated"`
}

// Building represents a building item
type Building struct {
	ID        sql.NullString `json:"id" rql:"filter"`
	Code      sql.NullString `json:"code" rql:"filter"`
	Tags      sql.NullString `json:"tags" rql:"filter"`
	Name      sql.NullString `json:"name" rql:"filter"`
	ShortName sql.NullString `json:"short_name" rql:"filter"`
	Address   struct {
		Street   sql.NullString `json:"street"`
		City     sql.NullString `json:"city"`
		Province sql.NullString `json:"province"`
		Country  sql.NullString `json:"country"`
		Postal   sql.NullString `json:"postal"`
	} `json:"address"`
	Coordinates struct {
		Latitude  sql.NullFloat64 `json:"latitude"`
		Longitude sql.NullFloat64 `json:"longitude"`
	} `json:"coordinates"`
	LastUpdated sql.NullString `json:"last_updated"`
}

// Food represents a food item
type Food struct {
	ID          sql.NullString `json:"id" rql:"filter"`
	Name        sql.NullString `json:"name" rql:"filter"`
	Description sql.NullString `json:"description" rql:"filter"`
	Tags        sql.NullString `json:"tags" rql:"filter"`
	Campus      sql.NullString `json:"campus" rql:"filter"`
	Address     sql.NullString `json:"address" rql:"filter"`
	Coordinates struct {
		Latitude  sql.NullFloat64 `json:"latitude"`
		Longitude sql.NullFloat64 `json:"longitude"`
	} `json:"coordinates"`
	Hours struct {
		Sunday struct {
			Closed sql.NullBool  `json:"closed"`
			Open   sql.NullInt64 `json:"open"`
			Close  sql.NullInt64 `json:"close"`
		} `json:"sunday"`
		Monday struct {
			Closed sql.NullBool  `json:"closed"`
			Open   sql.NullInt64 `json:"open"`
			Close  sql.NullInt64 `json:"close"`
		} `json:"monday"`
		Tuesday struct {
			Closed sql.NullBool  `json:"closed"`
			Open   sql.NullInt64 `json:"open"`
			Close  sql.NullInt64 `json:"close"`
		} `json:"tuesday"`
		Wednesday struct {
			Closed sql.NullBool  `json:"closed"`
			Open   sql.NullInt64 `json:"open"`
			Close  sql.NullInt64 `json:"close"`
		} `json:"wednesday"`
		Thursday struct {
			Closed sql.NullBool  `json:"closed"`
			Open   sql.NullInt64 `json:"open"`
			Close  sql.NullInt64 `json:"close"`
		} `json:"thursday"`
		Friday struct {
			Closed sql.NullBool  `json:"closed"`
			Open   sql.NullInt64 `json:"open"`
			Close  sql.NullInt64 `json:"close"`
		} `json:"friday"`
		Saturday struct {
			Closed sql.NullBool  `json:"closed"`
			Open   sql.NullInt64 `json:"open"`
			Close  sql.NullInt64 `json:"close"`
		} `json:"saturday"`
	} `json:"hours"`
	Image       sql.NullString   `json:"image" rql:"filter"`
	URL         sql.NullString   `json:"url" rql:"filter"`
	Twitter     sql.NullString   `json:"twitter" rql:"filter"`
	Facebook    sql.NullString   `json:"facebook" rql:"filter"`
	Attributes  []sql.NullString `json:"attributes"`
	LastUpdated sql.NullString   `json:"last_updated"`
}

// Parking represents a parking item
type Parking struct {
	ID          sql.NullString `json:"id" rql:"filter"`
	Name        sql.NullString `json:"name" rql:"filter"`
	Alias       sql.NullString `json:"alias" rql:"filter"`
	BuildingID  sql.NullString `json:"building_id" rql:"filter"`
	Description sql.NullString `json:"description" rql:"filter"`
	Campus      sql.NullString `json:"campus" rql:"filter"`
	Address     sql.NullString `json:"address" rql:"filter"`
	Coordinates struct {
		Latitude  sql.NullFloat64 `json:"latitude"`
		Longitude sql.NullFloat64 `json:"longitude"`
	} `json:"coordinates"`
	LastUpdated sql.NullString `json:"last_updated"`
}

// Service represents an service item
type Service struct {
	ID          sql.NullString `json:"id" rql:"filter"`
	Name        sql.NullString `json:"name" rql:"filter"`
	Alias       sql.NullString `json:"alias" rql:"filter"`
	BuildingID  sql.NullString `json:"building_id" rql:"filter"`
	Description sql.NullString `json:"description" rql:"filter"`
	Campus      sql.NullString `json:"campus" rql:"filter"`
	Address     sql.NullString `json:"address" rql:"filter"`
	Image       sql.NullString `json:"image" rql:"filter"`
	Coordinates struct {
		Latitude  sql.NullFloat64 `json:"latitude"`
		Longitude sql.NullFloat64 `json:"longitude"`
	} `json:"coordinates"`
	Tags        sql.NullString   `json:"tags" rql:"filter"`
	Attributes  []sql.NullString `json:"attributes"`
	LastUpdated sql.NullString   `json:"last_updated"`
}

// Exam represents an exam item
type Exam struct {
	ID         sql.NullString `json:"id" rql:"filter"`
	CourseID   sql.NullString `json:"course_id" rql:"filter"`
	CourseCode sql.NullString `json:"course_code" rql:"filter"`
	Campus     sql.NullString `json:"campus" rql:"filter"`
	Date       sql.NullString `json:"date" rql:"filter"`
	Start      sql.NullInt64  `json:"start" rql:"filter"`
	End        sql.NullInt64  `json:"end" rql:"filter"`
	Duration   sql.NullInt64  `json:"duration" rql:"filter"`
	Sections   []struct {
		LectureCode sql.NullString `json:"lecture_code"`
		Split       sql.NullString `json:"split"`
		Location    sql.NullString `json:"location"`
	} `json:"sections"`
	LastUpdated sql.NullString `json:"last_updated"`
}

// Eval represents an eval item
type Eval struct {
	ID     sql.NullString `json:"id" rql:"filter"`
	Name   sql.NullString `json:"name" rql:"filter"`
	Campus sql.NullString `json:"campus" rql:"filter"`
	Terms  []struct {
		Term     sql.NullString `json:"term"`
		Lectures []struct {
			LectureCode sql.NullString  `json:"lecture_code"`
			Firstname   sql.NullString  `json:"firstname"`
			Lastname    sql.NullString  `json:"lastname"`
			S1          sql.NullFloat64 `json:"s1"`
			S2          sql.NullFloat64 `json:"s2"`
			S3          sql.NullFloat64 `json:"s3"`
			S4          sql.NullFloat64 `json:"s4"`
			S5          sql.NullFloat64 `json:"s5"`
			S6          sql.NullFloat64 `json:"s6"`
			Invited     sql.NullInt64   `json:"invited"`
			Responses   sql.NullInt64   `json:"responses"`
		} `json:"lectures"`
	} `json:"terms"`
	LastUpdated sql.NullString `json:"last_updated"`
}

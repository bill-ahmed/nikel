package common

import "gopkg.in/guregu/null.v4"

// Course represents a course item
type Course struct {
	ID                         null.String `json:"id" gorm:"PRIMARY_KEY" rql:"filter"`
	Code                       null.String `json:"code" rql:"filter"`
	Name                       null.String `json:"name" rql:"filter"`
	Description                null.String `json:"description" rql:"filter"`
	Division                   null.String `json:"division" rql:"filter"`
	Department                 null.String `json:"department" rql:"filter"`
	Prerequisites              null.String `json:"prerequisites" rql:"filter"`
	Corequisites               null.String `json:"corequisites" rql:"filter"`
	Exclusions                 null.String `json:"exclusions" rql:"filter"`
	RecommendedPreparation     null.String `json:"recommended_preparation" rql:"filter"`
	Level                      null.String `json:"level" rql:"filter"`
	Campus                     null.String `json:"campus" rql:"filter"`
	Term                       null.String `json:"term" rql:"filter"`
	ArtsAndScienceBreadth      null.String `json:"arts_and_science_breadth" rql:"filter"`
	ArtsAndScienceDistribution null.String `json:"arts_and_science_distribution" rql:"filter"`
	UtmDistribution            null.String `json:"utm_distribution" rql:"filter"`
	UtscBreadth                null.String `json:"utsc_breadth" rql:"filter"`
	ApscElectives              null.String `json:"apsc_electives" rql:"filter"`
	MeetingSections            []struct {
		Code        null.String   `json:"code"`
		Instructors []null.String `json:"instructors"`
		Times       []struct {
			Day      null.String `json:"day"`
			Start    null.Int    `json:"start"`
			End      null.Int    `json:"end"`
			Duration null.Int    `json:"duration"`
			Location null.String `json:"location"`
		} `json:"times"`
		Size           null.Int    `json:"size"`
		Enrollment     null.Int    `json:"enrollment"`
		WaitlistOption null.Bool   `json:"waitlist_option"`
		Delivery       null.String `json:"delivery"`
	} `json:"meeting_sections"`
	LastUpdated null.String `json:"last_updated" rql:"filter"`
}

// Textbook represents a textbook item
type Textbook struct {
	ID      null.String `json:"id" rql:"filter"`
	Isbn    null.String `json:"isbn" rql:"filter"`
	Title   null.String `json:"title" rql:"filter"`
	Edition null.Int    `json:"edition" rql:"filter"`
	Author  null.String `json:"author" rql:"filter"`
	Image   null.String `json:"image" rql:"filter"`
	Price   null.Float  `json:"price" rql:"filter"`
	URL     null.String `json:"url" rql:"filter"`
	Courses []struct {
		ID              null.String `json:"id"`
		Code            null.String `json:"code"`
		Requirement     null.String `json:"requirement"`
		MeetingSections []struct {
			Code        null.String   `json:"code"`
			Instructors []null.String `json:"instructors"`
		} `json:"meeting_sections"`
	} `json:"courses"`
	LastUpdated null.String `json:"last_updated" rql:"filter"`
}

// Building represents a building item
type Building struct {
	ID        null.String `json:"id" rql:"filter"`
	Code      null.String `json:"code" rql:"filter"`
	Tags      null.String `json:"tags" rql:"filter"`
	Name      null.String `json:"name" rql:"filter"`
	ShortName null.String `json:"short_name" rql:"filter"`
	Address   struct {
		Street   null.String `json:"street"`
		City     null.String `json:"city"`
		Province null.String `json:"province"`
		Country  null.String `json:"country"`
		Postal   null.String `json:"postal"`
	} `json:"address"`
	Coordinates struct {
		Latitude  null.Float `json:"latitude"`
		Longitude null.Float `json:"longitude"`
	} `json:"coordinates"`
	LastUpdated null.String `json:"last_updated" rql:"filter"`
}

// Food represents a food item
type Food struct {
	ID          null.String `json:"id" rql:"filter"`
	Name        null.String `json:"name" rql:"filter"`
	Description null.String `json:"description" rql:"filter"`
	Tags        null.String `json:"tags" rql:"filter"`
	Campus      null.String `json:"campus" rql:"filter"`
	Address     null.String `json:"address" rql:"filter"`
	Coordinates struct {
		Latitude  null.Float `json:"latitude"`
		Longitude null.Float `json:"longitude"`
	} `json:"coordinates"`
	Hours struct {
		Sunday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"sunday"`
		Monday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"monday"`
		Tuesday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"tuesday"`
		Wednesday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"wednesday"`
		Thursday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"thursday"`
		Friday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"friday"`
		Saturday struct {
			Closed null.Bool `json:"closed"`
			Open   null.Int  `json:"open"`
			Close  null.Int  `json:"close"`
		} `json:"saturday"`
	} `json:"hours"`
	Image       null.String   `json:"image" rql:"filter"`
	URL         null.String   `json:"url" rql:"filter"`
	Twitter     null.String   `json:"twitter" rql:"filter"`
	Facebook    null.String   `json:"facebook" rql:"filter"`
	Attributes  []null.String `json:"attributes" rql:"filter"`
	LastUpdated null.String   `json:"last_updated" rql:"filter"`
}

// Parking represents a parking item
type Parking struct {
	ID          null.String `json:"id" rql:"filter"`
	Name        null.String `json:"name" rql:"filter"`
	Alias       null.String `json:"alias" rql:"filter"`
	BuildingID  null.String `json:"building_id" rql:"filter"`
	Description null.String `json:"description" rql:"filter"`
	Campus      null.String `json:"campus" rql:"filter"`
	Address     null.String `json:"address" rql:"filter"`
	Coordinates struct {
		Latitude  null.Float `json:"latitude"`
		Longitude null.Float `json:"longitude"`
	} `json:"coordinates"`
	LastUpdated null.String `json:"last_updated" rql:"filter"`
}

// Service represents an service item
type Service struct {
	ID          null.String `json:"id" rql:"filter"`
	Name        null.String `json:"name" rql:"filter"`
	Alias       null.String `json:"alias" rql:"filter"`
	BuildingID  null.String `json:"building_id" rql:"filter"`
	Description null.String `json:"description" rql:"filter"`
	Campus      null.String `json:"campus" rql:"filter"`
	Address     null.String `json:"address" rql:"filter"`
	Image       null.String `json:"image" rql:"filter"`
	Coordinates struct {
		Latitude  null.Float `json:"latitude"`
		Longitude null.Float `json:"longitude"`
	} `json:"coordinates"`
	Tags        null.String   `json:"tags" rql:"filter"`
	Attributes  []null.String `json:"attributes" rql:"filter"`
	LastUpdated null.String   `json:"last_updated" rql:"filter"`
}

// Exam represents an exam item
type Exam struct {
	ID         null.String `json:"id" rql:"filter"`
	CourseID   null.String `json:"course_id" rql:"filter"`
	CourseCode null.String `json:"course_code" rql:"filter"`
	Campus     null.String `json:"campus" rql:"filter"`
	Date       null.String `json:"date" rql:"filter"`
	Start      null.Int    `json:"start" rql:"filter"`
	End        null.Int    `json:"end" rql:"filter"`
	Duration   null.Int    `json:"duration" rql:"filter"`
	Sections   []struct {
		LectureCode null.String `json:"lecture_code"`
		Split       null.String `json:"split"`
		Location    null.String `json:"location"`
	} `json:"sections"`
	LastUpdated null.String `json:"last_updated" rql:"filter"`
}

// Eval represents an eval item
type Eval struct {
	ID     null.String `json:"id" rql:"filter"`
	Name   null.String `json:"name" rql:"filter"`
	Campus null.String `json:"campus" rql:"filter"`
	Terms  []struct {
		Term     null.String `json:"term"`
		Lectures []struct {
			LectureCode null.String `json:"lecture_code"`
			Firstname   null.String `json:"firstname"`
			Lastname    null.String `json:"lastname"`
			S1          null.Float  `json:"s1"`
			S2          null.Float  `json:"s2"`
			S3          null.Float  `json:"s3"`
			S4          null.Float  `json:"s4"`
			S5          null.Float  `json:"s5"`
			S6          null.Float  `json:"s6"`
			Invited     null.Int    `json:"invited"`
			Responses   null.Int    `json:"responses"`
		} `json:"lectures"`
	} `json:"terms"`
	LastUpdated null.String `json:"last_updated" rql:"filter"`
}

package models

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Week is our struct for users
type Week struct {
	ID           uint64 `gorm:"primary_key"`
	Date         time.Time
	Total        uint
	RepositoryID uint64
	Repository   Repository
}

// WeekForm is our struct to handle new users requests
type WeekForm struct {
	Date         time.Time
	Total        uint
	RepositoryID uint64
	Repository   Repository
}

// WeekJSON is the struct to return week in json
type WeekJSON struct {
	ID           uint64
	Date         time.Time
	Total        uint
	RepositoryID uint64
	Repository   Repository
}

// ValidateWeek takes a week form as parameter and check if its properties are valid
func ValidateWeek(week *WeekForm) error {
	_, err := govalidator.ValidateStruct(week)

	return err
}

package models

import (
	"github.com/asaskevich/govalidator"
)

// Contributor is our struct for users
type Contributor struct {
	ID           uint64 `gorm:"primary_key"`
	Name         string
	Total        uint
	Repository   Repository
	ImageLink    string
	RepositoryID uint64
}

// ContributorForm is our struct to handle new users requests
type ContributorForm struct {
	Name         string
	Total        uint
	Repository   Repository
	ImageLink    string
	RepositoryID uint64
}

// ContributorJSON is the struct to return contributor in json
type ContributorJSON struct {
	ID           uint64
	Name         string
	Total        uint
	Repository   Repository
	ImageLink    string
	RepositoryID uint64
}

// ValidateContributor takes a contributor form as parameter and check if its properties are valid
func ValidateContributor(contributor *ContributorForm) error {
	_, err := govalidator.ValidateStruct(contributor)

	return err
}

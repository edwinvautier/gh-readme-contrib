package models

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Repository is our struct for users
type Repository struct {
	ID        uint64 `gorm:"primary_key"`
	Name      string
	Author    string
	UpdatedAt time.Time
}

// RepositoryForm is our struct to handle new users requests
type RepositoryForm struct {
	Name      string
	Author    string
	UpdatedAt time.Time
}

// RepositoryJSON is the struct to return repository in json
type RepositoryJSON struct {
	ID        uint64
	Name      string
	Author    string
	UpdatedAt time.Time
}

// ValidateRepository takes a repository form as parameter and check if its properties are valid
func ValidateRepository(repository *RepositoryForm) error {
	_, err := govalidator.ValidateStruct(repository)

	return err
}

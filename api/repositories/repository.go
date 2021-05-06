package repositories

import (
	"errors"
	"time"

	"github.com/edwinvautier/gh-readme-contrib/api/models"
	"github.com/edwinvautier/gh-readme-contrib/shared/database"
	"github.com/jinzhu/gorm"
)

func CreateRepository(repository *models.Repository) error {
	repository.UpdatedAt = time.Now().Add(-7 * 25 * time.Hour)
	return database.Db.Debug().Create(repository).Error
}

func UpdateRepositoryByID(repository *models.Repository, id uint64) error {
	var err error
	var old models.Repository
	err = database.Db.Debug().Where("id = ?", id).First(&old).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("repository Not Found")
	}
	repository.ID = id

	err = database.Db.Debug().Save(&repository).Error
	if err != nil {
		return errors.New("Could'nt update repository")
	}

	return nil
}

func DeleteRepositoryByID(id uint64) (models.Repository, error) {

	var err error
	var repository models.Repository

	err = database.Db.Debug().Delete(&repository, id).Error
	if err != nil {
		return models.Repository{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return models.Repository{}, errors.New("repository Not Found")
	}

	return repository, err
}

func FindRepositoryByID(id uint64) (*models.Repository, error) {
	var err error
	var repository models.Repository
	err = database.Db.Debug().Model(models.Repository{}).Where("id = ?", id).Take(&repository).Error
	if err != nil {
		return &models.Repository{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &models.Repository{}, errors.New("repository Not Found")
	}

	return &repository, err
}

func FindRepositoryByAuthorAndName(name, author string) (*models.Repository, error) {
	var err error
	var repository models.Repository
	err = database.Db.Debug().Model(models.Repository{}).Where("name = ? AND author = ?", name, author).Take(&repository).Error
	if err != nil {
		return &models.Repository{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &models.Repository{}, errors.New("repository Not Found")
	}

	return &repository, err
}

func FindAllRepository() (*[]models.Repository, error) {
	var err error
	var repositoryList []models.Repository
	err = database.Db.Debug().Find(&repositoryList).Error
	if err != nil {
		return &[]models.Repository{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]models.Repository{}, errors.New("repository Not Found")
	}

	return &repositoryList, err
}

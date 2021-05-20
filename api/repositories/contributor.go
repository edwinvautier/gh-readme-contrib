package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/edwinvautier/gh-readme-contrib/api/models"
	"github.com/edwinvautier/gh-readme-contrib/shared/database"
	"github.com/edwinvautier/gh-readme-contrib/shared/env"
	"github.com/google/go-github/v35/github"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

func CreateContributor(contributor *models.Contributor) error {
	return database.Db.Debug().Create(contributor).Error
}

func UpdateContributorByID(contributor *models.Contributor, id uint64) error {
	var err error
	var old models.Contributor
	err = database.Db.Debug().Where("id = ?", id).First(&old).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("contributor Not Found")
	}
	contributor.ID = id

	err = database.Db.Debug().Save(&contributor).Error
	if err != nil {
		return errors.New("Could'nt update contributor")
	}

	return nil
}

func DeleteContributorsByRepositoryID(id uint64) (models.Contributor, error) {

	var err error
	var contributor models.Contributor

	err = database.Db.Debug().Where("repository_id = ?", id).Delete(&contributor).Error
	if err != nil {
		return models.Contributor{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return models.Contributor{}, errors.New("contributor Not Found")
	}

	return contributor, err
}

func FindContributorByID(id uint64) (*models.Contributor, error) {
	var err error
	var contributor models.Contributor
	err = database.Db.Debug().Model(models.Contributor{}).Where("id = ?", id).Take(&contributor).Error
	if err != nil {
		return &models.Contributor{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &models.Contributor{}, errors.New("contributor Not Found")
	}

	return &contributor, err
}

func FindTopContributorsByRepositoryID(id uint64) (*[]models.Contributor, error) {
	var err error
	var contributorList []models.Contributor
	err = database.Db.Debug().Where("repository_id = ?", id).Order("total desc").Limit(3).Find(&contributorList).Error
	if err != nil {
		return &[]models.Contributor{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]models.Contributor{}, errors.New("contributor Not Found")
	}

	return &contributorList, err
}

func FetchContributors(repository *models.Repository) ([]models.Contributor, error) {
	var contributors *[]models.Contributor

	// Check if its time to fetch again
	if shouldFetch(repository) {
		// Fetch and save weeks
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: env.GoDotEnvVariable("GH_API_TOKEN")},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)
		stats, _, err := client.Repositories.ListContributorsStats(context.TODO(), repository.Author, repository.Name)
		if err != nil {
			log.Error(err)
			return *contributors, err
		}
		DeleteContributorsByRepositoryID(repository.ID)
		for _, stat := range stats {
			var contributor models.Contributor
			contributor.Total = uint(*stat.Total)
			contributor.RepositoryID = repository.ID
			contributor.Name = *stat.Author.Login
			contributor.ImageLink = *stat.Author.AvatarURL

			CreateContributor(&contributor)
		}
		repository.UpdatedAt = time.Now()
		if err := UpdateRepositoryByID(repository, repository.ID); err != nil {
			return *contributors, err
		}
	}
	contributors, err := FindTopContributorsByRepositoryID(repository.ID)

	return *contributors, err
}

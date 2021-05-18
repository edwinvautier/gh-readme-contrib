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

func CreateWeek(week *models.Week) error {
	return database.Db.Debug().Create(week).Error
}

func UpdateWeekByID(week *models.Week, id uint64) error {
	var err error
	var old models.Week
	err = database.Db.Debug().Where("id = ?", id).First(&old).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("week Not Found")
	}
	week.ID = id

	err = database.Db.Debug().Save(&week).Error
	if err != nil {
		return errors.New("could'nt update week")
	}

	return nil
}

func DeleteWeeksByRepositoryID(id uint64) error {
	err := database.Db.Debug().Where("repository_id = ?", id).Delete(models.Week{}).Error
	if err != nil {
		return err
	}
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("week Not Found")
	}

	return err
}

func FindWeekByID(id uint64) (*models.Week, error) {
	var err error
	var week models.Week
	err = database.Db.Debug().Model(models.Week{}).Where("id = ?", id).Take(&week).Error
	if err != nil {
		return &models.Week{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &models.Week{}, errors.New("week Not Found")
	}

	return &week, err
}

func FindAllWeek() (*[]models.Week, error) {
	var err error
	var weekList []models.Week
	err = database.Db.Debug().Find(&weekList).Error
	if err != nil {
		return &[]models.Week{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]models.Week{}, errors.New("week Not Found")
	}

	return &weekList, err
}

func FindWeeksByRepositoryID(id uint64) (*[]models.Week, error) {
	var err error
	var weekList []models.Week
	err = database.Db.Debug().Order("date asc").Where("repository_id = ?", id).Find(&weekList).Error
	if err != nil {
		return &[]models.Week{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]models.Week{}, errors.New("weeks Not Found")
	}

	return &weekList, err
}

func FetchWeeks(repository *models.Repository) ([]models.Week, error) {
	var weeks *[]models.Week

	// Check if its time to fetch again
	if shouldFetch(repository) {
		// Fetch and save weeks
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: env.GoDotEnvVariable("GH_API_TOKEN")},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)
		stats, _, err := client.Repositories.ListCommitActivity(context.TODO(), repository.Author, repository.Name)
		if err != nil {
			log.Error(err)
			return *weeks, err
		}
		DeleteWeeksByRepositoryID(repository.ID)
		for _, stat := range stats {
			var week models.Week
			week.Date = stat.Week.Time
			week.Total = uint(*stat.Total)
			week.RepositoryID = repository.ID
			CreateWeek(&week)
		}
		repository.UpdatedAt = time.Now()
		if err := UpdateRepositoryByID(repository, repository.ID); err != nil {
			return *weeks, err
		}
	}
	weeks, err := FindWeeksByRepositoryID(repository.ID)

	return *weeks, err
}

func shouldFetch(repository *models.Repository) bool {
	lastMonday := GetFirstDateOfWeek()
	updatedAt := repository.UpdatedAt
	updatedAt = updatedAt.Add(7 * 24 * time.Hour)

	return updatedAt.Before(lastMonday)
}

/*
Gets the date of Monday this week
*/
func GetFirstDateOfWeek() (weekMonday time.Time) {
	now := time.Now()
 
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
 
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
}
package controllers

import (
	"net/http"

	"github.com/edwinvautier/gh-readme-contrib/api/models"
	"github.com/edwinvautier/gh-readme-contrib/api/repositories"
	"github.com/edwinvautier/gh-readme-contrib/shared/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// GetRepositoryByNameAndAuthor is the controller to get a repository by Name
func GetRepositoryByNameAndAuthor(c *gin.Context) {
	config := services.InitChartConfig(c)

	var repository *models.Repository
	var err error
	if repository, err = repositories.FindRepositoryByAuthorAndName(config.Name, config.Author); err != nil {
		// create repository
		repository.Author = config.Author
		repository.Name = config.Name
		if err := repositories.CreateRepository(repository); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	weeks, err := repositories.FetchWeeks(repository)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	config.WeeklyStats = weeks
	chart, err := services.GenerateChartFromContribs(config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Info(chart)
	c.Header("Content-type", "image/svg+xml")
	c.String(200, chart)
}

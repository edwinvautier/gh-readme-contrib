package controllers

import (
	"net/http"

	"github.com/edwinvautier/gh-readme-contrib/api/models"
	"github.com/edwinvautier/gh-readme-contrib/api/repositories"
	"github.com/edwinvautier/gh-readme-contrib/shared/services"
	"github.com/gin-gonic/gin"
)

// GetRepositoryByNameAndAuthor is the controller to get a repository by Name
func GetRepositoryByNameAndAuthor(c *gin.Context) {
	config := services.InitChartConfig(c)
	c.Header("Content-type", "image/svg+xml")

	var repository *models.Repository
	var err error
	if repository, err = repositories.FindRepositoryByAuthorAndName(config.Name, config.Author); err != nil {
		// create repository
		repository.Author = config.Author
		repository.Name = config.Name
		if err := repositories.CreateRepository(repository); err != nil {
			c.String(http.StatusOK, "<svg width=\"440\" height=\"270\" xmlns=\"http://www.w3.org/2000/svg\"><text>Empty data</text</svg>")
			return
		}
	}

	weeks, err := repositories.FetchWeeks(repository)
	if err != nil {
		c.String(http.StatusOK, "<svg width=\"440\" height=\"270\" xmlns=\"http://www.w3.org/2000/svg\"><text>Empty data</text</svg>")
		return
	}
	config.WeeklyStats = weeks
	chart, err := services.GenerateChartFromContribs(config)
	if err != nil {
		c.String(http.StatusOK, "<svg width=\"440\" height=\"270\" xmlns=\"http://www.w3.org/2000/svg\"><text>Empty data</text</svg>")
		return
	}

	c.Header("Cache-Control", "public, max-age=86400")
	c.String(http.StatusOK, chart)
}

func GetContributorsByNameAndAuthor(c *gin.Context) {
	config := services.InitChartConfig(c)
	c.Header("Content-type", "image/svg+xml")

	var repository *models.Repository
	var err error
	if repository, err = repositories.FindRepositoryByAuthorAndName(config.Name, config.Author); err != nil {
		// create repository
		repository.Author = config.Author
		repository.Name = config.Name
		if err := repositories.CreateRepository(repository); err != nil {
			c.String(http.StatusOK, "<svg width=\"440\" height=\"270\" xmlns=\"http://www.w3.org/2000/svg\"><text>Empty data</text</svg>")
			return
		}
	}

	contributors, err := repositories.FetchContributors(repository)
	if err != nil {
		c.String(http.StatusOK, "<svg width=\"440\" height=\"270\" xmlns=\"http://www.w3.org/2000/svg\"><text>Empty data</text</svg>")
		return
	}
	config.ContributorsStats = contributors
	chart, err := services.GenerateChartFromContributors(config)
	if err != nil {
		c.String(http.StatusOK, "<svg width=\"440\" height=\"270\" xmlns=\"http://www.w3.org/2000/svg\"><text>Empty data</text</svg>")
		return
	}
	c.Header("Cache-Control", "public, max-age=86400")
	c.String(http.StatusOK, chart)
}

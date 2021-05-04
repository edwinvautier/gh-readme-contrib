package controllers

import (
	"context"
	"net/http"

	"github.com/edwinvautier/gh-readme-contrib/api/models"
	"github.com/edwinvautier/gh-readme-contrib/api/repositories"
	"github.com/edwinvautier/gh-readme-contrib/shared/helpers"
	"github.com/edwinvautier/gh-readme-contrib/shared/services"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v35/github"
	log "github.com/sirupsen/logrus"
)

// CreateRepository is the controller to create a new repository
func CreateRepository(c *gin.Context) {
  var repositoryForm models.RepositoryForm
  if err := c.ShouldBindJSON(&repositoryForm); err != nil {
    c.JSON(http.StatusBadRequest, "invalid informations provided")
    return
  }

	if err := models.ValidateRepository(&repositoryForm); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	repository := models.Repository{ 
	  Name: repositoryForm.Name,
	  Author: repositoryForm.Author,
	  Base64: repositoryForm.Base64,
	}

	if err := repositories.CreateRepository(&repository); err != nil {
		c.JSON(http.StatusInternalServerError, "Couldn't create repository. Try again.")
		return
	}

	c.JSON(http.StatusOK, repository)
}

// GetRepositoryByNameAndAuthor is the controller to get a repository by Name
func GetRepositoryByNameAndAuthor(c *gin.Context) {
  name := c.Param("repository")
  author := c.Param("author")
	client := github.NewClient(nil)
	stats, _, err := client.Repositories.ListContributorsStats(context.TODO(), author, name)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	chart, err := services.GenerateChartFromContribs(stats)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	
	c.JSON(http.StatusOK, chart)
}

// GetAllRepository is the controller to get all repository
func GetAllRepository(c *gin.Context) {
	repositoryList, err := repositories.FindAllRepository();
  if err != nil {
		c.JSON(http.StatusNotFound, "Couldn't find repository. Try again.")
		return
	}

	c.JSON(http.StatusOK, repositoryList)
}

// UpdateRepository is the controller to update a repository
func UpdateRepository(c *gin.Context) {
  ID := helpers.ParseStringToUint64(c.Param("id")) 
  var repository models.Repository
  if err := c.ShouldBindJSON(&repository); err != nil {
    c.JSON(http.StatusBadRequest, "invalid informations provided")
    return
  }

	if err := repositories.UpdateRepositoryByID(&repository, ID); err != nil {
		c.JSON(http.StatusInternalServerError, "Couldn't update repository. Try again.")
		return
	}

	c.JSON(http.StatusOK, repository)
}

// DeleteRepositoryByID is the controller to delete a repository by id
func DeleteRepositoryByID(c *gin.Context) {
  ID := helpers.ParseStringToUint64(c.Param("id"))

	_, err := repositories.DeleteRepositoryByID(ID);
  if err != nil {
		c.JSON(http.StatusNotFound, "Couldn't find repository. Try again.")
		return
	}

	c.JSON(http.StatusOK, "repository deleted")
}
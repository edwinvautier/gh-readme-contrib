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
	log.Info(chart)
	svg := `
	<svg id="Layer_1" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="20px" height="20px" viewBox="0 0 20 20" enable-background="new 0 0 20 20" xml:space="preserve">
<path fill="#4D4D4D" d="M1.171,20c-0.292,0-0.585-0.111-0.809-0.334c-0.448-0.447-0.449-1.172-0.003-1.619l8.022-8.045L0.349,1.968 c-0.448-0.447-0.448-1.172,0-1.62c0.447-0.447,1.171-0.447,1.618,0l8.032,8.031l8.02-8.042c0.446-0.449,1.172-0.449,1.62-0.002 c0.448,0.447,0.448,1.171,0.002,1.62L11.618,10l8.033,8.033c0.447,0.447,0.447,1.172,0,1.619c-0.447,0.446-1.172,0.446-1.619,0 l-8.03-8.031l-8.02,8.043C1.758,19.888,1.464,20,1.171,20z"></path>
</svg>`
	c.Header("Content-type", "image/svg+xml")
	c.String(200, svg)
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
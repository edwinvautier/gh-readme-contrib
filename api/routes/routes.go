package routes

import (
	"github.com/edwinvautier/gh-readme-contrib/api/controllers"

	"github.com/gin-gonic/gin"
)

// Init initializes router with the following routes
func Init(r *gin.Engine) {
	r.GET("/", controllers.Healtcheck)

	api := r.Group("/api")
	{
		api.GET("/:author/:repository", controllers.GetRepositoryByNameAndAuthor)

		v1 := api.Group("/v1")
		{
			v1.GET("/activity/:author/:repository", controllers.GetRepositoryByNameAndAuthor)
			v1.GET("/contributors/:author/:repository", controllers.GetContributorsByNameAndAuthor)
		}
	}
}

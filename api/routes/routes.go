package routes

import (
	"github.com/edwinvautier/gh-readme-contrib/api/controllers"

	"github.com/gin-gonic/gin"
)

// Init initializes router with the following routes
func Init(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/:author/:repository", controllers.GetRepositoryByNameAndAuthor)
	}
}

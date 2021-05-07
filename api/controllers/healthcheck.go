package controllers

import "github.com/gin-gonic/gin"


func Healtcheck(c *gin.Context) {
	c.JSON(200, "api up")
}
package views

import (
	"errors"
	"fmt"
	"redis/controllers"
	"redis/models"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func GetEmails(c *gin.Context) {
	emails, err := controllers.GetEmails()
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, emails)
}

func AddEmail(c *gin.Context) {
	email := models.Emails{}
	c.BindJSON(&email)

	status, err := controllers.AddEmail(email)
	if !status {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(201, email)
}

func GetEmail(c *gin.Context) {
	searchEmail := c.Param("email")

	email, err := controllers.GetEmail(searchEmail)
	if err == redis.Nil {
		c.AbortWithError(404, errors.New("Email not found"))
		return
	} else if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, email)
}

func UpdateEmail(c *gin.Context) {
	searchEmail := c.Param("email")

	email := models.Emails{}
	c.BindJSON(&email)

	status, err := controllers.UpdateEmail(email, searchEmail)
	if !status && err == redis.Nil {
		c.AbortWithError(404, errors.New("Email not found"))
		return
	} else if !status && err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.Data(204, gin.MIMEJSON, nil)
}

func DeleteEmail(c *gin.Context) {
	searchEmail := c.Param("email")

	status, err := controllers.DeleteEmail(searchEmail)
	if !status && err == redis.Nil {
		fmt.Println("view redis nil")
		c.AbortWithError(404, errors.New("Email not found"))
		return
	} else if !status && err != nil {
		fmt.Println("view err != nil")
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.Data(200, gin.MIMEJSON, nil)
}

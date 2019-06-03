package main

import (
	"fmt"
	"redis/utils"
	"redis/views"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	cors "github.com/tommy351/gin-cors"
)

// GetEngine
func GetEngine() *gin.Engine {
	router := gin.Default()
	router.Use(gin.ErrorLoggerT(gin.ErrorTypePrivate))
	router.Use(cors.Middleware(cors.Options{}))
	router.Use(location.Default())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/v1/emails", views.GetEmails)
	router.POST("/v1/emails", views.AddEmail)
	router.GET("/v1/email/:email", views.GetEmail)
	router.PUT("/v1/email/:email", views.UpdateEmail)
	router.DELETE("/v1/email/:email", views.DeleteEmail)

	return router
}

func initApp() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Error(err)
	}

	err = utils.InitRedisClient()
	if err != nil {
		log.Error(err)
	}

	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetLevel(log.DebugLevel)
}

func main() {
	initApp()
	gin.SetMode(gin.DebugMode)
	defer utils.CloseRedisClient()
	router := GetEngine()
	router.Run(fmt.Sprintf(":%s", viper.GetString("server.port")))
}

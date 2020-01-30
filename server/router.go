package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"{{ .ProjectPath }}/config"
	"{{ .ProjectPath }}/controllers"
)

// NewRouter is constructor for router
func NewRouter() (*gin.Engine, error) {
	c := config.GetConfig()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins: c.GetStringSlice("server.cors"),
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))


	version := router.Group("/" + c.GetString("server.version"))

	healthController := controllers.NewHealthController()
	version.GET("/health", healthController.Index)

	return router, nil
}

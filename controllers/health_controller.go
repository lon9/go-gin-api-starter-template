package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController controller for health request
type HealthController struct{}

// NewHealthController is constructer for HealthController
func NewHealthController() *HealthController {
	return new(HealthController)
}

// Index is index route for health
func (hc *HealthController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"result":  "OK",
	})
}

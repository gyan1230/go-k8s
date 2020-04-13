package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Home :
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "In Home"})
}

//HealthHandler :
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Health is OK"})
}

//ReadinessHandler :
func ReadinessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Readiness is OK"})
}

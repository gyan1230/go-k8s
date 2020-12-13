package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var count int

//Home :
func Home(c *gin.Context) {
	count++
	v := fmt.Sprintf("no is %d", count)
	c.JSON(http.StatusOK, gin.H{"Visitor": v})
}

//HealthHandler :
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Health is OK"})
}

//ReadinessHandler :
func ReadinessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Readiness is OK"})
}

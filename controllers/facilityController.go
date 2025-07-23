package controllers

import (
	"net/http"

	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

func GetAllFacilitiesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		facilities, err := services.GetAllFacilities(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list facilities"})
			return
		}
		c.JSON(http.StatusOK, facilities)
	}
}

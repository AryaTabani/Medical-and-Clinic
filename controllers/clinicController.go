package controllers

import (
	"net/http"

	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

func GetAllClinicsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		clinics, err := services.GetAllClinics(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list clinics"})
			return
		}
		c.JSON(http.StatusOK, clinics)
	}
}
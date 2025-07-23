package controllers

import (
	"net/http"
	"strings"

	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

func GetDoctorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		specialty := c.Query("specialty")
		if strings.ToLower(specialty) == "all" {
			specialty = ""
		}

		doctors, err := services.GetDoctors(c.Request.Context(), specialty)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list doctors"})
			return
		}
		c.JSON(http.StatusOK, doctors)
	}
}

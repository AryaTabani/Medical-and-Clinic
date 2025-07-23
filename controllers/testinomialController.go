package controllers

import (
	"net/http"

	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

func GetAllTestimonialsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		testimonials, err := services.GetAllTestimonials(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list testimonials"})
			return
		}
		c.JSON(http.StatusOK, testimonials)
	}
}
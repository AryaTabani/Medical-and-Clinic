package controllers

import (
	"net/http"

	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

func GetAboutPageDataHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := services.GetAboutPageData(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}

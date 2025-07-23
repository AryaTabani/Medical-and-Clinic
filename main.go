package main

import (
	db "example.com/m/v2/DB"
	"example.com/m/v2/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	router := gin.Default()

	router.GET("/about", controllers.GetAboutPageDataHandler())
	router.GET("/facilities", controllers.GetAllFacilitiesHandler())
	router.GET("/doctors", controllers.GetDoctorsHandler())
	router.GET("/testimonials", controllers.GetAllTestimonialsHandler())

	router.Run(":8080")
}

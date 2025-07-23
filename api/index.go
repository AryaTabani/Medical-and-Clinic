package handler

import (
	"net/http"

	db "example.com/m/v2/DB"
	"example.com/m/v2/controllers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	db.InitDB()

	router = gin.Default()

	router.GET("/about", controllers.GetAboutPageDataHandler())
	router.GET("/facilities", controllers.GetAllFacilitiesHandler())
	router.GET("/doctors", controllers.GetDoctorsHandler())
	router.GET("/testimonials", controllers.GetAllTestimonialsHandler())
	router.GET("/clinics", controllers.GetAllClinicsHandler())
}


func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}

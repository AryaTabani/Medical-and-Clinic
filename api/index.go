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

	// ۲. ساخت روتر Gin
	router = gin.Default()

	// ۳. تعریف تمام مسیرها (کپی شده از main.go شما)
	router.GET("/about", controllers.GetAboutPageDataHandler())
	router.GET("/facilities", controllers.GetAllFacilitiesHandler())
	router.GET("/doctors", controllers.GetDoctorsHandler())
	router.GET("/testimonials", controllers.GetAllTestimonialsHandler())
	router.GET("/clinics", controllers.GetAllClinicsHandler())
}


func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}

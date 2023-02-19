package routes

import (
	"github.com/gin-gonic/gin"
	"quiz-3-sanbercode_golang/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/segitiga-sama-sisi/", controllers.SegitigaSamaSisi)
	router.GET("/persegi", controllers.Persegi)
	router.GET("/persegi-panjang", controllers.PersegiPanjang)
	router.GET("/lingkaran", controllers.Lingkaran)

	router.GET("/category", controllers.GetAllCategories)
	router.POST("/category", controllers.InsertCategory)
	router.PUT("/category/:id", controllers.UpdateCategory)
	router.DELETE("/category/:id", controllers.DeleteCategory)
	return router
}

package main

import (
	"rental-car/controllers"
	"rental-car/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	// Car Controller
	r.POST("/rentalcar/v1/post/car", controllers.PostCarV1)
	r.PUT("/rentalcar/v1/put/car/:id", controllers.PutCarV1)
	r.DELETE("/rentalcar/v1/delete/car/:id", controllers.DeleteCarV1)
	r.GET("/rentalcar/v1/get/car", controllers.GetCarV1)
	r.GET("/rentalcar/v1/get/car/:id", controllers.GetCarV1ById)

	// Order Controller
	r.POST("/rentalcar/v1/post/order", controllers.PostOrderV1)
	r.PUT("/rentalcar/v1/put/order/:id", controllers.PutOrderV1)
	r.DELETE("/rentalcar/v1/delete/order/:id", controllers.DeleteOrderV1)
	r.GET("/rentalcar/v1/get/order", controllers.GetOrderV1)
	r.GET("/rentalcar/v1/get/order/:id", controllers.GetOrderV1ById)

	r.Run() // listen and serve
}

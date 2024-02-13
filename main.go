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
	r.POST("/rentalcar/v1/post/car", controllers.PostCarV1)

	r.Run() // listen and serve on 0.0.0.0:8080
}

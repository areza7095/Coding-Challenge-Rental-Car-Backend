package main

import (
	"rental-car/initializers"
	"rental-car/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Car{})
	initializers.DB.AutoMigrate(&models.Order{})
}

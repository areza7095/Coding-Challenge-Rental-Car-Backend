package controllers

import (
	"rental-car/initializers"
	"rental-car/models"

	"github.com/gin-gonic/gin"
)

func PostCarV1(c *gin.Context) {
	// Get data req body
	var bodyCar struct {
		CarId     int
		CarName   string
		DayRate   float64
		MonthRate float64
		Image     string
	}

	c.Bind(&bodyCar)

	// Create a Post
	car := models.Car{CarID: bodyCar.CarId, CarName: bodyCar.CarName, DayRate: bodyCar.DayRate, MonthRate: bodyCar.MonthRate, Image: bodyCar.Image}

	result := initializers.DB.Create(&car)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return
	c.JSON(201, gin.H{
		"code":    201,
		"message": "OK",
		"data":    car,
	})
}

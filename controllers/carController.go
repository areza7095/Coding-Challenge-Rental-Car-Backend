package controllers

import (
	"net/http"
	"rental-car/initializers"
	"rental-car/models"

	"github.com/gin-gonic/gin"
)

func PostCarV1(c *gin.Context) {
	// Get data req body
	var bodyCar struct {
		CarName   string  `json:"car_name" binding:"required"`
		DayRate   float64 `json:"day_rate" binding:"required"`
		MonthRate float64 `json:"month_rate" binding:"required"`
		Image     string  `json:"image" binding:"required"`
	}

	if err := c.BindJSON(&bodyCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad Request: Invalid JSON payload",
		})
		return
	}

	// Create a Post
	car := models.Car{CarName: bodyCar.CarName, DayRate: bodyCar.DayRate, MonthRate: bodyCar.MonthRate, Image: bodyCar.Image}

	result := initializers.DB.Create(&car)

	if result.Error != nil {
		// Return error
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	// Return Response
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "OK",
	})
}

func GetCarV1(c *gin.Context) {
	//Get the All of Car Data
	var cars []models.Car

	if err := initializers.DB.Preload("Orders").Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to fetch cars with orders",
		})
		return
	}

	//Return Response
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": cars,
	})
}

func GetCarV1ById(c *gin.Context) {
	//Get Param
	id := c.Param("id")

	//Get the All of Car Data
	var cars models.Car
	result := initializers.DB.First(&cars, id)

	// Checking the data
	if result.Error != nil {
		// Returning 404
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		// Return nothing
		return
	}

	//Return Response
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": cars,
	})
}

func PutCarV1(c *gin.Context) {
	//Get Param
	id := c.Param("id")

	// Get data req body
	var bodyCar struct {
		CarName   string  `json:"car_name" binding:"required"`
		DayRate   float64 `json:"day_rate" binding:"required"`
		MonthRate float64 `json:"month_rate" binding:"required"`
		Image     string  `json:"image" binding:"required"`
	}

	if err := c.Bind(&bodyCar); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":    http.StatusServiceUnavailable,
			"message": "Bad Request: Invalid JSON payload",
		})
		return
	}

	// Find the data with ID
	var car models.Car
	result := initializers.DB.First(&car, id)

	// Sending 404
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Not Found",
		})
		return
	}

	//Update the data
	putResult := initializers.DB.Model(&car).Updates(models.Car{
		CarName:   bodyCar.CarName,
		DayRate:   bodyCar.DayRate,
		MonthRate: bodyCar.MonthRate,
		Image:     bodyCar.Image,
	})

	if putResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	// Return with response
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Updated Succesfully",
	})
}

func DeleteCarV1(c *gin.Context) {
	//Get Param
	id := c.Param("id")

	// Find the data with ID
	var car models.Car
	result := initializers.DB.First(&car, id)

	// Checking data
	if result.Error != nil {
		// Response with 404
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Not Found",
		})
		return
	}

	delResult := initializers.DB.Delete(&models.Car{}, id)

	// Checking error
	if delResult.Error != nil {
		// Response with 404
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad Request",
		})
		return
	}

	// Return Response
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Delete Succesfully",
	})
}

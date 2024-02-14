package controllers

import (
	"net/http"
	"rental-car/initializers"
	"rental-car/models"
	"time"

	"github.com/gin-gonic/gin"
)

func PostOrderV1(c *gin.Context) {
	// Get data req body
	var bodyOrder struct {
		CarID           int       `json:"car_id" binding:"required"`
		OrderDate       time.Time `json:"order_date" binding:"required"`
		PickupDate      time.Time `json:"pickup_date" binding:"required"`
		DropoffDate     time.Time `json:"dropoff_date" binding:"required"`
		PickupLocation  string    `json:"pickup_location" binding:"required"`
		DropoffLocation string    `json:"dropoff_location" binding:"required"`
	}

	if err := c.BindJSON(&bodyOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Bad Request: Invalid JSON payload",
		})
		return
	}
	// Create order
	order := models.Order{
		CarID:           bodyOrder.CarID,
		OrderDate:       bodyOrder.OrderDate,
		PickupDate:      bodyOrder.PickupDate,
		DropoffDate:     bodyOrder.DropoffDate,
		PickupLocation:  bodyOrder.PickupLocation,
		DropoffLocation: bodyOrder.DropoffLocation,
	}

	// Save to database
	result := initializers.DB.Create(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	// Return with response
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "OK",
	})
}

func GetOrderV1(c *gin.Context) {
	//Get the All of Order Data
	var orders []models.Order

	if err := initializers.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to fetch cars with orders",
		})
		return
	}

	//Return Response
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": orders,
	})
}

func GetOrderV1ById(c *gin.Context) {
	//Get Param
	id := c.Param("id")

	//Get the Order Data by ID
	var orders []models.Order

	if err := initializers.DB.First(&orders, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Not Found",
		})
		return
	}

	//Return Response
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": orders,
	})
}

func PutOrderV1(c *gin.Context) {
	//Get Param
	id := c.Param("id")

	// Get data req body
	var bodyOrder struct {
		CarID           int       `json:"car_id"`
		OrderDate       time.Time `json:"order_date"`
		PickupDate      time.Time `json:"pickup_date"`
		DropoffDate     time.Time `json:"dropoff_date"`
		PickupLocation  string    `json:"pickup_location"`
		DropoffLocation string    `json:"dropoff_location"`
	}

	if err := c.Bind(&bodyOrder); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":    http.StatusServiceUnavailable,
			"message": "Bad Request: Invalid JSON payload",
		})
		return
	}

	// Find the data with ID
	var order models.Order
	result := initializers.DB.First(&order, id)

	// Sending 404
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Not Found",
		})
		return
	}

	//Update the data
	putResult := initializers.DB.Model(&order).Updates(models.Order{
		OrderDate:       bodyOrder.OrderDate,
		PickupDate:      bodyOrder.PickupDate,
		DropoffDate:     bodyOrder.DropoffDate,
		PickupLocation:  bodyOrder.PickupLocation,
		DropoffLocation: bodyOrder.DropoffLocation,
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

func DeleteOrderV1(c *gin.Context) {
	//Get Param
	id := c.Param("id")

	// Find the data with ID
	var order models.Order
	result := initializers.DB.First(&order, id)

	// Checking data
	if result.Error != nil {
		// Response with 404
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Not Found",
		})
		return
	}

	delResult := initializers.DB.Delete(&models.Order{}, id)

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

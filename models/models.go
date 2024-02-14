package models

import (
	"time"
)

type Car struct {
	CarID     int     `gorm:"primaryKey"`
	CarName   string  `gorm:"size:50;not null"`
	DayRate   float64 `gorm:"not null"`
	MonthRate float64 `gorm:"not null"`
	Image     string  `gorm:"size:256;not null"`
	Orders    []Order `gorm:"foreignKey:CarID"`
}

type Order struct {
	OrderID         int       `gorm:"primaryKey"`
	CarID           int       `gorm:"column:car_id"`
	OrderDate       time.Time `gorm:"not null"`
	PickupDate      time.Time `gorm:"not null"`
	DropoffDate     time.Time `gorm:"not null"`
	PickupLocation  string    `gorm:"size:50;not null"`
	DropoffLocation string    `gorm:"size:50;not null"`
}

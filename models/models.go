package models

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	CarID     int     `gorm:"primaryKey"`
	CarName   string  `gorm:"size:50;not null"`
	DayRate   float64 `gorm:"not null"`
	MonthRate float64 `gorm:"not null"`
	Image     string  `gorm:"size:256;not null"`
}

type Order struct {
	gorm.Model
	OrderID         int       `gorm:"primaryKey"`
	CarID           int       `gorm:"not null"`
	Car             Car       `gorm:"foreignKey:CarID"` // Relationship with Car table
	OrderDate       time.Time `gorm:"not null"`
	PickupDate      time.Time `gorm:"not null"`
	DropoffDate     time.Time `gorm:"not null"`
	PickupLocation  string    `gorm:"size:50;not null"`
	DropoffLocation string    `gorm:"size:50;not null"`
}

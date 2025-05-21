package models

import "gorm.io/gorm"
import "time"

type Vehicles struct {
	gorm.Model
	ID               uint      `json:"id" gorm:"primaryKey"`
	vehicle_id       uint      `json:"vehicle_id"`
	brand            string    `json:"unique;not null" json:"email"`
	model            string    `json:"size:255;not null" json:"name"`
	year             time.Time `json:"year"`
	fuel_type        string    `json:"fuel_type"`
	fuel_consumption string    `json:"fuel_consumption"`
	status           string    `json:"status"`
}

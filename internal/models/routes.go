package models

import "gorm.io/gorm"
import "time"

type Routes struct {
	gorm.Model
	ID                  uint      `json:"id" gorm:"primaryKey"`
	departure_point     string    `json:"departure_point"`
	destination_point   string    `json:"destination_point"`
	distance            uint      `json:"distance"`
	travel_time         time.Time `json:"travel_time"`
	possible_stops      string    `json:"possible_stops"`
	possible_delays     string    `json:"possible_delays"`
	vehicles_vehicle_id uint      `json:"vehicles_vehicle_id"`
}

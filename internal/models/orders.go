package models

import "time"

type Orders struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	client_id      uint      `json:"client_id"`
	route_id       uint      `json:"route_id"`
	delivery_date  time.Time `json:"delivery_date"`
	cargo_quantity uint      `json:"cargo_quantity"`
	cargo_type     string    `json:"cargo_type"`
	cost           uint      `json:"cost"`
}

package models

import "time"

// Структура для учета операций с товарами
type InventoryTransaction struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	ProductID       uint      `json:"product_id"`
	SupplierID      uint      `json:"supplier_id"`
	Quantity        int       `json:"quantity"`
	TransactionType string    `json:"transaction_type"` // "incoming" или "outgoing"
	CreatedAt       time.Time `json:"created_at"`
}

package models

// Связь между продуктами и поставщиками
type ProductSupplier struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	ProductID  uint `json:"product_id"`
	SupplierID uint `json:"supplier_id"`
}

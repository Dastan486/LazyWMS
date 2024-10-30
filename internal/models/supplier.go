package models

type Supplier struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	ContactInfo string `json:"contact_info"`
}

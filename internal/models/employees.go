package models

// Структура для учета 
type Employees struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	full_name       string `json:"full_name"`
	position        string `json:"position"`
  phone           string `json:"phone"`
  email           string `json:email`

}
type Order_employees struct {
	ID         uint `json:"id" gorm:"primaryKey"`
  order_id   uint `json:"order_id"`
	employee_id uint `json:"employee_id"`
}

package models

type Clients struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	phone string `json:"phone"`
  email string `json:"email"`
  address string `json:"address"`
  collaboration_conditions string "json:collaboration_conditions"
}

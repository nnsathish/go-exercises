package models

// embed gorm.Model for consistency!
type Product struct {
  ID uint `json:"id" gorm:"primary_key"`
  Name string `json:"product_name"`
  Price float64 `json:"price"`
  Quantity uint `json:"quantity"`
}

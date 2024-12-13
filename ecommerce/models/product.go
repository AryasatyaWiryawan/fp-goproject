package models

type Product struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	UserID   uint    `json:"user_id"` // Foreign key untuk User
}

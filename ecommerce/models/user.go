package models

type User struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"` // Exclude from JSON responses
	Products []Product `gorm:"constraint:OnDelete:CASCADE;" json:"products"`
}

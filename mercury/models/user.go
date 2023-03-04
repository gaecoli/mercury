package models

type User struct {
	ID
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;index"`
	Password string `json:"-" gorm:"not null;default ''"`
	Timestamps
	SoftDeletes
}

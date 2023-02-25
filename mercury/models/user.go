package models

type User struct {
	ID
	Name     string `json:"name" gorm:"not null"`
	Number   string `json:"number" gorm:"not null;index"`
	Password string `json:"password" gorm:"not null;default ''"`
	Timestamps
	SoftDeletes
}

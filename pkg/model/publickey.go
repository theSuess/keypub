package model

type PublicKey struct {
	ID      string `json:"id"`
	Content string `json:"content" gorm:"not null"`
	Name    string `json:"name" gorm:"not null"`
	User    *User  `json:"user"`
	UserID  string `json:"userId" gorm:"not null"`
}

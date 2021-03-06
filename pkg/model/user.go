package model

type User struct {
	ID       string       `json:"id"`
	Username string       `json:"username" gorm:"UNIQUE"`
	Name     *string      `json:"name"`
	Keys     []*PublicKey `json:"keys"`
	Groups   []*Group     `json:"groups" gorm:"many2many:user_groups"`
}

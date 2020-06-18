package model

type Group struct {
	ID     string  `json:"id"`
	Name   string  `json:"name" gorm:"UNIQUE"`
	Users  []*User `json:"users" gorm:"many2many:user_groups"`
	Owners []*User `json:"owners" gorm:"many2many:groups_owners"`
}

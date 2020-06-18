package service

import (
	"github.com/jinzhu/gorm"
	"github.com/theSuess/keypub/pkg/model"
)

type KeyService struct {
	db *gorm.DB
}

func Key(db *gorm.DB) *KeyService {
	return &KeyService{
		db: db,
	}
}

func (ks *KeyService) AddKey(key *model.PublicKey) error {
	key.ID = generateID()
	return ks.db.Create(key).Error
}

func (ks *KeyService) Owner(key *model.PublicKey) (*model.User, error) {
	user := &model.User{}
	err := ks.db.Where(&model.User{ID: key.UserID}).First(user).Error
	return user, err
}

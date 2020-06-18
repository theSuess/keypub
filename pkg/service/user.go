package service

import (
	"github.com/jinzhu/gorm"
	"github.com/theSuess/keypub/pkg/model"
)

type UserService struct {
	db *gorm.DB
}

func User(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (us *UserService) Register(u *model.User) error {
	u.ID = generateID()
	return us.db.Create(u).Error
}

func (us *UserService) ByID(id string) (*model.User, error) {
	u := &model.User{}
	err := us.db.Where(model.User{ID: id}).First(u).Error
	return u, err
}

func (us *UserService) FindAll(limit int, offset int) ([]*model.User, error) {
	users := []*model.User{}
	err := us.db.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (us *UserService) GroupsOf(user *model.User, limit int, offset int) ([]*model.Group, error) {
	groups := []*model.Group{}
	err := us.db.Model(user).Limit(limit).Offset(offset).Related(&groups, "Groups").Error
	return groups, err
}

func (us *UserService) KeysOf(user *model.User, limit int, offset int) ([]*model.PublicKey, error) {
	keys := []*model.PublicKey{}
	err := us.db.Model(user).Limit(limit).Offset(offset).Related(&keys).Error
	return keys, err
}

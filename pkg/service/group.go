package service

import (
	"github.com/jinzhu/gorm"
	"github.com/theSuess/keypub/pkg/model"
)

type GroupService struct {
	db *gorm.DB
}

func Group(db *gorm.DB) *GroupService {
	return &GroupService{
		db: db,
	}
}

func (gs *GroupService) FindAll(limit int, offset int) ([]*model.Group, error) {
	groups := []*model.Group{}
	err := gs.db.Limit(limit).Offset(offset).Find(&groups).Error
	return groups, err
}

func (gs *GroupService) UsersOf(group *model.Group) ([]*model.User, error) {
	users := []*model.User{}
	err := gs.db.Model(group).Related(&users, "Users").Error
	return users, err
}

func (gs *GroupService) OwnersOf(group *model.Group) ([]*model.User, error) {
	users := []*model.User{}
	err := gs.db.Model(group).Related(&users, "Owners").Error
	return users, err
}

func (gs *GroupService) CreateGroup(group *model.Group, requester *model.User) error {
	group.ID = generateID()
	group.Users = []*model.User{requester}
	group.Owners = []*model.User{requester}
	return gs.db.Create(group).Error
}

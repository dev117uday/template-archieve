package db

import (
	"go-postgres/errset"
	"go-postgres/logrus"
	"go-postgres/model"

	"gorm.io/gorm/clause"
)

func AddUser(user *model.User) error {

	err := Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)

	if err.Error != nil {
		logrus.LogIt(1, "AddUser", "unable to add user to db", err.Error)
		return errset.ErrInternalServer
	}

	return nil
}

// get user from database by id
func GetUser(id string) (*model.User, error) {
	user := model.User{}

	err := Db.Where(model.User{Gid: id}).First(&user)

	if err.Error != nil {
		logrus.LogIt(1, "GetUser", "unable to get user from db", err.Error)
		return nil, errset.ErrInternalServer
	}
	return &user, nil
}

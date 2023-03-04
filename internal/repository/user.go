package repository

import (
	"sirawit/shop/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserQuery interface {
	Register(input model.User) (*model.User, error)
}

type userQuery struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserQuery {
	db.AutoMigrate(&model.User{})
	return &userQuery{db}
}

func (u *userQuery) Register(input model.User) (*model.User, error) {
	if err := u.db.Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func ConnectToUserDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

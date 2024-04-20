package userrepository

import (
	"game_suit/core/entity"
	userrepository "game_suit/core/repository"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) userrepository.UserRepository {
	return &repository{
		DB: db,
	}
}

func (r repository) GetUser(id string) (*entity.User, error) {
	var (
		user *entity.User

		db = r.DB
	)
	err := db.Where("id_game =?", id).First(&user).Error
	return user, err
}

func (r repository) GetUsers() ([]entity.User, error) {
	var user []entity.User
	err := r.DB.Find(&user).Error
	return user, err
}

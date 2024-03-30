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

func (r repository) GetUser(id int) (*entity.User, error) {
	var (
		user *entity.User

		db = r.DB
	)

	err := db.Where("id =?", id).First(user).Error

	return user, err

}

package repository

import "game_suit/core/entity"

type UserRepository interface {
	GetUser(id string) (*entity.User, error)
	GetUsers() ([]entity.User, error)
}

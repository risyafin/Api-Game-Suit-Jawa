package repository

import "game_suit/core/entity"

type UserRepository interface {
	GetUser(id int) (*entity.User, error)
}

package module

import (
	"game_suit/core/entity"
	"game_suit/core/repository"
)

type UserUsecase interface {
	GetUser(id string) (*entity.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (e *userUsecase) GetUser(id string) (*entity.User, error) {
	return e.userRepository.GetUser(id)
}

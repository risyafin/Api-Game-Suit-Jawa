package module

import (
	"game_suit/core/entity"
	"game_suit/core/repository"
)

type UserUsecase interface {
	GetUser(id int) (*entity.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserRepository(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (e *userUsecase) GetUser(id int) (*entity.User, error) {
	return e.userRepository.GetUser(id)
}

package user

import (
	"microblog/internal/core/domain"
	"microblog/internal/core/port/repository"
	"microblog/internal/core/port/service/user"
)

type userService struct {
	Repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) user.UserService {
	return &userService{
		Repository: userRepository,
	}
}

func (u *userService) Create(user *domain.User) error {
	err := u.Repository.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) Get(ID int64) (*domain.User, error) {
	user, err := u.Repository.Get(ID)
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

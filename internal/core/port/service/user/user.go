package user

import "microblog/internal/core/domain"

type UserService interface {
	Create(*domain.User) (error)
	Get(int64) (*domain.User, error)
}



package repository

import "microblog/internal/core/domain"

type UserRepository interface {
	Save(*domain.User) ( error)
	Get(int64) (*domain.User, error)
}

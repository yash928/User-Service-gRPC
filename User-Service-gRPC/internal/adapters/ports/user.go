package ports

import (
	"user-service-grpc/internal/core/user"

	"github.com/google/uuid"
)

type UserDB interface {
	FindUserById(id uuid.UUID) (*user.User, error)
	SaveUser(userDet *user.User) error
	FindUserGivenId(ids []string) ([]user.User, error)
	FindUserByFilter(filter user.Filter) ([]user.User, error)
}

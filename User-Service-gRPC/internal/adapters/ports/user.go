package ports

import (
	"user-service-grpc/internal/core/user"

	"github.com/google/uuid"
)

type UserDB interface {
	FindUserById(id uuid.UUID) (*user.User, error)
	SaveUser(userDet *user.User) error
	FindAllUsers() ([]user.User, error)
}

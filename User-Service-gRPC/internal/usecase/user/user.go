package user

import (
	"context"
	"errors"
	"user-service-grpc/internal/adapters/ports"
	"user-service-grpc/internal/core/user"
	"user-service-grpc/pkg/logging"

	"github.com/google/uuid"
)

type UserUsecaseImpl struct {
	userDB ports.UserDB
	log    logging.Service
}

func NewUserUsecase(userDB ports.UserDB, logs logging.Service) user.UserUsecase {
	return &UserUsecaseImpl{
		userDB: userDB,
		log:    logs,
	}
}

func (u *UserUsecaseImpl) CreateUser(ctx context.Context, userDet *user.User) error {

	err := u.userDB.SaveUser(userDet)
	if err != nil {
		u.log.ErrorWithContext(ctx, "SaveUser Error=", err)
		return user.ErrSomethingWentWrong
	}

	return nil
}

func (u *UserUsecaseImpl) FindUserById(ctx context.Context, id string) (*user.User, error) {

	uid, err := uuid.Parse(id)
	if err != nil {
		u.log.ErrorWithContext(ctx, "Error While Parsing the Id, Error=", err)
		return nil, user.ErrInvalidID
	}

	userDet, err := u.userDB.FindUserById(uid)
	if err != nil {
		u.log.ErrorWithContext(ctx, "FindUserById Error=", err)
		if errors.Is(err, user.ErrDocumentNotFound) {
			return nil, user.ErrUserNotFound
		}
		return nil, user.ErrSomethingWentWrong
	}

	return userDet, nil
}

func (u *UserUsecaseImpl) FindUsersListFromID(ctx context.Context, ids []string) ([]user.User, error) {

	users, err := u.userDB.FindUserGivenId(ids)
	if err != nil {
		u.log.ErrorWithContext(ctx, "FindAllUsers Error=", err)
		return nil, user.ErrSomethingWentWrong
	}

	return users, nil
}

func (u *UserUsecaseImpl) FindUserByFilter(ctx context.Context, filter user.Filter) ([]user.User, error) {

	userDet, err := u.userDB.FindUserByFilter(filter)
	if err != nil {
		u.log.ErrorWithContext(ctx, "FindUserByFilter Error=", err)
		return nil, user.ErrSomethingWentWrong
	}

	return userDet, nil
}

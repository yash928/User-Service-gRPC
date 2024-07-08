package user

import (
	"context"
	"errors"
	"user-service-client/internal/core/user"
	userSrv "user-service-client/internal/interfaces/output/grpc/user"
	"user-service-client/pkg/logging"
)

type UserUsecaseImpl struct {
	userServer userSrv.UserServiceClient
	log        logging.Service
}

func NewUserUsecase(userServer userSrv.UserServiceClient, logs logging.Service) user.UserUsecase {
	return &UserUsecaseImpl{
		userServer: userServer,
		log:        logs,
	}
}

// func (u *UserUsecaseImpl) CreateUser(ctx context.Context, userDet *user.User) error {

// 	err := u.userDB.SaveUser(userDet)
// 	if err != nil {
// 		u.log.ErrorWithContext(ctx, "SaveUser Error=", err)
// 		return user.ErrSomethingWentWrong
// 	}

// 	return nil
// }

func (u *UserUsecaseImpl) FindUserById(ctx context.Context, id string) (*user.User, error) {

	userDet, err := u.userServer.FindUserById(ctx, &userSrv.FindUserByIdInput{
		Id: id,
	})
	if err != nil {
		u.log.ErrorWithContext(ctx, "FindUserById Error=", err)
		if errors.Is(err, user.ErrDocumentNotFound) {
			return nil, user.ErrUserNotFound
		}
		return nil, user.ErrSomethingWentWrong
	}

	return &user.User{
		Name:          userDet.UserDet.Name,
		Address:       userDet.UserDet.Address,
		City:          userDet.UserDet.City,
		State:         userDet.UserDet.State,
		Country:       userDet.UserDet.Country,
		Pincode:       userDet.UserDet.Pincode,
		PhoneNo:       userDet.UserDet.PhoneNumber,
		MaritalStatus: userDet.UserDet.MaritalStatus,
		Height:        userDet.UserDet.Height,
	}, nil
}

// func (u *UserUsecaseImpl) FindAllUsers(ctx context.Context) ([]user.User, error) {

// 	users, err := u.userDB.FindAllUsers()
// 	if err != nil {
// 		u.log.ErrorWithContext(ctx, "FindAllUsers Error=", err)
// 		return nil, user.ErrSomethingWentWrong
// 	}

// 	return users, nil
// }

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
		Id:            userDet.UserDet.Id,
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

func (u *UserUsecaseImpl) FindUserListByID(ctx context.Context, ids []string) ([]user.User, error) {

	if len(ids) <= 0 {
		u.log.ErrorWithContext(ctx, "No IDs to read")
		return []user.User{}, nil
	}

	users, err := u.userServer.FindUsersListFromID(ctx, &userSrv.FindUsersListFromIDReq{
		Id: ids,
	})
	if err != nil {
		u.log.ErrorWithContext(ctx, "FindAllUsers Error=", err)
		return nil, user.ErrSomethingWentWrong
	}

	var userList []user.User

	for _, userDet := range users.UserDet {
		userInfo := user.User{
			Id:            userDet.Id,
			Name:          userDet.Name,
			Address:       userDet.Address,
			City:          userDet.City,
			State:         userDet.State,
			Country:       userDet.Country,
			Pincode:       userDet.Pincode,
			PhoneNo:       userDet.PhoneNumber,
			MaritalStatus: userDet.MaritalStatus,
			Height:        userDet.Height,
		}
		userList = append(userList, userInfo)
	}

	return userList, nil
}

func (u *UserUsecaseImpl) FindUserByFilter(ctx context.Context, filter user.Filter) ([]user.User, error) {

	users, err := u.userServer.FindUserByFilter(ctx, &userSrv.Filter{
		MaritalStatus: filter.MaritalStatus,
		Country:       filter.Country,
	})
	if err != nil {
		u.log.ErrorWithContext(ctx, "FindUserByFilter Error=", err)
		return nil, user.ErrSomethingWentWrong
	}

	var userList []user.User

	for _, userDet := range users.UserDet {
		userInfo := user.User{
			Id:            userDet.Id,
			Name:          userDet.Name,
			Address:       userDet.Address,
			City:          userDet.City,
			State:         userDet.State,
			Country:       userDet.Country,
			Pincode:       userDet.Pincode,
			PhoneNo:       userDet.PhoneNumber,
			MaritalStatus: userDet.MaritalStatus,
			Height:        userDet.Height,
		}
		userList = append(userList, userInfo)
	}

	return userList, nil
}

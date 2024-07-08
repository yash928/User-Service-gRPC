package user

import (
	"context"
	"log"
	"user-service-grpc/internal/core/user"
)

type Server struct {
	UnimplementedUserServiceServer
	userUC user.UserUsecase
}

func NewServer(userUC user.UserUsecase) UserServiceServer {
	return &Server{
		userUC: userUC,
	}
}

func (s *Server) FindUserById(ctx context.Context, req *FindUserByIdInput) (*FindUserByIdResponse, error) {

	userDet, err := s.userUC.FindUserById(ctx, req.Id)
	if err != nil {
		log.Println("FindUserById Error=", err)
		return nil, err
	}

	return &FindUserByIdResponse{
		UserDet: &User{
			Id:            string(userDet.Id.String()),
			Name:          userDet.Name,
			Address:       userDet.Address,
			City:          userDet.City,
			State:         userDet.State,
			Country:       userDet.Country,
			Pincode:       userDet.Pincode,
			PhoneNumber:   userDet.PhoneNo,
			MaritalStatus: userDet.MaritalStatus,
			Height:        userDet.Height,
		},
	}, nil

}

func (s *Server) FindUsersListFromID(ctx context.Context, req *FindUsersListFromIDReq) (*FindUsersListFromIDResponse, error) {

	users, err := s.userUC.FindUsersListFromID(ctx, req.Id)
	if err != nil {
		log.Print("FindUsersListFromID Error=", err)
		return nil, err
	}

	var userList []*User

	for _, userDet := range users {
		userInfo := &User{
			Id:            string(userDet.Id.String()),
			Name:          userDet.Name,
			Address:       userDet.Address,
			City:          userDet.City,
			State:         userDet.State,
			Country:       userDet.Country,
			Pincode:       userDet.Pincode,
			PhoneNumber:   userDet.PhoneNo,
			MaritalStatus: userDet.MaritalStatus,
			Height:        userDet.Height,
		}
		userList = append(userList, userInfo)
	}

	return &FindUsersListFromIDResponse{
		UserDet: userList,
	}, nil
}

func (s *Server) FindUserByFilter(ctx context.Context, filter *Filter) (*FindUserByFilterResp, error) {
	users, err := s.userUC.FindUserByFilter(ctx, user.Filter{
		Country:       filter.Country,
		MaritalStatus: filter.MaritalStatus,
	})
	if err != nil {
		log.Print("FindUserByFilter Error=", err)
		return nil, err
	}

	var userList []*User

	for _, userDet := range users {
		userInfo := &User{
			Id:            string(userDet.Id.String()),
			Name:          userDet.Name,
			Address:       userDet.Address,
			City:          userDet.City,
			State:         userDet.State,
			Country:       userDet.Country,
			Pincode:       userDet.Pincode,
			PhoneNumber:   userDet.PhoneNo,
			MaritalStatus: userDet.MaritalStatus,
			Height:        userDet.Height,
		}
		userList = append(userList, userInfo)
	}

	return &FindUserByFilterResp{
		UserDet: userList,
	}, nil
}

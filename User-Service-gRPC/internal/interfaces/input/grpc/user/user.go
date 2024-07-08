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

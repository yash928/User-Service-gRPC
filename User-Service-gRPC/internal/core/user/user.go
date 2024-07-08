package user

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Country       string    `json:"country"`
	Pincode       string    `json:"pincode"`
	PhoneNo       string    `json:"phone_number"`
	MaritalStatus string    `json:"marital_status"`
	Height        float32   `json:"height"`
}

type Filter struct {
	Country       string `json:"country"`
	MaritalStatus string `json:"marital_status"`
}

type UserUsecase interface {
	CreateUser(ctx context.Context, userDet *User) error
	FindUsersListFromID(ctx context.Context, ids []string) ([]User, error)
	FindUserById(ctx context.Context, id string) (*User, error)
	FindUserByFilter(ctx context.Context, filter Filter) ([]User, error)
}

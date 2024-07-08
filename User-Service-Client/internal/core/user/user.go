package user

import (
	"context"
)

type User struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Address       string  `json:"address"`
	City          string  `json:"city"`
	State         string  `json:"state"`
	Country       string  `json:"country"`
	Pincode       string  `json:"pincode"`
	PhoneNo       string  `json:"phone_number"`
	MaritalStatus string  `json:"marital_status"`
	Height        float32 `json:"height"`
}

type UserUsecase interface {
	// CreateUser(ctx context.Context, userDet *User) error
	// FindAllUsers(ctx context.Context) ([]User, error)
	FindUserById(ctx context.Context, id string) (*User, error)
}
